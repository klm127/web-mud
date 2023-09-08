package base

import "sync"

var mutex_last_id sync.Mutex
var last_id int64
var Actors map[int64]*Actor

func init() {
	last_id = 0
	mutex_last_id.Unlock()
	Actors = make(map[int64]*Actor)
}

func nextId() int64 {
	var found_id int64
	mutex_last_id.Lock()
	_, ok := Actors[last_id]
	if ok {
		for ok {
			last_id++
			_, ok = Actors[last_id]
		}
		found_id = last_id
	} else {
		found_id = last_id
		last_id++
	}
	mutex_last_id.Unlock()
	return found_id
}

func removeActor(id int64) {
	delete(Actors, id)
}
