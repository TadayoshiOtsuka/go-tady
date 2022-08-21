package migration

import (
	"context"
	"log"

	"GO_TADY_PACKAGE_NAME/app/pkg/rdb"
	"GO_TADY_PACKAGE_NAME/app/pkg/rdb/orm/ent"
	"GO_TADY_PACKAGE_NAME/app/pkg/rdb/orm/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

type MigrateHandler struct {
	client *ent.Client
}

func NewMigrateHandler() *MigrateHandler {
	c := rdb.NewSQLHandler().Client

	return &MigrateHandler{client: c}
}

func (h *MigrateHandler) Migrate() {
	err := h.client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("Migrate Error: %v", err)
	}
}
