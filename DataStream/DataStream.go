package DataStream

import (
	"fmt"
	"sync"
)

type DataStream struct {
	Closed		bool
	Data		*[]byte
	ReadLock	sync.Mutex
	WriteLock	sync.Mutex
}

func(d *DataStream) Buffered() int {
	return len(*d.Data)
}

func(d *DataStream) Close() {
	d.WriteLock.Lock()
	defer d.WriteLock.Unlock()

	d.Closed = true
}

func(d *DataStream) Write(buf *[]byte) {
	if d.Closed {
		panic(fmt.Errorf("DataStream is already closed"))
	}

	d.WriteLock.Lock()
	defer d.WriteLock.Unlock()

	for i := 0; i < len(*buf); i++ {
		*d.Data = append(*d.Data, (*buf)[i])
	}
}

func(d *DataStream) Read(buf *[]byte, sz *int) bool {
	if len(*d.Data) == 0 && d.Closed {
		return false
	}

	d.ReadLock.Lock()
	defer d.ReadLock.Unlock()

	for len(*d.Data) == 0 {}

	if len(*buf) > 0 {
		*sz = copy(*buf, *d.Data)
		*d.Data = (*d.Data)[*sz:]
	} else {
		*buf = *d.Data
		*d.Data = (*d.Data)[len(*buf):]
		*sz = len(*buf)
	}

	return true
}

func New() *DataStream {
	cl := new(DataStream)
	cl.Data = new([]byte)
	cl.Closed = false
	return cl
}