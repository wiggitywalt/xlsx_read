package mongoService
import (
	"gopkg.in/mgo.v2"
	"time"
)
//Rajesh mlab: "ds011158.mlab.com:11158"
//my mlab: "ds043982.mlab.com:43982"

func CreateMongoSession() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
    //rajesh
		// Addrs:    []string{"ds011158.mlab.com:11158"},

    //mine
    Addrs:    []string{"ds043982.mlab.com:43982"}, 

		Timeout:  60 * time.Second,
		Database: "dev_skills",
		Username: "disney",
		Password: "disney",
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	return session

}
