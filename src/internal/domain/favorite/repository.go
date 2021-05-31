package favorite

import (
	"context"
)

type FavoriteRepository interface {
	Create(ctx context.Context, f *Favorite) error
}
