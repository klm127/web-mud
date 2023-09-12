package iworld

// Anything that exists in the world such that it can be interacted with.
type IExists interface {
	// Get this things name
	Name() string
	// Get this things description
	Desc() string
	// Get the image, if extant
	Img() *string
}
