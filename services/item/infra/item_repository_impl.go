package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/domain"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type itemRepository struct {
	db *database.DB
}

func NewItemRepository(db *database.DB) domain.ItemRepository {
	return &itemRepository{db: db}
}

// GetByID implements domain.ItemRepository.
func (repo *itemRepository) GetByID(ctx context.Context, id domain.ItemID) (*domain.Item, error) {
	tx, err := database.GetTxByContext(ctx)
	if err != nil {
		var err error
		tx, err = repo.db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		defer func() {
			if err := recover(); err != nil {
				tx.Rollback()
			}
		}()
		defer tx.Commit()
	}

	var item item
	err = tx.NewSelect().
		Model(&item).
		Where("? = ?", bun.Ident("id"), id).
		Scan(ctx)
	if err != nil {
		switch errors.Cause(err) {
		//TODO: add custom error for not found
		default:
			return nil, errors.WithStack(err)
		}
	}

	return item.ConvertToEntity(), nil
}

// GetByIDs implements domain.ItemRepository.
func (*itemRepository) GetByIDs(ctx context.Context, id domain.ItemID) ([]*domain.Item, error) {
	panic("unimplemented")
}
