package link

import (
	"context"

	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent/click"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/db/ent/link"
	"github.com/misshanya/wb-tech-l3/url-shortener/internal/models"
)

type userAgentStats struct {
	UserAgent string `json:"user_agent"`
	Count     int    `json:"count"`
}
type linkStatistics []userAgentStats

func (r *repo) GetLinkStatistics(ctx context.Context, linkID int64) (models.LinkStatistics, error) {
	var internalStats linkStatistics

	err := r.client.Click.
		Query().
		Where(click.HasLinkWith(link.ID(linkID))).
		GroupBy(click.FieldUserAgent).
		Aggregate(ent.Count()).
		Scan(ctx, &internalStats)
	if err != nil {
		return nil, err
	}

	result := make(models.LinkStatistics, len(internalStats))
	for i, stats := range internalStats {
		result[i] = models.UserAgentStats{
			UserAgent: stats.UserAgent,
			Count:     stats.Count,
		}
	}

	return result, nil
}
