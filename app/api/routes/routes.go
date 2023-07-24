package routes

import (
	"net/http"
	"store/app/api/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new_product", controllers.NewProduct)
	http.HandleFunc("/insert_product", controllers.Insert)
	http.HandleFunc("/edit_product", controllers.Edit)
	http.HandleFunc("/update_product", controllers.Update)
	http.HandleFunc("/delete_product", controllers.Delete)
}
