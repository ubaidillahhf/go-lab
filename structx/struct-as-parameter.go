package structx

import "fmt"

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func DisplayUser(user Student) (result string) {
	result = fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return
}
