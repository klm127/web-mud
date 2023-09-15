package iworld

type IMoveReq interface {
	// if able to move; enough time has passed
	CanMove() bool
	// Get repeats; how many more moves remaining
	GetRepeats() int8
	// Get the direction; one of enum
	GetDirection() string
	// Move to a given room. nil could be passed for a failed move.
	Move(to IRoom)
	// Get what initiated this move
	GetMover() IBeing
}
