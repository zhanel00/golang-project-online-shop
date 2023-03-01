package Basic

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func Register(email, password string) {
	f, err := os.OpenFile("data/login_data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	ErrorHandler(err)

	defer f.Close()

	_, err = fmt.Fprintln(f, email, password)
	ErrorHandler(err)
}

func Authorize(email, password string) bool {
	var logins []string

	f, err := os.Open("data/login_data.txt")
	ErrorHandler(err)

	Scanner := bufio.NewScanner(f)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		logins = append(logins, Scanner.Text())
	}

	for index := 0; index < len(logins); index += 2 {
		if logins[index] == email {
			if logins[index+1] == password {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return false
}
