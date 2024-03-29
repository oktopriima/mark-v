package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/oktopriima/mark-v/configurations"
)

func MssSqlConnection(cfg configurations.Config) (*gorm.DB, error) {
	dbUser := cfg.GetString(`mssql.user`)
	dbPass := cfg.GetString(`mssql.pass`)
	dbName := cfg.GetString(`mssql.database`)
	dbHost := cfg.GetString(`mssql.address`)
	dbPort := cfg.GetString(`mssql.port`)

	db, err := gorm.Open("mssql", "sqlserver://"+dbUser+":"+dbPass+"@"+dbHost+":"+dbPort+"?database="+dbName+"")
	if err != nil {
		return nil, err
	}
	return db, nil
}
