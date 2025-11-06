package notification

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/errorz"
	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/models"
)

func (s *Service) Cancel(ctx context.Context, id uuid.UUID) error {
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}

	if n.Status != models.StatusScheduled {
		return fmt.Errorf("%w: current status is %s", errorz.NotificationIsNotCancellable, n.Status)
	}

	return s.repo.UpdateStatus(ctx, id, models.StatusCancelled)
}
