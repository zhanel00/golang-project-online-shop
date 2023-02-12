package Basic

import (
	"fmt"
	"log"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func AskCredentials() (string, string) {
	var email, password string
	fmt.Println("Enter email: ")
	_, err := fmt.Scanf("%s\n", &email)
	ErrorHandler(err)
	fmt.Println("Enter password: ")
	_, err = fmt.Scanf("%s\n", &password)
	ErrorHandler(err)
	return email, password
}
