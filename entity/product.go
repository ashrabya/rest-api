package entity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

var ErrNoProduct = errors.New("no product found")

func GetProducts() ([]byte, error) {
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetProduct(id string) (Product, error) {
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return Product{}, err
	}
	//read product
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return Product{}, err
	}
	//iterate through product array
	for i := 0; i < len(products); i++ {
		//if we find one product with the given ID
		if products[i].ID == id {
			return products[i], nil
		}
	}
	return Product{}, err
}
func DeleteProduct(id string) error {
	//Read json file
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}
	//iterate through product array
	for i := 0; i < len(products); i++ {
		if products[i].ID == id {
			products = removeElement(products, i)
			updatedData, err := json.Marshal(products)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNoProduct
}

func AddProduct(product []Product) error {
	var products []Product
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}

	for _, product := range products {
		products = append(products, product)
	}

	updatedData, err := json.Marshal(products)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./data/data.jscon", updatedData, os.ModePerm)
	return nil

}
func removeElement(arr []Product, index int) []Product {
	ret := make([]Product, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}
