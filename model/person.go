package model

type Person struct {
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

func CreatePerson(p Person) bool {
	return true
}
