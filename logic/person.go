package logic

import (
	"bookkeeping/dao"
	"bookkeeping/model"
	"fmt"
)

func CreatePerson(p model.Person) bool {
	sql_statement := "SELECT COUNT(id) FROM person WHERE id=$1;"
	rows, err := dao.PostgresDB.Query(sql_statement, p.ID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			fmt.Println(err.Error())
		}
	}
	if count > 0 {
		fmt.Println("has account : ", p.ID)
		return false
	}
	p.Password = Sha1Hash(p.Password)
	sql_statement = "INSERT INTO person (id, name, password, email) VALUES ($1, $2, $3, $4);"
	_, err = dao.PostgresDB.Exec(sql_statement, p.ID, p.Name, p.Password, p.Email)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func HasAccount(p model.Person) bool {
	p.Password = Sha1Hash(p.Password)
	sql_statement := "SELECT COUNT(id) FROM person WHERE id=$1 AND password=$2;"
	rows, err := dao.PostgresDB.Query(sql_statement, p.ID, p.Password)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			fmt.Println(err.Error())
		}
	}
	if count == 1 {
		return true
	}
	fmt.Println("has no account : ", p.ID)
	return false
}
