package store

type Key interface {
	Less(interface{}) bool
	HasPrefix(interface{}) bool
}
type Value interface {
	V()
}

type Store interface {
	Get(Key) (Value, error)
	Put(Key, Value) error
	Delete(Key) error
	PrefixScan(Key, int) []interface{}
}
