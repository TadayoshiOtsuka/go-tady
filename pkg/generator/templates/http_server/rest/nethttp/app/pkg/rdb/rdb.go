package rdb

import (
	"fmt"
	"log"

	"GO_TADY_PACKAGE_NAME/app/pkg/config"
	"GO_TADY_PACKAGE_NAME/app/pkg/rdb/orm/ent"

	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	Client *ent.Client
}

func NewSQLHandler() *SQLHandler {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.AppConfig.DB.UserName,
		config.AppConfig.DB.Password,
		config.AppConfig.DB.Host,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.DBName,
	)

	c, err := ent.Open(config.AppConfig.DB.Driver, dsn)
	if err != nil {
		log.Panicf("DB Connection Error:%v", err)
	}

	return &SQLHandler{Client: c}
}
