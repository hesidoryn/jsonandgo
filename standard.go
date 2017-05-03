package main

import (
	"encoding/json"
	"fmt"
)

type Object struct {
	Pers Person  `json:"person"`
	Comp Company `json:"company"`
}

type Person struct {
	Name struct {
		First    string `json:"first"`
		Last     string `json:"last"`
		FullName string `json:"fullName"`
	} `json:"name"`
	Github struct {
		Handle    string `json:"handle"`
		Followers int    `json:"followers"`
	} `json:"github"`
	Avatars []struct {
		Url  string `json:"url"`
		Type string `json:"type"`
	} `json:"avatars"`
}

type Company struct {
	Name string `json:"name"`
}

func standard(data []byte) {
	fmt.Println("--------------------Standard methods------------------")
	fmt.Println("----------------Unmarshal---------------")
	objects := []Object{}
	json.Unmarshal(data, &objects)
	fmt.Println(objects)

	fmt.Println("----------------Marshal---------------")
	b, _ := json.Marshal(objects)
	fmt.Println(string(b))

	fmt.Println("----------------Unmarshal in interface---------------")
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)

	fmt.Println("----------------Creating array of interfaces---------------")
	m := f.([]interface{})

	fmt.Println("--Access first element--")
	first := m[0].(map[string]interface{})
	fmt.Println(first)

	fmt.Println("-Access first element company name-")
	fmt.Println((first["company"].(map[string]interface{})["name"]))
}
