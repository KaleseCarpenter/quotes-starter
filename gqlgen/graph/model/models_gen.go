// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// Author of a Quote
type Author struct {
	// The id of the author
	ID string `json:"id"`
	// The name of the person who created the quote
	Name string `json:"name"`
}

type Quotes struct {
	// The id of a quote
	ID string `json:"id"`
	// All quotes are strings
	Quote string `json:"quote"`
	// Author is the creator of the quote
	Author *Author `json:"author"`
}
