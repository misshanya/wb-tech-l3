package link

import (
	"context"
	"errors"
	"fmt"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/errorz"
	"github.com/misshanya/wb-tech-l3/url-shortener/pkg/base62"
)

func (s *service) Create(ctx context.Context, url string) (string, error) {
	// Check if URL exists in db
	id, err := s.repo.GetIDByURL(ctx, url)
	if err == nil {
		return base62.Encode(id), nil
	} else if !errors.Is(err, errorz.LinkNotFound) {
		return "", fmt.Errorf("failed to get id by url: %w", err)
	}

	link, err := s.repo.Create(ctx, url)
	if err != nil {
		return "", fmt.Errorf("failed to create link: %w", err)
	}

	return base62.Encode(link.ID), nil
}
