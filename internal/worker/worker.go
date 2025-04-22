package worker

import (
	"context"
	"fmt"
	"github.com/floriansw/go-hll-rcon/rconv2"
	"github.com/floriansw/go-hll-rcon/rconv2/api"
	"github.com/floriansw/hll-geofences/internal"
	"log/slog"
	"slices"
	"time"
)

type worker struct {
	pool               *rconv2.ConnectionPool
	l                  *slog.Logger
	axisFences         []internal.Fence
	alliesFences       []internal.Fence
	punishAfterSeconds time.Duration

	sessionTicker *time.Ticker
	playerTicker  *time.Ticker
	punishTicker  *time.Ticker

	current        *api.GetSessionResponse
	outsidePlayers map[string]time.Time
}

func NewWorker(l *slog.Logger, pool *rconv2.ConnectionPool, c internal.Server) *worker {
	punishAfterSeconds := 10
	if c.PunishAfterSeconds != nil {
		punishAfterSeconds = *c.PunishAfterSeconds
	}
	return &worker{
		l:                  l,
		pool:               pool,
		punishAfterSeconds: time.Duration(punishAfterSeconds) * time.Second,
		axisFences:         c.AxisFence,
		alliesFences:       c.AlliesFence,

		sessionTicker:  time.NewTicker(5 * time.Second),
		playerTicker:   time.NewTicker(500 * time.Millisecond),
		punishTicker:   time.NewTicker(time.Second),
		outsidePlayers: map[string]time.Time{},
	}
}

func (w *worker) Run(ctx context.Context) {
	if err := w.populateSession(ctx); err != nil {
		w.l.Error("fetch-session", "error", err)
		return
	}

	go w.pollSession(ctx)
	go w.pollPlayers(ctx)
	go w.punishPlayers(ctx)
}

func (w *worker) populateSession(ctx context.Context) error {
	return w.pool.WithConnection(ctx, func(c *rconv2.Connection) error {
		si, err := c.SessionInfo(ctx)
		if err != nil {
			return err
		}
		w.current = si
		return nil
	})
}

func (w *worker) punishPlayers(ctx context.Context) {
	for {
		select {
		case <-w.punishTicker.C:
			for id, t := range w.outsidePlayers {
				println(id, t.String(), time.Since(t).String(), w.punishAfterSeconds.String())
				if time.Since(t) > w.punishAfterSeconds {
					w.punishPlayer(ctx, id)
					delete(w.outsidePlayers, id)
				}
			}
		}
	}
}

func (w *worker) punishPlayer(ctx context.Context, id string) {
	err := w.pool.WithConnection(ctx, func(c *rconv2.Connection) error {
		return c.PunishPlayer(ctx, id, fmt.Sprintf("%s outside the playarea", w.punishAfterSeconds.String()))
	})
	if err != nil {
		w.l.Error("poll-session", "error", err)
	}
}

func (w *worker) pollSession(ctx context.Context) {
	for {
		select {
		case <-w.sessionTicker.C:
			if err := w.populateSession(ctx); err != nil {
				w.l.Error("poll-session", "error", err)
			}
		}
	}
}

func (w *worker) pollPlayers(ctx context.Context) {
	for {
		select {
		case <-w.playerTicker.C:
			err := w.pool.WithConnection(ctx, func(c *rconv2.Connection) error {
				players, err := c.Players(ctx)
				if err != nil {
					return err
				}
				for _, player := range players.Players {
					go w.checkPlayer(ctx, player)
				}
				return nil
			})
			if err != nil {
				w.l.Error("poll-session", "error", err)
			}
		}
	}
}

var (
	alliedTeams = []api.PlayerTeam{
		api.PlayerTeamB8a,
		api.PlayerTeamDak,
		api.PlayerTeamGb,
		api.PlayerTeamRus,
		api.PlayerTeamUs,
	}
)

func (w *worker) checkPlayer(ctx context.Context, p api.GetPlayerResponse) {
	if !p.Position.IsSpawned() {
		return
	}
	g := p.Position.Grid(w.current)
	var fences []internal.Fence
	if slices.Contains(alliedTeams, p.Team) {
		fences = w.alliesFences
	} else {
		fences = w.axisFences
	}
	for _, f := range fences {
		if f.Includes(g) {
			delete(w.outsidePlayers, p.Id)
			return
		}
	}
	if _, ok := w.outsidePlayers[p.Id]; ok {
		return
	}
	w.outsidePlayers[p.Id] = time.Now()
	w.l.Info("player-outside-fence", "player", p.Name, "grid", g)
	err := w.pool.WithConnection(ctx, func(c *rconv2.Connection) error {
		return c.MessagePlayer(ctx, p.Name, fmt.Sprintf("You are outside of the designated playarea! Please go back to the battlefield immediately.\n\nYou will be punished in %s", w.punishAfterSeconds.String()))
	})
	if err != nil {
		w.l.Error("message-player-outside-fence", "player", p.Name, "grid", g, "error", err)
	}
}
