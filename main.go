package main

import(
	"fmt"
	"sharath/books"
	"sharath/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)


func initDatabase(){
	var err error
	database.DBConn , err = gorm.Open("mysql", database.DNS)
	if err!=nil{
		panic("failed to connect to the database")
	}
	fmt.Println("Connection opened to database sucessfully")
	database.DBConn.AutoMigrate(&books.Books{})
	fmt.Println("Database Migrated")
}

func setupRoutes(app *fiber.App){
	app.Post("/api/v1/books/upload",books.AddNewBook)
	app.Get("/api/v1/books/:id",books.GetBookById)
	app.Get("/api/v1/books",books.GetAllBooks)
	app.Put("/api/v1/books/:id",books.UpdateBook)
	app.Delete("/api/v1/books/:id",books.DeleteBook)
}


func main(){
	app:=fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(8000)
	defer database.DBConn.Close()
}