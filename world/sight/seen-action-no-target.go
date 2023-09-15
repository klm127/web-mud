package sight

import (
	"fmt"

	"github.com/pwsdc/web-mud/interfaces/iworld"
)

type seen_solo_action struct {
	msg_to_bystanders *string
	msg_to_source     *string
	source            iworld.IExists
}

func NewSeen(source iworld.IExists, msg_to_source *string, msg_to_bystanders *string) iworld.ISeen {
	return &seen_solo_action{
		msg_to_bystanders: msg_to_bystanders,
		msg_to_source:     msg_to_source,
		source:            source,
	}
}

func (s *seen_solo_action) GetSource() iworld.IExists {
	return s.source
}

func (s *seen_solo_action) GetTarget() iworld.IExists {
	return nil
}

func (s *seen_solo_action) GetMessage(being iworld.IBeing) string {
	if being == s.source {
		if s.msg_to_source != nil {
			return *s.msg_to_source
		} else {
			return "You do something, but you're not sure what."
		}
	} else {
		if s.msg_to_bystanders != nil {
			if s.source == nil {
				return *s.msg_to_bystanders
			} else {
				return fmt.Sprintf("%s %s", s.source.Name(), *s.msg_to_bystanders)
			}
		} else {
			if s.source == nil {
				return "You see something happen nearby."
			} else {
				return fmt.Sprintf("%s does something confusing.", s.source.Name())
			}
		}
	}
}
