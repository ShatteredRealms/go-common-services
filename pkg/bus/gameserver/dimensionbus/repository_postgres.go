package dimensionbus

import (
	"context"
	"fmt"

	"github.com/ShatteredRealms/go-common-service/pkg/srospan"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository struct {
	gormdb *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Dimension{})
	return &postgresRepository{gormdb: db}
}

// Save implements DimensionRepository.
func (p *postgresRepository) Save(ctx context.Context, data any) error {
	dimension, ok := data.(Dimension)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data)
	}

	updateSpanWithDimension(ctx, dimension.Id)
	return p.db(ctx).Create(&dimension).Error
}

// Delete implements DimensionRepository.
func (p *postgresRepository) Delete(ctx context.Context, dimensionId string) error {
	updateSpanWithDimension(ctx, dimensionId)
	return p.db(ctx).Clauses(clause.Returning{}).Delete(&Dimension{}, "id = ?", dimensionId).Error
}

// GetById implements DimensionRepository.
func (p *postgresRepository) GetById(ctx context.Context, dimensionId string) (dimension *Dimension, _ error) {
	result := p.db(ctx).Where("id = ?", dimensionId).Find(&dimension)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	updateSpanWithDimension(ctx, dimensionId)
	return dimension, nil
}

// GetAll implements DimensionRepository.
func (p *postgresRepository) GetAll(ctx context.Context) (dimensions *Dimensions, _ error) {
	return dimensions, p.db(ctx).Find(dimensions).Error
}

func (p *postgresRepository) db(ctx context.Context) *gorm.DB {
	return p.gormdb.WithContext(ctx)
}

func updateSpanWithDimension(ctx context.Context, dimensionId string) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		srospan.DimensionId(dimensionId),
	)
}
