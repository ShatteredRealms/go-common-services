package mapbus

import (
	"time"

	"github.com/ShatteredRealms/go-common-service/pkg/bus"
	"github.com/google/uuid"
)

type Message struct {
	// Id is the unique identifier of the character
	Id uuid.UUID `json:"id"`

	// Deleted is a flag indicating if the character has been deleted
	Deleted bool `json:"deleted"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BusReader bus.MessageBusReader[Message]
type BusWriter bus.MessageBusWriter[Message]

func (m Message) GetType() bus.BusMessageType {
	return bus.BusMessageType("sro.gameserver.map")
}

func (m Message) GetId() string {
	return m.Id.String()
}

func (m Message) WasDeleted() bool {
	return m.Deleted
}
