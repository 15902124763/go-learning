package mapper

import (
	"go-learning/dao"
	. "github.com/jinzhu/gorm"
)
var db *DB

func Save(product *T_product) {
	db = dao.GetDb()
	db.Create(product)
}

func Del(id int)  {
	db := dao.GetDb()
	db.Delete(T_product{},"id=?", id)
}

func Update(product *T_product) {
	db = dao.GetDb()
	//db.Model(&T_product{}).Where("id=?", product.Id).Update(
	//	"product_name", product.Product_name,
	//	"product_url",product.Product_url,
	//	"updated_at",product.Updated_at,
	//	"updated_by",product.Updated_by,)
	db.Exec("UPDATE t_product SET product_name=?,product_url=?,updated_at=?,updated_by=? WHERE id=?",
		product.Product_name,product.Product_url,product.Updated_at,product.Updated_by, product.Id)
}

func Select(id int) T_product  {
	db := dao.GetDb()
	var t_p T_product
	db.Raw("SELECT * FROM t_product WHERE id = ?", id).Scan(&t_p)
	return t_p
}

func SelectAll() []T_product  {
	db := dao.GetDb()
	//var pList []T_product
	var pList []T_product
	db.Find(&pList)
	return pList
}

type (
	T_product struct {
		Id int
		Product_name string
		Product_url string
		Created_at string
		Created_by string
		Updated_at string
		Updated_by string
	}
)