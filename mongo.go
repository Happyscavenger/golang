package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
#mongo数据库链接
const URL = "xxx"

func main() {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(nil)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("caiji")
	collection := db.C("fangtianxia")

	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println(countNum)

	var results []map[string]interface{}
	collection.Find(bson.M{"tag": "1"}).All(&results)
	for _, result := range results {
		fmt.Println(result["url"])
	}
}
