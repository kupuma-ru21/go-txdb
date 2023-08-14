package services

import (
	"context"
	"go-txdb/graph/db"
	"go-txdb/graph/model"
	"strconv"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type linkService struct {
	exec boil.ContextExecutor
}

func (u *linkService) CreateLink(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	newLink := db.Link{
		Title:   null.StringFrom(input.Title),
		Address: null.StringFrom(input.Address),
	}
	err := newLink.Insert(ctx, u.exec, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &model.Link{
		ID:      strconv.Itoa(newLink.ID),
		Title:   newLink.Title.String,
		Address: newLink.Address.String,
	}, nil
}
