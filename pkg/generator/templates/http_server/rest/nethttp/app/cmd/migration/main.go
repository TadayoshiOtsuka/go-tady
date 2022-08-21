package main

import (
	"GO_TADY_PACKAGE_NAME/app/individual/migration"
)

func main() {
	migration.NewMigrateHandler().Migrate()
}
