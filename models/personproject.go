package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
)


type Personproject struct {
  Userid string
	WBS string
	Startdate string
	ID bson.ObjectId `bson:"_id,omitempty"`
}


func AddPersonProject(session *mgo.Session, newPersonProject Personproject) {
	c := session.DB("dev_skills").C("personprojects")
	fmt.Println(c)
	fmt.Println("Done");

	err := c.Insert(&newPersonProject)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPersonProjects(session *mgo.Session) []Project{
	var results []Project
	err := session.DB("dev_skills").C("personprojects").Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

func RemoveAllPersonProjects(session *mgo.Session){
	_, err := session.DB("dev_skills").C("personprojects").RemoveAll(bson.M{})
	if err != nil{
		panic(err)
	}
}