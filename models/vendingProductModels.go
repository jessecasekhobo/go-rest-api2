package models

import (
	"database/sql"

	"go-rest-api/entities"
)

//ProductModel model
type ProductModel struct {
	Db *sql.DB
}

//FindAll
func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("[Vending_MachineDB].[dbo].[GetVendingMachineProducts]")
	if err != nil {
		return nil, err
	} else {
		products := []entities.Product{}
		for rows.Next() {
			var prodid int
			var prodname string
			var prodpicture string
			var price int
			err2 := rows.Scan(&prodid, &prodname, &prodpicture, &price)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{prodid, prodname, prodpicture, price}
				products = append(products, product)
			}
		}

		return products, nil
	}
}
