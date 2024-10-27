package sourcequery

import (
	"context"
	"fmt"
	"steamserverlauncher/iterate"

	"github.com/rumblefrog/go-a2s"
)

type SourceQueryPlayer struct {
	Name, Score string
}

func (sq *SourceQueryIntegration) CurrentPlayers(ctx context.Context) ([]SourceQueryPlayer, error) {
	result, err := sq.client.QueryPlayer()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch current players: %w", err)
	}

	return iterate.Map(result.Players, func(player *a2s.Player) SourceQueryPlayer {
		return SourceQueryPlayer{Name: player.Name, Score: fmt.Sprint(player.Score)}
	}), nil
}
