package infra

import (
	"ddd_demo/internal/conf"
	"ddd_demo/internal/infra/ent"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"time"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db *ent.Client
}

func NewData(c *conf.Conf) (*Data, func(), error) {
	drv, err := sql.Open(c.Data.Database.Driver, c.Data.Database.Dsn)
	if err != nil {
		return nil, nil, err
	}
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	d := &Data{
		db: ent.NewClient(ent.Driver(drv)),
	}
	return d, func() {
		d.db.Close()
	}, nil
}
