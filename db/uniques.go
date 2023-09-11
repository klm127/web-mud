package db

import (
	"context"
)

/*
Determines if a name is unique in the system. I.e: not used by any users or beings.
*/
func (s *tStore) UniqueName(txt string) bool {
	_, err := s.Query.GetBeingByName(context.Background(), txt)
	if err != nil {
		return false
	}
	_, err = s.Query.GetUserByName(context.Background(), txt)
	if err != nil {
		return false
	}
	return true
}
