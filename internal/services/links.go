package services

import (
	"context"

	"github.com/steliosmagalios/link-shortener/internal/database/repository"
)

type LinkService struct {
	queries *repository.Queries
	ctx     *context.Context
}

func NewLinkService(queries *repository.Queries, ctx *context.Context) LinkService {
	return LinkService{
		queries: queries,
		ctx:     ctx,
	}
}

func (s *LinkService) GetAll() ([]repository.Link, error) {
	return s.queries.GetAllLinks(*s.ctx)
}

func (s *LinkService) GetLinkFromSlug(slug string) (string, error) {
	return s.queries.GetLinkFromSlug(*s.ctx, slug)
}
