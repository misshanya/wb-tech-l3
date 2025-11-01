package notification

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/wb-go/wbf/retry"
	"github.com/wb-go/wbf/zlog"
)

func (s *Sender) SendNotification(ctx context.Context, title, content, receiver string) error {
	url := fmt.Sprintf("%s/%s", s.ntfyURL, receiver)

	return retry.Do(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(content))
		if err != nil {
			zlog.Logger.Error().
				Err(err).
				Msg("failed to create http request to send notification via ntfy")
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Title", title)

		resp, err := s.client.Do(req)
		if err != nil {
			zlog.Logger.Error().
				Err(err).
				Msg("failed to send message")
			return fmt.Errorf("failed to send request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				zlog.Logger.Error().
					Err(err).
					Msg("failed to read response body")
			}
			zlog.Logger.Error().
				Err(err).
				Int("status_code", resp.StatusCode).
				Str("body", string(respBody)).
				Msg("failed to send message: unexpected status code")
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		return nil
	}, s.retry)
}
