package main

import (
	"bookkeeping/router"
)

func main() {
	// if dao.ConnectPostgresDB() == false {
	// 	fmt.Println("cann't connect to postgresSQL Database")
	// 	return
	// }
	// defer dao.PostgresDB.Close()
	// if dao.InitDB() == false {
	// 	fmt.Println("cann't initialize postgresSQL Database")
	// 	return
	// }
	router := router.SetUp()
	router.Run()
}
