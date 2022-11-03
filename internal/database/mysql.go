package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"init_project/internal/utils"
)

const (
	MysqlDriver = "mysql"
	PGDriver    = "postgres"
)

type DataSource struct {
	*sqlx.DB
	config string
}

type DataStore struct {
	DataSource
}

func NewDB(driver, configUrl string) *DataStore {
	ds := DataSource{
		DB:     open(driver, configUrl),
		config: configUrl,
	}
	return &DataStore{
		DataSource: ds,
	}
}

func open(driver, config string) *sqlx.DB {
	if driver != MysqlDriver && driver != PGDriver {
		utils.Logger.Work.Errorf("not supported driver: %s", driver)
	}
	utils.Logger.Work.Debug("connect to %s", config)
	db, err := sqlx.Connect(driver, config)
	if err != nil {
		utils.Logger.Work.Errorf("database connection failed. err: %s", err)
	}
	switch driver {
	case "mysql":
		db.SetMaxIdleConns(5)
	case "postgres":
		db.SetMaxIdleConns(5)
	}
	return db
}
