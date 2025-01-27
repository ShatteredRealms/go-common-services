package mapbus

import (
	"context"

	"github.com/ShatteredRealms/go-common-service/pkg/bus"
	"github.com/ShatteredRealms/go-common-service/pkg/common"
	"github.com/google/uuid"
)

type Service interface {
	bus.BusProcessor[Message]
	GetMaps(ctx context.Context) (*Maps, error)
	GetMapById(ctx context.Context, mapId string) (*Map, error)
}

type service struct {
	bus.DefaultBusProcessor[Message]
}

func NewService(
	repo Repository,
	mapBus bus.MessageBusReader[Message],
) Service {
	return &service{
		DefaultBusProcessor: bus.DefaultBusProcessor[Message]{
			Reader: mapBus,
			Repo:   repo,
		},
	}
}

// GetMapById implements MapService.
func (d *service) GetMapById(ctx context.Context, mapId string) (*Map, error) {
	id, err := uuid.Parse(mapId)
	if err != nil {
		return nil, common.ErrInvalidId
	}

	return d.Repo.(Repository).GetById(ctx, &id)
}

// GetMaps implements MapService.
func (d *service) GetMaps(ctx context.Context) (*Maps, error) {
	return d.Repo.(Repository).GetAll(ctx)
}
