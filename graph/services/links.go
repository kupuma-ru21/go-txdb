package services

import (
	"context"
	"go-txdb/graph/db"
	"go-txdb/graph/model"
	"strconv"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (u *linkService) UpdateLink(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	idFromInput, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}
	dbLink, err := db.Links(
		qm.Select(
			db.LinkColumns.ID,
			db.LinkColumns.Title,
			db.LinkColumns.Address,
		),
		db.LinkWhere.ID.EQ(idFromInput),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	dbLink.Title = null.StringFrom(input.Title)
	dbLink.Address = null.StringFrom(input.Address)
	if _, err := dbLink.Update(ctx, u.exec, boil.Infer()); err != nil {
		return nil, err
	}

	return &model.Link{
		ID:      input.ID,
		Title:   dbLink.Title.String,
		Address: dbLink.Address.String,
	}, nil
}
