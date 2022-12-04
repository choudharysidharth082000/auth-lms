package validate

// import (
// 	"fmt"
// )

// func LoginValidate(email string, password string) bool {
// 	fmt.Println("Login Validate : ", email, password)
// 	//password length
// 	lengthEmail := len(email)
// 	lengthPassword := len(password)
// 	if lengthEmail < 4 {
// 		return false
// 	}
// 	if lengthPassword < 8 {
// 		return false
// 	}
// 	return true
// }

// signup validations =
func SignupValidate(email string, password string, username string, firstname string, lastname string, phone string) bool {
	if len(email) < 8 {
		return false
	}
	if len(password) < 8 {
		return false
	}
	if len(username) < 4 {
		return false
	}
	if len(firstname) < 4 {
		return false
	}
	if len(lastname) < 4 {
		return false
	}
	if len(phone) < 10 {
		return false
	}
	return true
}
