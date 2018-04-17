package main

/*
	车站系统
*/

import "fmt"
import (
	_ "vendors/mysql"
	"database/sql"
)

func main(){
	fmt.Println("hello world!!!!!")
	db, err := sql.Open("mysql", "root:920717@/mrzhang")

	if err == nil{
		fmt.Println(db)
	}
}
