package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/inventory-app/backend/golang/common/conf"
	"github.com/inventory-app/backend/golang/common/utils"
	"github.com/inventory-app/backend/golang/models"
	"github.com/inventory-app/backend/golang/repository"
	"github.com/inventory-app/backend/golang/validaton"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
	"log"
)

func CreateNewProduct(product models.Product) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	if resp, err := validaton.ProductValidation(product); err != nil {
		return resp, err
	}
	//Connect database and store login data
	if userResp, err := repository.CreateProductRepository(product); err != nil {
		resp["statusCode"] = 500
		resp["message"] = "New Product not added. Please add proper details of product."
		return resp, err
	} else {
		return userResp, err
	}
}

func CreateNewCategory(category models.Category) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	if resp, err := validaton.CategoryValidation(category); err != nil {
		return resp, err
	}
	//Connect database and store login data
	if userResp, err := repository.CreateCategoryRepository(category); err != nil {
		resp["statusCode"] = 500
		resp["message"] = "New Product not added. Please add proper details of product."
		return resp, err
	} else {
		return userResp, err
	}
}

func GetCategory() (map[string]interface{}, error) {

	var resp = make(map[string]interface{})
	var category []models.Category
	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.CATEGORY_COLLECTION)
	if err := collection.Find(nil).All(&category); err != nil {
		log.Fatal("Error at Getting Category List:", err)
	}
	resp["statusCode"] = 201
	resp["category"] = category
	return resp, err
}

func OrderPaymentService(email string) (map[string]interface{}, error) {

	var resp = make(map[string]interface{})
	var cart []models.Cart
	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.CART_COLLECTION)
	if err := collection.Find(bson.M{"email": email}).All(&cart); err != nil {
		log.Fatal("Error at Getting Category List:", err)
	}

	items := []models.Items{}
	var totalPrice float64

	for _, value := range cart {
		item := models.Items{}
		for _, v := range value.Items {
			item.ItemName = v.ProductName
			item.ItemPrice = v.ProductPrice
			item.Quntity = v.Quantity
			totalPrice += v.ProductPrice * float64(v.Quantity)
			items = append(items, item)
		}
	}

	invoice := models.Invoice{
		cart[0].Name,
		cart[0].Email,
		items,
		totalPrice,
		"Your order will be placed with in 2-3 working days",
	}

	if resp, err = repository.OrderPaymentRepository(invoice); err != nil {
		return resp, err

	} else {
		resp["statusCode"] = 201
		resp["Message"] = "Thank you for the shopping!!!!"
		resp["Invoice"] = invoice
		return resp, err
	}
}

func AddProductToCart(userContext jwt.MapClaims, cartItems []models.ProductItems) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	if userResp, err := repository.AddProductToCartRepository(userContext, cartItems); err != nil {
		resp["message"] = userResp
		return resp, err
	} else {
		resp["message"] = userResp
		return userResp, nil
	}
}

func OrderList() (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	var category []models.Category
	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.ORDER_COLLECTION)
	if err := collection.Find(nil).All(&category); err != nil {
		log.Fatal("Error at Getting Order List", err)
	}
	resp["statusCode"] = 201
	resp["category"] = category
	return resp, err
}
