package biz

import (
	"context"
	"social/todo/list/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}
type DeleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *DeleteItemBiz {
	return &DeleteItemBiz{store: store}
}

func (biz *DeleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
