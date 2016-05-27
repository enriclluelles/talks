package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

//START EXPLICIT OMIT
type Body struct {
	Response Response
}

type Response struct {
	Data Data
}

type Data struct {
	Users []User
}

type User struct {
	LoginMethods []LoginMethod `json:"loginMethods"`
}

type LoginMethod struct {
	Data Data2
}

type Data2 struct {
	DeviceType string    `json:"deviceType"`
	Date       time.Time `json:"date"`
}

//END EXPLICIT OMIT

func main() {
	var r Body
	f, err := os.Open("sample.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(f).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
	js, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(js))
}
