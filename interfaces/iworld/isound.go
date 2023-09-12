package iworld

type ISound interface {
	GetMessage() string
	HasSpecialMessage() bool
	GetSpecialMessage() []byte
	GetSourceName() string
	GetSource() IExists
	GetLoudness() int
}
