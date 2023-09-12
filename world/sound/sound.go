package sound

import (
	"github.com/pwsdc/web-mud/interfaces/iworld"
)

type sound struct {
	msg      string
	source   iworld.IExists
	loudness int
}

func (s *sound) GetMessage() string {
	return s.msg
}

func (s *sound) HasSpecialMessage() bool {
	return false
}

func (s *sound) GetSpecialMessage() []byte {
	return []byte(s.source.Name() + " says " + s.msg)
}

func (s *sound) GetSourceName() string {
	return s.source.Name()
}

func (s *sound) GetSource() iworld.IExists {
	return s.source
}

func (s *sound) GetLoudness() int {
	return s.loudness
}

func New(source iworld.IExists, msg string, loudness int) iworld.ISound {
	return &sound{
		msg:      msg,
		loudness: loudness,
		source:   source,
	}
}
