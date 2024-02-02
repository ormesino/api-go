package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ormesino/e-commerce/internal/database"
	"github.com/ormesino/e-commerce/internal/service"
	"github.com/ormesino/e-commerce/internal/webserver"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommerce")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	categoryDB := database.NewCategoryDB(db)
	productDB := database.NewProductDB(db)

	categoryService := service.NewCategoryService(*categoryDB)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/categories", webCategoryHandler.GetCategories)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/products", webProductHandler.GetProducts)
	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product/category/{id}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("💡 Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
