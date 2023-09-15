package iworld

type ISeen interface {
	GetSource() IExists
	GetTarget() IExists
	GetMessage(IBeing) string
}
