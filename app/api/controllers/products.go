package controllers

import (
	"log"
	"net/http"
	"store/domain"
	"store/infra/repository"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	productsRepository := repository.ProductsRepository{}

	products := productsRepository.ListAll()
	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "POST" {
		product := domain.Product{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
		}
		product.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Erro na conversão: ", err)
		}
		product.Amount, err = strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Erro na conversão: ", err)
		}

		productsRepository := repository.ProductsRepository{}
		productsRepository.CreateProduct(&product)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	productsRepository := repository.ProductsRepository{}
	productsRepository.Delete(ID)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	productRepository := repository.ProductsRepository{}
	product := productRepository.Get(ID)

	templates.ExecuteTemplate(w, "EditProduct", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "POST" {
		product := domain.Product{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
		}
		product.ID, err = strconv.Atoi(r.FormValue("ID"))
		if err != nil {
			log.Println("Erro na conversão: ", err)
		}
		product.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Erro na conversão: ", err)
		}
		product.Amount, err = strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Erro na conversão: ", err)
		}

		productsRepository := repository.ProductsRepository{}
		productsRepository.Update(&product)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
