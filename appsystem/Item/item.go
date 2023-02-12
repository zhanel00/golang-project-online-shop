package Item

import (
	"App-project/appsystem/Basic"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Item struct {
	Id     int     `json:"Id"`
	Name   string  `json:"Name"`
	Price  int     `json:"Price"`
	Rating float64 `json:"Rating"`
	Voted  int     `json:"Voted"`
}

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
