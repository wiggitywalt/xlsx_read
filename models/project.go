package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
)


type Project struct {
	CostCenter string
	CCName string
	CCOwner string
	WBSName string
	WBS string
	ID        bson.ObjectId `bson:"_id,omitempty"`
}


func AddProject(session *mgo.Session, newProject Project) {
	c := session.DB("dev_skills").C("projects")
	fmt.Println(c)
	fmt.Println("Done");

	err := c.Insert(&newProject)
	if err != nil {
		log.Fatal(err)
	}
}

func GetProjects(session *mgo.Session) []Project{
	var results []Project
	err := session.DB("dev_skills").C("projects").Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

func RemoveAllProjects(session *mgo.Session){
	_, err := session.DB("dev_skills").C("projects").RemoveAll(bson.M{})
	if err != nil{
		panic(err)
	}
	fmt.Println("Ok, all projects have been removed.")
}