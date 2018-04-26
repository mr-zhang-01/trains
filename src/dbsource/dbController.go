package dbsource

import (
	"database/sql"

	"conf"
	"fmt"
)

var (
	//dbsource
	MySqlHanle *MySQL
	Errors error
)

func init(){
	MySqlHanle = &MySQL{}
	MySqlHanle.Start()
}

type MySQL struct{
	source *sql.DB
}

func(m *MySQL)Start()(){
	m.source, Errors = sql.Open(conf.DbConf["driver"], conf.DbConf["user"] + ":" + conf.DbConf["password"] + "@" + "/" + conf.DbConf["dbname"])
}

func(m *MySQL)Add(tableNmae string, data map[string]interface{})(){
	fmt.Println(tableNmae)

	fmt.Println(data["name"])
	fmt.Println(data["area_id"])
}

func(m *MySQL)Edit()(){

}

func(m *MySQL)Delete()(){

}

func(m *MySQL)Search()(){

}
