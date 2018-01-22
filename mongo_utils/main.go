package main 

import (
	"fmt"
	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DbName string = "golang"

func findById (coll string, id string) {
	// collection:= 
}

var _session *mongo.Session 

func getDBSession() (mongo.Session error){
	if !_session {
		session, err:= mongo.Dial("localhost")
		if err != nil {
			return nil, err
		}
		session.SetMode(mgo.Monotonic, true)
		return session, nil
	}
}


func main (){
	session, err: = getDBSession()
	if err != nil {
		panic(err)
		return
	}
	defer session.Close()
}