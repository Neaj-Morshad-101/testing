package library

import (
	"context"
	"github.com/Neaj-Morshad-101/testing/books"
)

func FetchSummary(ctx context.Context, book *books.Book) (string, error) {
	return "Jean Valjean", nil
}
