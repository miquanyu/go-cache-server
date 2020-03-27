package cache

// #include "rocksdb/c.h"
// #cgo CFLAGS: -I/usr/local/Cellar/rocksdb/6.7.3/include
// #cgo LDFLAGS: -L/usr/local/Cellar/rocksdb/6.7.3/lib -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"

type rocksdbCache struct {
	db *C.rocksdb_t
	ro *C.rocksdb_readoptions_t
	wo *C.rocksdb_writeoptions_t
	e  *C.char
	ch chan *pair
}
