package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/misshanya/wb-tech-l3/delayed-notifier/internal/infra/telegram/notification/dto"
	"github.com/wb-go/wbf/retry"
	"github.com/wb-go/wbf/zlog"
)

func (s *Sender) SendNotification(ctx context.Context, title, content, receiver string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.botApiKey)

	body := &dto.Message{
		ChatID:    receiver,
		Text:      fmt.Sprintf("%s\n\n%s", title, content),
		ParseMode: "markdownV2",
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	return retry.Do(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyJson))
		if err != nil {
			zlog.Logger.Error().
				Err(err).
				Msg("failed to create http request to send notification")
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

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
