package dbsource

import (
	"database/sql"

	"conf"
)

var (
	MySqlHanle *MySQL
)

type MySQL struct{
	source *sql.DB
}

func(m *MySQL)Start()(){
	m.source = sql.Open()
}
