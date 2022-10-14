package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kalesecarpenter/quotes-starter/gqlgen/graph/generated"
	"github.com/kalesecarpenter/quotes-starter/gqlgen/graph/model"
)

// GetRandomQuote is the resolver for the getRandomQuote field.
func (r *queryResolver) GetRandomQuote(ctx context.Context) (*model.Quote, error) {
	//defer response.Body.Close() -- delay the execution of the function or method or an anonymous method until the nearby functions returns
	request, err := http.NewRequest("GET", "http://34.160.48.181/quotes", nil)
	request.Header.Set("x-api-key", "COCKTAILSAUCE")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, _ := client.Do(request)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var newGoQuote model.Quote
	json.Unmarshal(responseData, &newGoQuote)

	/*fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Quote))

	for i := 0; i < len(responseObject.Quote); i++ {
		fmt.Println(responseObject.Quote[i].ID.Author)
	}*/
	return &newGoQuote, nil
}

/*quote := &model.Quote{
	ID:     "12-3422-1-23432-009954",
	Quote:  "I love being my authentic self",
	Author: "Kalese Carpenter",
}*/
//return quote, nil

// GetQuoteByID is the resolver for the getQuoteById field.
func (r *queryResolver) GetQuoteByID(ctx context.Context, id string) (*model.Quote, error) {
	panic(fmt.Errorf("not implemented: GetQuoteByID - getQuoteById"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
