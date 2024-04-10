package pgmanager

import (
	"bannerlord/internal/config"
	"bannerlord/internal/services/storage"
	"context"
	"github.com/go-pg/pg/v10"
)

type pgManager struct {
	cfg config.Config
	db  *pg.DB
}

func New(cfg *config.Config) storage.Service {
	return &pgManager{cfg: *cfg}
}

func (p *pgManager) Connect() error {
	p.db = pg.Connect(&pg.Options{
		Addr:            p.cfg.Host + ":" + p.cfg.Port,
		User:            p.cfg.UserName,
		Password:        p.cfg.Password,
		Database:        p.cfg.DBName,
		MaxRetries:      3,
		MaxRetryBackoff: 3,
	})

	if err := p.db.Ping(context.Background()); err != nil {

		return err
	}

	return nil
}
