package pgprovider

import (
	"bannerlord/config"
	"bannerlord/internal/services"
	"context"
	"github.com/go-pg/pg/v10"
)

type pgProvider struct {
	cfg config.Config
	db  *pg.DB
}

func New(cfg *config.Config) services.Storage {
	return &pgProvider{cfg: *cfg}
}

func (p *pgProvider) Connect() error {
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
