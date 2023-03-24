package books

import (
	"sharath/database"

	"strconv"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Books struct {
	gorm.Model
	Name          string `json:"name"`
	Author        string `json:"author"`
	YearOfPublish int    `json:"year"`
	Genre         string `json:"genre"`
}


func AddNewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Books)
	if err := c.BodyParser(&book); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse JSON"})
		return
	}

	var existingBook Books
	db.Where("name = ?",book.Name).First(&existingBook)
	if existingBook.ID != 0 {
		id:=strconv.FormatUint(uint64(existingBook.ID), 10)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Book With This Name Already Exists with Id: "+id})
		return
	}else{
		db.Create(&book)
	    c.JSON(book)
	}
}



func GetAllBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Books
	db.Find(&books)
	c.JSON(books)
}



func GetBookById(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Books
	db.Find(&book, id)
	if book.Name == "" {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unable to find book with given Id: " + id})
		return
	}
	c.JSON(book)
}



func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	book := new(Books)

	db.First(&book, id)
	if book.Name == "" {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unable to find book with given Id: " + id})
		return
	}

	if err := c.BodyParser(&book); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse JSON"})
		return
	}

	db.Save(&book)
	c.JSON(book)
}



func DeleteBook(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var book Books
	db.First(&book, id)
	if book.Name == "" {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book Not Found With This Id: " + id})
		return
	}
	db.Delete(&book)
	c.Status(fiber.StatusOK).JSON(fiber.Map{"sucess": "Book with Id: " + id + " Deleted Sucessfully"})
}
