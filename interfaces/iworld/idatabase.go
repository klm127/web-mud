package iworld

import "time"

// Objects that are related to underlying database data
type IDatabase[T interface{}] interface {

	// Retrieve the underyling database data
	GetData() T
	// Save (save to database) if dirty
	Save()
	// Return whether this is dirty
	IsDirty() bool
	// Get the database record id
	GetId() int64
	// Get the time when this was instanced from database data
	GetInstancedTime() time.Time
	// Get the last time this was interacted with
	GetLastInteractionTime() time.Time
	// Get time since last interaction
	GetTimeSinceLastInteraction() time.Duration
	// Refresh time since last interaction
	RefreshTime()
}
