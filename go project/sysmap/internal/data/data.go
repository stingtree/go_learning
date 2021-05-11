package data

import (
	"context"
	"fmt"
	"sysmap/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"sysmap/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	client, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		fmt.Errorf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Errorf("failed creating schema resources: %v", err)
	}
	d := &Data{
		db: client,
	}

	return d, func() {
		fmt.Println("closing the data resources")
		if err := d.db.Close(); err != nil {
			fmt.Println(err)
		}
	}, nil
}
