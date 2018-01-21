package lookup_pipeline

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"




	"gopkg.in/mgo.v2/bson"

	"os"
	"encoding/json"
)

type Person struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Phone string
	Children bson.ObjectId `bson:"children,omitempty"`
}

type Child struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Type string
	Value string
}

//type Per map[string]interface{}



func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	servs := session.Mode()
	fmt.Println(servs)
	c := session.DB("golangdb").C("people")
	childColl:= session.DB("golangdb").C("peopleChild")
	child:= &Child{Type:"some1", Value:"hello"}


	//fmt.Println(res)
	//var child Child
	err = childColl.Find(bson.M{"type":"some1"}).One(&child)
	if err != nil {
		panic(err)
	}
	fmt.Println(child)

	//err = c.Insert(&Person{Name:"Ale", Phone:"+55 53 8116 9639", Children:child.ID})
	//if err != nil {
	//	log.Fatal(err)
	//}



	//result := Person{}
	//fmt.Println(c.Count())
	//explanation := bson.M{}
	//err = c.Find(bson.M{"name": "Ale"}).Explain(explanation)
	//var m []Person
	pipeline:= []bson.M{
		bson.M{"$lookup":bson.M{"from":"peopleChild", "localField":"children", "foreignField":"_id", "as":"children"},
	}}
	result:=[]bson.M{}

	err = c.Pipe(pipeline).All(&result)
	b, err:= json.Marshal(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		os.Stdout.Write(b)
		//n, err:= os.Hostname()

	}

	//fmt.Println("Phone:", result.Phone)
}
