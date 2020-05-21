package validaton

import (
	"github.com/inventory-app/backend/golang/models"
)

//TODO Validations..
func ProductValidation(product models.Product) (map[string]interface{}, error) {
	var errorResponse = make(map[string]interface{})
	return errorResponse, nil
}

func CategoryValidation(category models.Category) (map[string]interface{}, error) {
	var errorResponse = make(map[string]interface{})
	return errorResponse, nil
}
