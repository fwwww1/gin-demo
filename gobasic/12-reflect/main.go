package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reflectTest(arg interface{}) {
	fmt.Println("arg type is", reflect.TypeOf(arg))
	fmt.Println("arg type is", reflect.ValueOf(arg))
}

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	var num float64 = 1.2343
	reflectTest(num)

	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhang"}}
	//编码过程  结构体--> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json Marshal err", err)
		return
	}
	fmt.Printf("jsonSre=%s\n", jsonStr)

	//解码过程  json-->结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unMarshal err", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
}
