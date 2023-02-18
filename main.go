package main

import (
	"App-project/appsystem/Basic"
	"App-project/appsystem/Item"
	"fmt"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/mainpage.html")
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/searchform.html")
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/registration.html")
	})
	http.HandleFunc("/searchform", Item.PrintByName)
	http.HandleFunc("/items", Item.ReturnAllItems)
	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		Basic.Register(email, password)

		fmt.Fprintf(w, "Registration went successfully")
	})
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func main() {
	//var u User = User{}
	//var authorized bool = false
	//for true {
	//	fmt.Println("Hello, choose an option: ")
	//	fmt.Println("                         1. Register")
	//	fmt.Println("                         2. Authorize")
	//	fmt.Println("                         3. Search items by name")
	//	fmt.Println("                         4. Search items by price")
	//	fmt.Println("                         5. Search items by rating")
	//	fmt.Println("                         6. Give a rate for an item")
	//	var option string
	//	_, err := fmt.Scanf("%s\n", &option)
	//	Basic.ErrorHandler(err)
	//	if option == "1" {
	//		authorized = false
	//		var email string
	//		var password string
	//		email, password = Basic.AskCredentials()
	//		u.Email = email
	//		u.Password = password
	//		u.Register()
	//	} else if option == "2" {
	//		var email string
	//		var password string
	//		email, password = Basic.AskCredentials()
	//		u.Email = email
	//		u.Password = password
	//		if u.Authorization() {
	//			authorized = true
	//		}
	//	} else if option == "3" {
	//		var itemName string
	//		fmt.Println("Enter name of the item: ")
	//		_, err := fmt.Scanf("%s\n", &itemName)
	//		Basic.ErrorHandler(err)
	//		items, index := Item.SearchByName(itemName)
	//		Item.PrintItems([]Item.Item{items[index]})
	//	} else if option == "4" {
	//		var itemPrice1, itemPrice2 int
	//		fmt.Println("Enter price values between which the search will be done: ")
	//		_, err := fmt.Scanf("%d %d\n", &itemPrice1, &itemPrice2)
	//		Basic.ErrorHandler(err)
	//		items := Item.FilterByPrice(itemPrice1, itemPrice2)
	//		Item.PrintItems(items)
	//	} else if option == "5" {
	//		var rating float64
	//		fmt.Println("Enter rating value by which items will be filtered out: ")
	//		_, err := fmt.Scanf("%f\n", &rating)
	//		Basic.ErrorHandler(err)
	//		items := Item.FilterByRating(rating)
	//		Item.PrintItems(items)
	//	} else if option == "6" {
	//		if authorized == false {
	//			fmt.Println("You are not authorized, authorize first")
	//		} else {
	//			var itemName string
	//			var givenRate float64
	//			fmt.Println("Enter item name and your rate given to it on a 5-point scale: ")
	//			_, err := fmt.Scanf("%s %f\n", &itemName, &givenRate)
	//			Basic.ErrorHandler(err)
	//			u.Rate(itemName, givenRate)
	//		}
	//	} else if option == "7" {
	//		Basic.HandleRequest()
	//	} else {
	//		fmt.Println("There is no such function yet")
	//	}
	//}
	HandleRequest()
}
