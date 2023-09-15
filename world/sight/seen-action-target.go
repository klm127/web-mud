package sight

import (
	"fmt"

	"github.com/pwsdc/web-mud/interfaces/iworld"
)

// Not yet used - builder?
type seen_targetted_action struct {
	msg_to_target      *string
	msg_to_bystanders  *string
	msg_to_source      *string
	append_target_name bool
	source             iworld.IExists
	target             iworld.IExists
}

func (s *seen_targetted_action) GetTarget() iworld.IExists {
	return s.target
}

func (s *seen_targetted_action) GetSource() iworld.IExists {
	return s.source
}

func (s *seen_targetted_action) GetMessage(being iworld.IBeing) string {
	if being == s.source {
		return s.seenBySource(being)
	} else if being == s.target {
		return s.seenByTarget(being)
	} else {
		return s.seenByBystander(being)
	}
}

func (s *seen_targetted_action) seenBySource(being iworld.IBeing) string {
	if s.msg_to_source == nil {
		if s.append_target_name && s.target != nil {
			return fmt.Sprintf("You see yourself do something strange to %s.", s.target.Name())
		} else {
			return fmt.Sprintf("You see yourself do something to nothing.")
		}
	} else {
		if s.append_target_name {
			if s.target != nil {
				return fmt.Sprintf("You %s %s.", *s.msg_to_source, s.target.Name())
			} else {
				return fmt.Sprintf("You %s.", *s.msg_to_source)
			}
		} else {
			return fmt.Sprintf("You %s.", *s.msg_to_source)
		}
	}
}

func (s *seen_targetted_action) seenByTarget(being iworld.IBeing) string {
	if s.msg_to_target == nil {
		if s.source == nil {
			return fmt.Sprintf("You see something happen to you.")
		} else {
			return fmt.Sprintf("%s does something to you.", s.source.Name())
		}
	} else {
		if s.source == nil {
			return fmt.Sprintf(*s.msg_to_target)
		} else {
			return fmt.Sprintf("%s %s", s.source.Name(), *s.msg_to_target)
		}
	}
}

func (s *seen_targetted_action) seenByBystander(being iworld.IBeing) string {
	if s.msg_to_bystanders == nil {
		if s.source == nil {
			if s.target == nil || !s.append_target_name {
				return fmt.Sprintf("You see something strange nearby.")
			} else {
				return fmt.Sprintf("Something strange happens to %s.", s.target.Name())
			}
		} else {
			if s.target == nil || !s.append_target_name {
				return fmt.Sprintf("%s does something strange.", s.source.Name())
			} else {
				return fmt.Sprintf("%s does something strange to %s.", s.source.Name(), s.target.Name())
			}
		}
	} else {
		if s.source == nil {
			if s.target == nil || !s.append_target_name {
				return fmt.Sprintf(*s.msg_to_bystanders)
			} else {
				return fmt.Sprintf("%s %s.", *s.msg_to_bystanders, s.target.Name())
			}
		} else {
			if s.target == nil || !s.append_target_name {
				return fmt.Sprintf("%s %s.", *s.msg_to_bystanders, s.source.Name())
			} else {
				return fmt.Sprintf("%s %s %s.", s.source.Name(), *s.msg_to_bystanders, s.target.Name())
			}
		}
	}
}
