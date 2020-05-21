package context

import (
	"github.com/inventory-app/backend/golang/common/db"
	"gopkg.in/mgo.v2"
	"log"
)

type DataBase struct {
	DataBaseSession map[string]*mgo.Session
}

var AppContext *DataBase

func init() {
	AppContext = new(DataBase)
}

func ApplicationContext() {
	var err error
	var appDb = DataBase{}

	if AppContext.DataBaseSession, err = db.ConnectDB(); err != nil {
		log.Println("Can't connect to mongodb database")
	} else {
		appDb.DataBaseSession = AppContext.DataBaseSession
	}
}
