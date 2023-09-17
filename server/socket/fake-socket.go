package socket

import "sync"

type fakeSocket struct {
	pending    [][]byte
	received   [][]byte
	rec_mutex  sync.Mutex
	pend_mutex sync.Mutex
}

func NewFakeSocket() *fakeSocket {
	return &fakeSocket{
		pending:  make([][]byte, 0),
		received: make([][]byte, 0),
	}
}
func (fs *fakeSocket) ReadMessage() (int, []byte, error) {
	var found []byte
	for len(found) == 0 {

	}
	fs.rec_mutex.Lock()
	found = fs.received[0]
	// unshift 1
	return 1, found, nil
}

func (fs *fakeSocket) WriteMessage(mtype int, val []byte, e error) error {
	fs.pend_mutex.Lock()
	fs.pending = append(fs.pending, val)
	return nil
}

func (fs *fakeSocket) Close() error {
	return nil
}
