package main 

import (
	//"fmt"
	mongo "gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
	//"github.com/fatih/structs"
	"fmt"

	"reflect"
)

const DbName string = "golangdb"

func findById (coll string, id string) (bson.M, error){
	 collection:= _session.DB(DbName).C(coll)
	 document:=bson.M{}
	 hex:= bson.ObjectIdHex(id)
	 err := collection.FindId(hex).One(&document)
	 if err != nil {
			return nil, err
	 }
	 return document, nil
}



func findOne (coll string, query map[string]interface{}) (bson.M, error){
	collection:= _session.DB(DbName).C(coll)
	document:= bson.M{}
	err := collection.Find(query).One(&document)
	if err != nil {
		return nil, err
	}
	return document, nil
}

var _session *mongo.Session 

func getDBSession() (*mongo.Session, error){
	if _session == nil {
		session, err:= mongo.Dial("localhost")
		if err != nil {
			return nil, err
		}
		session.SetMode(mongo.Monotonic, true)
		return session, nil
	} else {
		return _session, nil
	}

}


func main (){
	session, err:= getDBSession()
	if err != nil {
		panic(err)
		return
	} else {
		_session = session
	}
	defer session.Close()
	docByID, err:= findById("people", "5a651bc9515e23335e0bed75")
	if err != nil {
		panic(err)
	}

	docByField, err:= findOne("people", map[string]interface{}{"name":"Ale"})
	if err != nil {
		panic(err)
	}
	fmt.Println("found by id:", docByID, reflect.TypeOf(docByID["phone"]))
	fmt.Println("found by field:", docByField, reflect.TypeOf(docByID["phone"]))

}