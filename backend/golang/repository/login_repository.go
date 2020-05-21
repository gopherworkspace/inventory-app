package repository

import (
	"fmt"
	"github.com/inventory-app/backend/golang/common/conf"
	"github.com/inventory-app/backend/golang/common/utils"
	"github.com/inventory-app/backend/golang/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
Inside Repository perform database operations
All services database logic's are written in repository files.
*/

func RegistrationRepository(registrationData models.User) (map[string]interface{}, error) {
	var user models.User
	var resp = make(map[string]interface{})
	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.UserCollecion)

	if err := collection.Find(bson.M{"email": registrationData.Email}).One(&user); err != nil {
		fmt.Println(registrationData.Email)

		if registrationData.Email == "admin@gmail.com" {
			registrationData.Role = "ADMIN"
		}

		if err := collection.Insert(registrationData); err != nil {
			panic(err)
		}
		resp["statusCode"] = 201
		resp["message"] = "New user created"

		return resp, err
	} else {
		resp["statusCode"] = 200
		resp["message"] = "User already exist."
		return resp, nil

	}
	return resp, nil
}

func LoginRepository(loginData models.Login) (models.Register, error) {
	var userData models.Register
	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()
	collection := session.DB(conf.Database).C(conf.UserCollecion)
	if err := collection.Find(bson.M{"email": loginData.Email, "password": loginData.Password}).One(&userData); err != nil {
		return userData, err
	}
	return userData, nil
}

func UpdateRepository(updateData models.User) (string, error) {
	var responseMessage string
	session, err := mgo.Dial(conf.MongoUrl)
	utils.Check(err)
	defer session.Close()

	filter := bson.M{"email": updateData.Email}
	change := bson.M{
		"$set": bson.M{
			"username": updateData.Name,
			"password": updateData.Password,
			"email":    updateData.Email,
			"address":  updateData.Address,
		},
	}

	collection := session.DB(conf.Database).C(conf.UserCollecion)

	err = collection.Update(filter, change)
	if err != nil {
		fmt.Println(err)
		return responseMessage, err
	}
	message := "User name : " + updateData.Name + " is updated successfully with email address ( " + updateData.Email + " )"
	responseMessage = message
	return responseMessage, nil
}
