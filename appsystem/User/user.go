package User

import (
	"App-project/appsystem/Basic"
	"App-project/appsystem/Item"
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type User struct {
	firstName string
	lastName  string
	Email     string
	Password  string
}

func (u *User) Register() {
	f, err := os.OpenFile("data/login_data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	Basic.ErrorHandler(err)

	defer f.Close()

	_, err = fmt.Fprintln(f, u.Email, u.Password)
	Basic.ErrorHandler(err)
}

func (u *User) Authorization() bool {
	email := u.Email
	password := u.Password
	var logins []string

	f, err := os.Open("data/login_data.txt")
	Basic.ErrorHandler(err)

	Scanner := bufio.NewScanner(f)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		logins = append(logins, Scanner.Text())
	}

	for index := 0; index < len(logins); index += 2 {
		if logins[index] == email {
			if logins[index+1] == password {
				fmt.Println("You are authorized")
				return true
			} else {
				fmt.Println("Incorrect credentials, try again")
				return false
			}
		} else {
			fmt.Println("Such email address does not exist, try again")
			return false
		}
	}
	return false
}

func (u *User) Rate(itemName string, rate float64) {
	items, index := Item.SearchByName(itemName)
	var newRate = (items[index].Rating*float64(items[index].Voted) + rate) / float64(items[index].Voted+1)
	items[index].Rating = math.Round(newRate*100) / 100
	items[index].Voted += 1

	file, err := json.MarshalIndent(&items, "", " ")
	Basic.ErrorHandler(err)
	_ = os.WriteFile("data/items.json", file, 0644)
}
