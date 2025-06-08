package service

import (
	"context"

	"github.com/keribend/hello-web/internal/repository"
)

type Service struct {
	repo *repository.Queries
}

func New(r *repository.Queries) *Service {
	return &Service{r}
}

func (s *Service) FindAllEvents(ctx context.Context) ([]repository.Event, error) {
	return s.repo.FindAllEvents(ctx)
}

func (s *Service) FindEvent(ctx context.Context, id int64) (repository.Event, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) AddCheckinToEvent(ctx context.Context, eventId int64) error {
	return s.repo.InsertCheckinForEvent(ctx, eventId)
}
