package promotion

import (
	"context"

	"github.com/ddoffy/shopping/model"
)

type EUsecase interface {
	Fetch(ctx context.Context) ([]*model.Promotion, error)
	Store(ctx context.Context, a *model.Promotion) (int64, error)
	Delete(ctx context.Context, id int) (bool, error)
}
