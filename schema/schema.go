package schema

import (
	"time"

	"github.com/google/uuid"
)

// Arriving : true if incoming, false if outgoing
// Granted : true if access request granted
type Activity struct {
	UserId    uuid.UUID `json:"userid"`
	Arriving  bool      `json:"arriving"`
	Granted   bool      `json:"granted"`
	TimeStamp time.Time `json:"timestamp"`
}

func NewActivity(userid uuid.UUID, arriving bool, granted bool, timing time.Time) *Activity {
	return &Activity{
		UserId:    userid,
		Arriving:  arriving,
		TimeStamp: timing,
		Granted:   granted,
	}
}
