package services

import (
	"context"
	"go-txdb/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type LinkService interface {
	CreateLink(ctx context.Context, input model.CreateLinkInput) (*model.Link, error)
}

type Services interface {
	LinkService
}

type services struct {
	*linkService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		linkService: &linkService{exec: exec},
	}
}
