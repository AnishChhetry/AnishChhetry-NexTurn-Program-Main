package main

import (
	"Blog_Management_System/dbconfig"
	"Blog_Management_System/handlers"
	"Blog_Management_System/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	dbconfig.Init()
	defer dbconfig.Close()

	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthenticationMiddleware)

	apiRouter := router.PathPrefix("/").Subrouter()
	apiRouter.HandleFunc("/blog/{id}", handlers.GetBlogByID).Methods("GET")
	apiRouter.HandleFunc("/blogs", handlers.GetAllBlogs).Methods("GET")
	apiRouter.HandleFunc("/blog/{id}", handlers.UpdateBlogByID).Methods("PUT")
	apiRouter.HandleFunc("/blog/{id}", handlers.DeleteBlogByID).Methods("DELETE")
	apiRouter.HandleFunc("/blog", handlers.CreateBlog).Methods("POST")

	apiRouter.Use(middleware.ValidateJSONMiddleware)

	fmt.Println("Server started at port 8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
