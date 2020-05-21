package conf

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		Host          string `yaml:"host"`
		ServerAddress string `yaml:"serverAddress"`
	} `yaml:"server"`
	Database struct {
		Username      string `yaml:"user"`
		Password      string `yaml:"pass"`
		MongoUrl      string `yaml:"mongoUrl"`
		Database      string `yaml:"dbName"`
		UserCollecion string `yaml:"userCollecion"`
	} `yaml:"database"`
}

var Cfg Config

func init() {
	//f, err := os.Open("backend/golang/common/conf/config.yml")
	f, err := os.Open("backend/golang/common/conf/config.yml")
	if err != nil {
		log.Fatal("2--->", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Fatal("1", err)
	}
}

/*
Database and server details.
Always fetch this details from properties or yml file.
*/
const (
	MongoUrl            = "127.0.0.1"
	Database            = "inventoryDB"
	UserCollecion       = "user"
	PRODUCT_COLLECTION  = "product"
	CATEGORY_COLLECTION = "category"
	CART_COLLECTION     = "cart"
	ORDER_COLLECTION    = "order"
	PAYMENT_COLLECTION  = "payment"
	ServerAddress       = "127.0.0.1:8080"
)

/*
Unique key for JWT Token.
Do not store credentials inside projects.
*/
var JWT_KEY = []byte("inventory@123")
