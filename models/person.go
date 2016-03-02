package models
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
)

type Person struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Name string
}

func AddPerson(session *mgo.Session, newPerson Person) {
	c := session.DB("dev_skills").C("persons")
	fmt.Println(c)
	fmt.Println("Done");

	err := c.Insert(&newPerson)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPersons(session *mgo.Session) []Person {
	var results []Person
	err := session.DB("dev_skills").C("persons").Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}


func GetPerson(session *mgo.Session, name string) Person{
	result := Person{}
	err := session.DB("dev_skills").C("persons").Find(bson.M{"name": name}).Select(bson.M{"name": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func RemoveAllPersons(session *mgo.Session){
	_, err := session.DB("dev_skills").C("persons").RemoveAll(bson.M{})
	if err != nil{
		panic(err)
	}

}
