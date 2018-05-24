package cache

// StoreManager for cache interface
type StoreManager interface {
	Store(interface{}, interface{}) (bool, error)
	Load(interface{}) (interface{}, error)
	Delete(interface{}) error
}
