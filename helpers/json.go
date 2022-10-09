package helpers

import (
	"encoding/json"
	"final-project/models"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
)

var Conn string

type DB struct {
}

func JsonDB(conn string) *DB {
	SetConn(conn)
	return &DB{}
}

func SetConn(conn string) {
	Conn = conn
}

func (db *DB) Update() {
	data := models.SensorData{
		Water: rand.Intn(100),
		Wind:  rand.Intn(100),
	}

	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(Conn, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) Get() models.SensorData {
	content, err := ioutil.ReadFile(Conn)
	if err != nil {
		log.Fatal(err)
	}

	data := models.SensorData{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
