package main

import (
	"fmt"
	"net/http"

	"com.namycodes/internal/blog"
	"com.namycodes/migrations"
	"com.namycodes/pkg/config"
	"github.com/gin-gonic/gin"
)

func main(){
	db:= config.ConnectDatabase()
	migrations.RunMigrations(db)

	router := gin.Default()
   
	blogRepository := blog.NewRepository(db)
	blogService := blog.NewService(blogRepository)
	blogHandler := blog.NewHandler(blogService)

	blogs := router.Group("/api/v1/blogs")
	{
		blogs.GET("/", blogHandler.GetAllBlogs)
		blogs.POST("/", blogHandler.CreateBlog)
		blogs.GET("/:id", blogHandler.GetBlogById)
		blogs.DELETE("/:id", blogHandler.DeleteBlogById)
	}

	

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	fmt.Println("Server Started Successfully")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	
}