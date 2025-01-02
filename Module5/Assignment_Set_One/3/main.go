package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func AddProduct(id int, name string, priceInput string, stock int) error {
	price, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		return fmt.Errorf("Invalid price: %v", err)
	}

	if stock < 0 {
		return errors.New("Stock cannot be negative")
	}

	inventory = append(inventory, Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	})
	return nil
}

func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("Stock cannot be negative")
	}

	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}

	return errors.New("Product not found")
}

func SearchProduct(search interface{}) (*Product, error) {
	switch v := search.(type) {
	case int:
		for _, product := range inventory {
			if product.ID == v {
				return &product, nil
			}
		}
	case string:
		for _, product := range inventory {
			if product.Name == v {
				return &product, nil
			}
		}
	default:
		return nil, errors.New("Invalid search type")
	}

	return nil, errors.New("Product not found")
}
func DisplayInventory() {
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inventory {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func SortInventory(by string) error {
	switch by {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		return errors.New("Invalid sort option")
	}

	return nil
}

func main() {
	_ = AddProduct(1, "Shoes", "2599", 54)
	_ = AddProduct(2, "Hat", "799", 43)
	_ = AddProduct(3, "Shirt", "999", 64)
	_ = AddProduct(4, "Trouser", "1999", 40)

	fmt.Println("Initial Inventory:")
	DisplayInventory()

	fmt.Println("\nUpdating stock:")
	_ = UpdateStock(1, 70)
	DisplayInventory()

	fmt.Println("\nSorting by price:")
	_ = SortInventory("price")
	DisplayInventory()

	fmt.Println("\nSorting by stock:")
	_ = SortInventory("stock")
	DisplayInventory()

	fmt.Println("\nSearching for product by name (Shoes):")
	product, err := SearchProduct("Shoes")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found: %+v\n", *product)
	}

	fmt.Println("\nSearching for product by ID (4):")
	product, err = SearchProduct(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found: %+v\n", *product)
	}
}
