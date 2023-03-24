package database

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	DBConn *gorm.DB
)

const DNS string = "admin:adminbooks@tcp(database-1.csgjrk0tttja.ap-south-1.rds.amazonaws.com:3306)/booksDB?parseTime=true"