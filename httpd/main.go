package main

import (
	"fmt"

	"go-rest-api/entities"
	"go-rest-api/models"
	"go-rest-api/platform/datafeed"

	"go-rest-api/config"
)

//Application main execution point
func main() {
	fmt.Println("Hello World!")
	listener()
	//DB Connection string

}

//API Listener (Note: r.Run() keeps the listener active)
func listener() {
	/*r := gin.Default()

			r.GET("/ping", handler.PingGet())
			r.Run()
			//	prod := []entities.Product{}
			/*var i = 0
	         prod, nil := getinventoryItems()
	        for i < len(prod) {
		    println(prod[i].Price, nil)
		    i++


		}*/

	feed := datafeed.New()
	fmt.Println(feed)

	//feed.Add(datafeed.Item{"Helow", " fff"})
	feed.Add(datafeed.Item{"hello", "gg"})
	fmt.Println(feed)
}

//Used to get data from MsSQL database
func getinventoryItems() ([]entities.Product, error) {
	db, err := config.GetMsSQLDB() //opens SQL connection only
	products := []entities.Product{}
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Products: ", len(products))
		}

	}
	productModel := models.ProductModel{
		Db: db,
	}
	products, nil := productModel.FindAll()
	return products, nil
}
