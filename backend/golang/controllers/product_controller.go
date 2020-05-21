package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/inventory-app/backend/golang/models"
	"github.com/inventory-app/backend/golang/services"
	"log"
	"net/http"
)

func (c *Controller) CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	var Product models.Product

	if err := json.NewDecoder(r.Body).Decode(&Product); err != nil {
		log.Fatal(err)
		return
	}
	respnce, err := services.CreateNewProduct(Product)

	if err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}
	resp, err := json.Marshal(respnce)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return
}

func (c *Controller) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		log.Fatal(err)
		return
	}
	respnce, err := services.CreateNewCategory(category)

	if err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}
	resp, err := json.Marshal(respnce)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return
}

func (c *Controller) GetCategory(w http.ResponseWriter, r *http.Request) {
	if respnce, err := services.GetCategory(); err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	} else {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

}

func (c *Controller) OrderProduct(w http.ResponseWriter, r *http.Request) {

	if respnce, err := services.OrderList(); err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	} else {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

}

func (c *Controller) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	var cartItems []models.ProductItems
	if err := json.NewDecoder(r.Body).Decode(&cartItems); err != nil {
		log.Fatal(err)
		return
	}

	tokenString, err := r.Cookie("access_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(tokenString.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("access_token"), nil
	})

	if respnce, err := services.AddProductToCart(claims, cartItems); err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	} else {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

}

func (c *Controller) OrderList(w http.ResponseWriter, r *http.Request) {
	if respnce, err := services.GetCategory(); err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	} else {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

}

func (c *Controller) OrderPayment(w http.ResponseWriter, r *http.Request) {

	tokenString, err := r.Cookie("access_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(tokenString.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("access_token"), nil
	})

	email := fmt.Sprintf("%v", claims["email"])

	if respnce, err := services.OrderPaymentService(email); err != nil {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	} else {
		resp, err := json.Marshal(respnce)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		return
	}

}
