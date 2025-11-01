package notification

import (
	"net/http"
	"time"

	"github.com/wb-go/wbf/retry"
)

type Sender struct {
	client  *http.Client
	ntfyURL string
	retry   retry.Strategy
}

func New(
	client *http.Client,
	ntfyURL string,
	retryAttempts int,
	retryDelay time.Duration,
	retryBackoff float64,
) *Sender {
	return &Sender{
		client:  client,
		ntfyURL: ntfyURL,
		retry: retry.Strategy{
			Attempts: retryAttempts,
			Delay:    retryDelay,
			Backoff:  retryBackoff,
		},
	}
}
