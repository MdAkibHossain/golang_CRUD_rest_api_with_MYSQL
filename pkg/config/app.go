package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:bagdha1122@tcp(127.0.0.1:3306)/simplerest")
	//d, err := gorm.Open("mysql", "root:bagdha1122/simplerest?charset=utf8&parsetime=True&loc=Tocal")
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
