package sourcequery

import (
	"fmt"
	"time"

	"github.com/rumblefrog/go-a2s"
)

const (
	timeoutDuration = 15 * time.Second
)

type SourceQueryIntegration struct {
	client *a2s.Client
}

func New(address string) (*SourceQueryIntegration, error) {
	client, err := a2s.NewClient(
		address, a2s.TimeoutOption(timeoutDuration),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new source query client: %w", err)
	}

	return &SourceQueryIntegration{
		client: client,
	}, nil
}

func (sq *SourceQueryIntegration) Close() error {
	return sq.client.Close()
}
