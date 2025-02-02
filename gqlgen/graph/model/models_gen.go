// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewQuote struct {
	// Input a new quote to the database
	Quote string `json:"quote"`
	// Author is the creator of the quote
	Author string `json:"author"`
}

type Quote struct {
	// The id of a quote
	ID string `json:"id"`
	// All quotes are strings
	Quote string `json:"quote"`
	// Author is the creator of the quote
	Author string `json:"author"`
}
