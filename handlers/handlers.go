package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest-api/entity"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	data, err := entity.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(data)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	// Read product ID
	productID := mux.Vars(r)["id"]
	product, err := entity.GetProduct(productID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseData, err := json.Marshal(product)
	if err != nil {
		// Check if it is No product error or any other error
		if errors.Is(err, entity.ErrNoProduct) {
			// Write Header if no related product found.
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	// Write body with found product
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(responseData)
}

// DeleteProductHandler deletes the product with given ID.
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// Read product ID
	productID := mux.Vars(r)["id"]
	err := entity.DeleteProduct(productID)
	if err != nil {
		// Check if it is No product error or any other error
		if errors.Is(err, entity.ErrNoProduct) {
			// Write Header if no related product found.
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	// Write Header with Accepted Status (done operation)
	w.WriteHeader(http.StatusAccepted)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Check if data is proper JSON (data validation)
	var product []entity.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Invalid Data Format"))
		return
	}
	err = entity.AddProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Added New Product"))
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

	productID := mux.Vars(r)["id"]
	err := entity.DeleteProduct(productID)
	if err != nil {
		if errors.Is(err, entity.ErrNoProduct) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	// Read incoming JSON from request body
	data, err := ioutil.ReadAll(r.Body)
	// If no body is associated return with StatusBadRequest
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product []entity.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Invalid Data Format"))
		return
	}

	err = entity.AddProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
