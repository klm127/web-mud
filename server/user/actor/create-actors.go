package actor

import (
	"sync"

	"github.com/pwsdc/web-mud/interfaces/iserver"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

// A map of all active actors. The key is not related to their user ID
var actors map[int64]iactor.IActor

// The next available ID in the map.
var next_id int64

// A mutex to control map and ID access.
var actors_map_mutex sync.Mutex

// The default commands to be loaded into every new actor.
var default_command_groups []iactor.ICommandGroup

func init() {
	next_id = 0
	actors = make(map[int64]iactor.IActor)
	default_command_groups = make([]iactor.ICommandGroup, 0)
}

// Should only be called while mutex is locked, e.g., by StartActor
func getNextId() int64 {
	_, ok := actors[next_id]
	for ok {
		next_id++
		_, ok = actors[next_id]
	}
	return next_id
}

func removeActor(id int64) {
	delete(actors, id)
}

func SetDefaultCommandGroups(groups ...iactor.ICommandGroup) {
	default_command_groups = append(default_command_groups, groups...)
}

func StartActor(connection iserver.IConnection) iactor.IActor {
	actors_map_mutex.Lock()
	id := getNextId()
	actors[id] = newActor(id, connection)
	for _, group := range default_command_groups {
		actors[id].SetCommandGroup(group)
	}
	actors_map_mutex.Unlock()
	return actors[id]
}

func GetActor(id int64) (iactor.IActor, bool) {
	actor, ok := actors[id]
	return actor, ok
}

func Traverse(cb func(*map[int64]iactor.IActor), lock bool) {
	if lock {
		actors_map_mutex.Lock()
	}
	cb(&actors)
	if lock {
		actors_map_mutex.Unlock()
	}
}

func logoutAnyExtantActors(user_id int64) {
	for _, v := range actors {
		user := v.GetUser()
		if user != nil {
			if user.ID == user_id {
				v.Errorf("You have logged in from another location.")
				v.Disconnect()
			}

		}
	}
}
