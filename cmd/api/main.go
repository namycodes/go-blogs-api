package main

import (
	"fmt"
	"log"
	"net/http"

	"com.namycodes/internal/authentication"
	"com.namycodes/internal/blog"
	"com.namycodes/internal/middleware"
	"com.namycodes/migrations"
	"com.namycodes/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	dotenverr := godotenv.Load()
	if dotenverr != nil {
		log.Fatal("Error loading .env file")
	}

	db:= config.ConnectDatabase()
	migrations.RunMigrations(db)

	router := gin.Default()
   
	blogRepository := blog.NewRepository(db)
	blogService := blog.NewService(blogRepository)
	blogHandler := blog.NewHandler(blogService)
    authRepository := authentication.NewRepository(db)
	authService := authentication.NewService(authRepository)
	authHandler := authentication.NewHandler(authService)

	blogs := router.Group("/api/v1/blogs")
	 blogs.Use(middleware.JwtAuth())
	{
		blogs.GET("/", blogHandler.GetAllBlogs)
		blogs.POST("/", blogHandler.CreateBlog)
		blogs.GET("/:id", blogHandler.GetBlogById)
		blogs.DELETE("/:id", blogHandler.DeleteBlogById)
		blogs.PATCH("/:id", blogHandler.UpdateBlog)
	}

	authentication := router.Group("/api/v1/auth")
	{
		authentication.POST("/register", authHandler.RegisterUser)
		authentication.POST("/login", authHandler.Login)
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