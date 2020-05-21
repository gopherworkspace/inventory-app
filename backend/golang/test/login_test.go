package test

import (
	"fmt"
	"github.com/inventory-app/backend/golang/models"
	"github.com/inventory-app/backend/golang/services"
	"testing"
)

//Registration and login testing multiple parameters.
// So we took array of struct as request data.
//

func Test_Register(t *testing.T) {
	registerData := []models.Register{
		{"Kapil Yadav", "P@ss0rd", "kapil@gmail.com", "#101 51 street Bangalore"},
		{"", "P@ss0rd", "kapil@gmail.com", "#101 51 street Bangalore"},
		{"Kapil Yadav", "", "kapil@gmail.com", "#101 51 street Bangalore"},
		{"Kapil Yadav", "P@ss0rd", "", "#101 51 street Bangalore"},
		{"Kapil Yadav", "P@ss0rd", "kapil@gmail.com", ""},
	}

	for _, item := range registerData {
		response, err := services.Register(item)
		if err != nil {
			t.Errorf(fmt.Sprintf("Error message : %v  and request data : %v", response["message"], item))
		} else {
			t.Logf(fmt.Sprintf("Success message : %v  and request data : %v", response["message"], item))
		}
	}
}

func Test_Login(t *testing.T) {
	loginData := []models.Login{
		{"Kapil Yadav", "P@ss0rd"},
		{"", "P@ss0rd"},
		{"", ""},
		{"", "P@ss0rd"},
	}

	for _, item := range loginData {
		response, err := services.LoginService(item)
		if err != nil {
			t.Errorf(fmt.Sprintf("Error message : %v  and request data : %v", response["message"], item))
		} else {
			t.Logf(fmt.Sprintf("Success message : %v  and request data : %v", response["message"], item))
		}
	}
}
