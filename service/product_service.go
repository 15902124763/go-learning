package service

import (
	"go-learning/dao/mapper"
	"time"
)

func Save(bo *ProductBO) {
	do := new(mapper.T_product)
	do.Product_name = bo.ProductName
	do.Product_url = bo.ProductUrl
	do.Created_at = time.Now().Format("2006-01-02 15:04:05")
	do.Created_by = "sys"
	do.Updated_at = time.Now().Format("2006-01-02 15:04:05")
	mapper.Save(do)
}

func Del(id int)  {
	mapper.Del(id)
}

func Update(bo *ProductBO) {
	do := new(mapper.T_product)
	do.Id = bo.Id
	do.Product_name = bo.ProductName
	do.Product_url = bo.ProductUrl
	do.Updated_at = time.Now().Format("2006-01-02 15:04:05")
	do.Updated_by = "sys"
	mapper.Update(do)
}

func Select(id int) mapper.T_product {
	return mapper.Select(id)
}

// 成员变量首字母要大写
type (
	ProductBO struct {
		Id  int `json:"id"`
		ProductName string  `json:"productName"`
		ProductUrl string `json:"productUrl"`
	}
)
