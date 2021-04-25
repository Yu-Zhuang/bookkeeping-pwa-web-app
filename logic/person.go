package logic

import (
	"bookkeeping/dao"
	"bookkeeping/model"
)

func CreatePerson(p model.Person) bool {
	sql_statement := "SELECT id FROM person WHERE id=$1;"
	rows, err := dao.PostgresDB.Query(sql_statement, p.ID)
	if err != nil {
		return false
	}
	defer rows.Close()
	for rows.Next() {
		return false
	}

	p.Password = Sha1Hash(p.Password)
	sql_statement = "INSERT INTO person (id, name, password, email) VALUES ($1, $2, $3, $4);"
	_, err = dao.PostgresDB.Exec(sql_statement, p.ID, p.Name, p.Password, p.Email)
	if err != nil {
		return false
	}
	return true
}

func HasAccount(p model.Person) bool {
	p.Password = Sha1Hash(p.Password)
	sql_statement := "SELECT id, password FROM person WHERE id=$1 AND password=$2;"
	rows, err := dao.PostgresDB.Query(sql_statement, p.ID, p.Password)
	if err != nil {
		return false
	}
	for rows.Next() {
		return true
	}
	return false
}
