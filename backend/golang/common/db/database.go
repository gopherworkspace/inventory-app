package db

import (
	"github.com/inventory-app/backend/golang/common/conf"
	"gopkg.in/mgo.v2"
	"log"
)

func ConnectDB() (map[string]*mgo.Session, error) {
	mgoDBSession := make(map[string]*mgo.Session)

	//MongoDB session

	if session, err := mgo.Dial(conf.MongoUrl); err != nil {
		log.Println("Can't connect to mongodb database")
		return nil, err
	} else {
		mgoDBSession["coachingApp"] = session
		log.Println("Mongodb started successfully!")
		return mgoDBSession, nil
	}
}
