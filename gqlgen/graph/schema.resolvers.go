package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kalesecarpenter/quotes-starter/gqlgen/graph/generated"
	"github.com/kalesecarpenter/quotes-starter/gqlgen/graph/model"
)

// Quote is the resolver for the quote field.

func (r *queryResolver) Quote(ctx context.Context) (*model.Quote, error) {
	quote := model.Quote{
		ID:     "1234",
		Quote:  "Happy Coding",
		Author: "Kaye Carp",
	}

	return &quote, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
