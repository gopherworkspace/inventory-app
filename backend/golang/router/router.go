package router

import (
	"github.com/gorilla/mux"
	"github.com/inventory-app/backend/golang/common/auth"
	"github.com/inventory-app/backend/golang/common/utils"
	"github.com/inventory-app/backend/golang/controllers"
	"log"
	"net/http"
)

var controller = &controllers.Controller{}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Role        []string
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{"Health Check", "GET", "/healthcheck", controller.HealthCheck, []string{"ADMIN"}},

	Route{"Register", "POST", "/register", controller.Register, []string{"ADMIN"}},

	Route{"Login", "POST", "/login", controller.Login, []string{"ADMIN"}},

	Route{"Logout", "POST", "/logout", controller.Logout, []string{"ADMIN"}},

	Route{"Update", "POST", "/update", controller.Update, []string{"ADMIN"}},

	Route{"Create New Product", "POST", "/newProduct", controller.CreateNewProduct, []string{"ADMIN"}},

	Route{"Create New Category", "POST", "/newCategory", controller.CreateNewCategory, []string{"ADMIN"}},

	Route{"Category List", "GET", "/category/list", controller.GetCategory, []string{"ADMIN"}},

	Route{"Order Product", "POST", "/order/product", controller.OrderProduct, []string{"ADMIN"}},

	Route{"Add Product To Cart", "POST", "/cart/add/product", controller.AddProductToCart, []string{"ADMIN"}},

	Route{"List of Product", "POST", "/order/list", controller.OrderList, []string{"ADMIN"}},

	Route{"Order Payment", "POST", "/order/payment", controller.OrderPayment, []string{"ADMIN"}},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(CommonMiddleWare)

	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		if !utils.Contains(route.Role, "ADMIN") {
			//http.Redirect(w, r, EMPTY_STRING, http.StatusUnauthorized)
		}

		s := router.PathPrefix("/auth").Subrouter()
		s.Use(auth.JwtVerify)
	}
	return router
}
func CommonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(res, req)
	})
}
