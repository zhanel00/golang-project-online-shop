package Item

import (
	"App-project/appsystem/Basic"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

type Item struct {
	Id     int     `json:"Id"`
	Name   string  `json:"Name"`
	Price  int     `json:"Price"`
	Rating float64 `json:"Rating"`
	Voted  int     `json:"Voted"`
}

var items []Item

func FilterByPrice(price1 int, price2 int) []Item {
	var items []Item
	f, err := os.Open("data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteArray, _ := io.ReadAll(f)
	json.Unmarshal(byteArray, &items)

	var filtered []Item

	for i := range items {
		if items[i].Price >= price1 && items[i].Price <= price2 {
			filtered = append(filtered, items[i])
		}
	}
	return filtered
}

func FilterByRating(rating float64) []Item {
	var items []Item
	f, err := os.Open("data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteArray, _ := io.ReadAll(f)
	json.Unmarshal(byteArray, &items)

	var filtered []Item

	for i := range items {
		if items[i].Rating >= rating {
			filtered = append(filtered, items[i])
		}
	}
	return filtered
}

func SearchByName(name string) (items []Item, index int) {
	f, err := os.Open("data/items.json")
	Basic.ErrorHandler(err)
	defer f.Close()

	byteArray, _ := io.ReadAll(f)
	json.Unmarshal(byteArray, &items)

	for i := range items {
		if name == items[i].Name {
			return items, i
		}
	}
	return items, -1
}

func GetItems() {
	f, err := os.Open("data/items.json")
	Basic.ErrorHandler(err)
	defer f.Close()

	byteArray, _ := io.ReadAll(f)
	json.Unmarshal(byteArray, &items)
}

func PrintItems(items []Item) {
	for _, item := range items {
		fmt.Println("  id: ", item.Id)
		fmt.Println("  name: ", item.Name)
		fmt.Println("  price: ", item.Price)
		fmt.Println("  rating: ", item.Rating)
		fmt.Println("  number of votes: ", item.Voted)
		fmt.Println("-----------------------")
	}
}

func ReturnAllItems(w http.ResponseWriter, r *http.Request) {
	GetItems()
	for _, item := range items {
		fmt.Fprintln(w, "  id: ", item.Id)
		fmt.Fprintln(w, "  name: ", item.Name)
		fmt.Fprintln(w, "  price: ", item.Price)
		fmt.Fprintln(w, "  rating: ", item.Rating)
		fmt.Fprintln(w, "  number of votes: ", item.Voted)
		fmt.Fprintln(w, "-----------------------")
	}
}

func PrintByName(w http.ResponseWriter, r *http.Request) {
	GetItems()
	fmt.Println("Got in here")
	name := r.FormValue("itemname")
	fmt.Println(name)
	for _, item := range items {
		if item.Name == name {
			fmt.Fprintf(w, "Name: %s\nId: %d\nPrice: %d\nRating: %f\nVoted: %d\n-----------------", item.Name, item.Id, item.Price, item.Rating, item.Voted)
		}
	}
}

func PrintByPrice(w http.ResponseWriter, r *http.Request) {
	GetItems()
	fmt.Println()
	priceLower, err := strconv.Atoi(r.FormValue("pricevalue1"))
	Basic.ErrorHandler(err)
	priceHigher, err := strconv.Atoi(r.FormValue("pricevalue2"))
	Basic.ErrorHandler(err)
	for _, item := range items {
		if item.Price >= priceLower && item.Price <= priceHigher {
			fmt.Fprintf(w, "Name: %s\nId: %d\nPrice: %d\nRating: %f\nVoted: %d\n-----------------\n", item.Name, item.Id, item.Price, item.Rating, item.Voted)
		}
	}
}

func Rate(itemName string, rate float64) {
	items, index := SearchByName(itemName)
	var newRate = (items[index].Rating*float64(items[index].Voted) + rate) / float64(items[index].Voted+1)
	items[index].Rating = math.Round(newRate*100) / 100
	items[index].Voted += 1

	file, err := json.MarshalIndent(&items, "", " ")
	Basic.ErrorHandler(err)
	_ = os.WriteFile("data/items.json", file, 0644)
}

func PostRating(w http.ResponseWriter, r *http.Request) {
	GetItems()
	fmt.Println()
	name := r.FormValue("itemname")
	rating, err := strconv.ParseFloat(r.FormValue("itemrating"), 64)
	Basic.ErrorHandler(err)
	Rate(name, rating)
	fmt.Fprintf(w, "Thank you, your rating has been posted!")
}
