package service

import (
	"context"
	"log"

	"github.com/sebastianmendez/trail-browser/internal/store"
	"github.com/sebastianmendez/trail-browser/internal/store/model"
)

// this layer takes the responsability to handle any extra processing if required, it'd be scalable by defining an interface
// that allow multiple concrete objects to implement the same set of functions
func List(ctx context.Context, params map[string]string) ([]model.Trail, error) {
	res, err := store.List(ctx, params)
	if err != nil {
		log.Fatalf("Error fetching the data: %v", err)
	}
	return res, nil
}
