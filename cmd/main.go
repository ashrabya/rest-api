package main

import (
	"fmt"
	"github.com/market/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
handler()


}

func handler (){
	r := mux.NewRouter()
	r.HandleFunc("/products",handlers.GetProductsHandler).Methods(http.MethodGet)
	r.HandleFunc("/products/create", handlers.CreateProductHandler).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", handlers.GetProductHandler).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", handlers.DeleteProductHandler).Methods(http.MethodDelete)
	r.HandleFunc("/products/{id}", handlers.UpdateProductHandler).Methods(http.MethodPut)

	http.ListenAndServe(":8080",r)
	fmt.Printf("Application running on port %v\n ", "8080:")
}