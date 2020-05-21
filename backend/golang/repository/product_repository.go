package repository

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/inventory-app/backend/golang/common/conf"
	"github.com/inventory-app/backend/golang/common/utils"
	"github.com/inventory-app/backend/golang/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func CreateProductRepository(product models.Product) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})

	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.PRODUCT_COLLECTION)
	if err := collection.Insert(product); err != nil {
		log.Fatal("Error at creating new product :", err)
	}
	resp["statusCode"] = 201
	resp["message"] = "New product added to Database"
	return resp, err
}

func CreateCategoryRepository(category models.Category) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})

	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.CATEGORY_COLLECTION)
	//category.Id  = bson.NewObjectId()
	if err := collection.Insert(category); err != nil {
		log.Fatal("Error at creating new category :", err)
	}
	resp["statusCode"] = 201
	resp["message"] = "New category added to Database"
	return resp, err
}

func AddProductToCartRepository(userContext jwt.MapClaims, cartItems []models.ProductItems) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})
	var user models.User
	email := userContext["email"]
	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()

	if err := session.DB(conf.Database).C(conf.UserCollecion).Find(bson.M{"email": email}).One(&user); err != nil {
		resp["statusCode"] = 500
		resp["message"] = "Un Authorised User!!!"

		return resp, err
	} else {
		cart := models.Cart{
			Name:  user.Name,
			Email: user.Email,
			Items: cartItems,
		}

		collection := session.DB(conf.Database).C(conf.CART_COLLECTION)
		if err := collection.Insert(cart); err != nil {
			log.Fatal("Error at creating new category :", err)
		}
		resp["statusCode"] = 201
		resp["message"] = "New category added to Database"
		return resp, nil
	}
}

func OrderPaymentRepository(invoice models.Invoice) (map[string]interface{}, error) {
	var resp = make(map[string]interface{})

	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.PAYMENT_COLLECTION)
	if err := collection.Insert(invoice); err != nil {
		log.Fatal("Error at placing payment. Please dont refresh the page:", err)
		resp["statusCode"] = 201
		resp["message"] = "Your Payment Failed!!!. Please don't refresh th page."
	}
	return resp, err
}
