package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/HannaAlda/forloop/database"
	"github.com/tikv/client-go/v2/rawkv"
)

func main() {

	content, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var objmap map[string]interface{}
	err2 := json.Unmarshal(content, &objmap)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}

	client, err := database.Connect()
	if err != nil {
		println(err)
	}

	insertData(client, objmap)
}

func insertData(client *rawkv.Client, objmap map[string]interface{}) {
	for key, data := range objmap {

		typeData := reflect.TypeOf(data).Kind()
		if typeData == reflect.Interface {
			b, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			var subdata map[string]interface{}
			err = json.Unmarshal(b, &subdata)
			if err != nil {
				println(err)
			}
			insertData(client, subdata)
		} else if !(typeData == reflect.String) {
			b, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			var subdata []map[string]interface{}
			err = json.Unmarshal(b, &subdata)
			if err != nil {
				println(err)
			}
			for _, d := range subdata {
				insertData(client, d)
			}
		} else {
			err := client.Put(context.TODO(), []byte(key), []byte(data.(string)))
			if err != nil {
				println(err)
			}
		}
	}
}
