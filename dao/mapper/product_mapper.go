package mapper

import (
	"go-learning/dao"
	. "github.com/jinzhu/gorm"
)
var db *DB

func Save(product *Product)  {
	db = dao.GetDb()
	db.Save(product)
}

type Product struct {
	product_name string
	product_url string
	created_at string
}
