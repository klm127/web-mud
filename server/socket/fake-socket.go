package socket

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

/*
Implements an interface like the gorilla websocket, but is actually used with a series of HTTP routes.

Instead of directly sending messages to and receiving messages from the client, fakeSocket caches the messages in slices. As the client polls the relevant
*/
type fakeSocket struct {
	// to transmit to client
	pending [][]byte
	// sent from client
	received   [][]byte
	rec_mutex  sync.Mutex
	pend_mutex sync.Mutex
	closed     bool
}

func NewFakeSocket() *fakeSocket {
	return &fakeSocket{
		pending:  make([][]byte, 0),
		received: make([][]byte, 0),
	}
}
func (fs *fakeSocket) ReadMessage() (int, []byte, error) {
	fmt.Println("reading messages")
	var found []byte
	for len(fs.received) == 0 {
		time.Sleep(1000)
	}
	fs.rec_mutex.Lock()
	found = fs.received[0]
	fs.received = fs.received[1:]
	fs.rec_mutex.Unlock()
	// unshift 1
	return 1, found, nil
}

func (fs *fakeSocket) addMessage(s string) {
	fs.rec_mutex.Lock()
	fs.received = append(fs.received, []byte(s))
	fs.rec_mutex.Unlock()
}

func (fs *fakeSocket) popPending() [][]byte {
	fs.pend_mutex.Lock()
	out := fs.pending
	fs.pending = make([][]byte, 0)
	fs.pend_mutex.Unlock()
	return out
}

func (fs *fakeSocket) WriteMessage(mtype int, val []byte) error {
	fs.pend_mutex.Lock()
	fs.pending = append(fs.pending, val)
	fs.pend_mutex.Unlock()
	return nil
}

func (fs *fakeSocket) Close() error {
	sr := sock_response{"close", ""}
	b, _ := json.Marshal(sr)
	fs.pending = append(fs.pending, b)
	fs.closed = true
	return nil
}
