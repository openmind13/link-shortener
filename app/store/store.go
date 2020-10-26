package store

import "github.com/openmind13/link-shortener/app/model"

// Store ...
type Store interface {
	Add(model.Data) error
	Get(model.Data) (model.Data, error)
}
