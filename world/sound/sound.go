package sound

import "github.com/pwsdc/web-mud/interfaces/iworld"

type sound struct {
	msg      string
	source   iworld.IExists
	loudness int
}

func (s *sound) GetMessage() string {
	return s.msg
}

func (s *sound) GetSourceName() string {
	return s.source.Name()
}
