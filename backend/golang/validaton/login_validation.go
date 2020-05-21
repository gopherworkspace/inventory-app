package validaton

import (
	"errors"
	"github.com/inventory-app/backend/golang/models"
	"regexp"
	"strings"
)

/*
	Validation for all controllers request data.
*/
func RegistrationValidation(registrationData models.User) (map[string]interface{}, error) {
	var errorResponse = make(map[string]interface{})

	if strings.Trim(registrationData.Name, " ") == "" {
		message := "User name is missing."
		errorResponse["StatusCode"] = 204
		errorResponse["message"] = message
		return errorResponse, errors.New("User name is missing.")
	}

	if strings.Trim(registrationData.Password, " ") == "" {
		message := "Please enter password."
		errorResponse["StatusCode"] = 204
		errorResponse["message"] = message
		return errorResponse, errors.New(message)
	}

	if strings.Trim(registrationData.Address, " ") == "" {
		message := "Please enter address."
		errorResponse["StatusCode"] = 204
		errorResponse["message"] = message
		return errorResponse, errors.New(message)
	}

	if strings.Trim(registrationData.Email, " ") == "" {
		message := "Email address missing."
		errorResponse["StatusCode"] = 204
		errorResponse["message"] = message
		return errorResponse, errors.New(message)
	}
	return errorResponse, nil
}

func LoginValidation(loginRquestData models.Login) (map[string]interface{}, error) {
	var errorResponse = make(map[string]interface{})
	var errorMessage = ""

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(loginRquestData.Email) {
		errorMessage = " Please enter proper email address."
	}

	if strings.Trim(loginRquestData.Email, " ") == "" {
		errorMessage = " Email address is missing."
	}

	if strings.Trim(loginRquestData.Password, " ") == "" {
		errorMessage += " Please enter password."
	}

	if errorMessage != "" {
		errorResponse["message"] = errorMessage
		errorResponse["StatusCode"] = 204
		return errorResponse, errors.New(errorMessage)
	}
	return nil, nil
}
