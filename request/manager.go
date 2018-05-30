package request

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"

	"github.com/jorgevgut/goworkshop/cache"
)

const (
	defaultMaxConns = 3
)

// Manager is an interface for perfoming requests with additional
// configurations and settings
type Manager interface {
	SetMaxConnections(int)
	UseCache(cache.StoreManager)
	Requester
}

var instance *ManagerImpl

// ManagerImpl implements the Manager interface
type ManagerImpl struct {
	maxConns       int
	cache          cache.StoreManager // zero value false, by default cache is not used
	cacheAvailable bool
	client         *http.Client
	requestQueue   chan func()
	currentConns   int
	mutex          *sync.Mutex
	//currentRequests     map[string]struct{} // this is a hashSet, also is not thread safe(use sync)
	syncCurrentRequests *sync.Map
}

// Instance is a singleton constructor for a request Manager
func Instance() *ManagerImpl {
	if instance == nil {
		instance = new(ManagerImpl)
		instance.maxConns = defaultMaxConns
		instance.mutex = &sync.Mutex{}
		instance.client = &http.Client{}
		instance.requestQueue = make(chan func(), instance.maxConns)
		//instance.currentRequests = make(map[string]struct{})
		instance.syncCurrentRequests = new(sync.Map)
		for i := 0; i < instance.maxConns; i++ {
			go instance.processRequests() //1 at a time
		}
	}
	return instance
}

func (manager *ManagerImpl) processRequests() {
	for f := range manager.requestQueue {
		f()
		fmt.Println(manager.getStatus())
	}
}

func (manager *ManagerImpl) pushToQueue(f func()) {
	manager.requestQueue <- f
}

// SetMaxConnections sets how many simultaneous requests are performed
// with the HTTP client
func (manager *ManagerImpl) SetMaxConnections(maxValue int) {
	manager.maxConns = maxValue
}

// UseCache specifies that cache will be needed
func (manager *ManagerImpl) UseCache(store cache.StoreManager) {
	manager.cache = store
	manager.cacheAvailable = true
}

// String implements Stringer interface
func (manager *ManagerImpl) String() string {
	return fmt.Sprintf(
		"ManagerImpl type:\n-max simultaneous connections:%d\n-using cache:%t\n",
		manager.maxConns,
		manager.cache)
}

// Request method sends HTTP requests based on given options
func (manager *ManagerImpl) Request(opts Options) chan *http.Response {
	responseChan := make(chan *http.Response, 1)
	// starts rutine, will execute on separate thread while channel is inmediatly
	// returned to the consumer
	manager.pushToQueue(func() {
		//mutex := manager.mutex
		//mutex.Lock()
		// set info
		manager.currentConns++
		//manager.currentRequests[opts.URL] = struct{}{}
		manager.syncCurrentRequests.Store(opts.URL, struct{}{})

		// Get from cache, if not available perform req
		var response *http.Response
		response, inCache := manager.getFromCache(opts)
		// if not in cache
		if !inCache {
			req, _ := http.NewRequest(
				opts.Method,
				opts.URL,
				nil)
			// gather headers
			for key, value := range opts.Headers {
				req.Header.Add(key, value)
			}

			cacheResponseChannel := make(chan interface{}, 1)   // buffered channel is highly important
			manager.cache.Store(opts.URL, cacheResponseChannel) // nolint: errcheck
			//mutex.Unlock()

			// ommiting error management :(
			response, _ = manager.client.Do(req)
			cacheResponseChannel <- response
			/*_, err := manager.cache.Store(opts.URL, response)

			if err != nil {
				responseChan <- nil
			} //*/
		}

		fmt.Println(response.Status)
		fmt.Println(manager.getStatus())
		responseChan <- response // write the response to the channel
		manager.currentConns--   // not thread safe
		//delete(manager.currentRequests, opts.URL)
		manager.syncCurrentRequests.Delete(opts.URL)
		//}()
	})
	return responseChan // channel is resturned asynchronously
}

// Private functions usually can be added at the bottom of a go file

func (manager *ManagerImpl) getStatus() string {
	var currentRequests bytes.Buffer
	/*for key := range manager.currentRequests {
		currentRequests.WriteString(key)
	}*/
	manager.syncCurrentRequests.Range(func(key interface{}, value interface{}) bool {
		currentRequests.WriteString(key.(string))
		return true
	})
	return fmt.Sprintf(
		"Current connections %v:\n--requests: %v \n",
		manager.currentConns,
		currentRequests.String(),
	)
}

// check's from cache and returns ok true if successful
func (manager *ManagerImpl) getFromCache(opts Options) (*http.Response, bool) {
	if !manager.cacheAvailable {
		fmt.Println("No Cahce available")
		return nil, false
	}

	value, err := manager.cache.Load(opts.URL)
	if err != nil {
		fmt.Println(err)
	}
	switch value.(type) {

	case *http.Response:
		fmt.Println("------getting cached response status:------",
			value.(*http.Response).Status,
			"-----",
		)
		return value.(*http.Response), true
	}

	return nil, false
}
