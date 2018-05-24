package cache

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// SimpleStore implements the Store Manager interface using go maps
type SimpleStore struct {
	data  map[interface{}]interface{}
	mutex *sync.Mutex
}

// NewSimpleStore creates a SimpleStore based StoreManager
func NewSimpleStore() *SimpleStore {
	store := new(SimpleStore)
	store.data = make(map[interface{}]interface{})
	store.mutex = &sync.Mutex{}
	return store
}

// Store a value
func (store *SimpleStore) Store(key interface{}, value interface{}) (bool, error) {
	store.mutex.Lock()
	store.data[key] = value
	store.mutex.Unlock()
	return true, nil
}

// Load a value from the cache
func (store *SimpleStore) Load(key interface{}) (interface{}, error) {
	value, ok := store.data[key]
	if !ok {
		return nil, errors.New("Value is not stored")
	}
	fmt.Printf("this is a Load type is %v...", reflect.TypeOf(value).String())

	switch value.(type) {
	case chan interface{}:
		store.mutex.Lock()
		result := <-value.(chan interface{}) // blocking call, might  work on Go 1.2
		close(value.(chan interface{}))      // critical
		store.data[key] = result
		fmt.Printf("storing %s", store.data[key])
		store.mutex.Unlock()
		return result, nil
	default:
		fmt.Println("DID Not a channel")
	}
	return value, nil
}

// Delete a value from the cache
func (store *SimpleStore) Delete(key interface{}) error {
	_, ok := store.data[key]
	if !ok {
		return errors.New("Value does not exist")
	}
	delete(store.data, key)
	return nil
}
