package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func gabstest(data []byte) {
	fmt.Println("--------------------Using gabs------------------")
	jsonParsed, _ := gabs.ParseJSON(data)
	children, _ := jsonParsed.Children()
	for _, child := range children {
		fmt.Println(child)

		fmt.Println("--------Sugar methods----------")
		value, ok := child.Path("person.github.followers").Data().(float64)
		fmt.Printf("Followers: %f\n", value)
		fmt.Printf("Followers exist: %t\n", ok == true)

		companyName := child.Search("company", "name").Data().(string)
		fmt.Printf("company name: %s\n", companyName)

		value, ok = child.Path("does.not.exist").Data().(float64)
		fmt.Printf("does.not.exist: %f\n", value)
		fmt.Printf("does.not.exist: %t\n", !ok == true)
		// value == 0.0, ok == false

		exists := child.Exists("person", "github")
		fmt.Printf("Person github exists: %t\n", exists)

		exists = child.Exists("person", "linkedin")
		fmt.Printf("Person linkedin does not exist: %t\n", !exists)

		exists = child.ExistsP("person.facebook")
		fmt.Printf("Person facebook does not exist: %t\n", !exists)

		// Access embedded object
		names, _ := child.Search("person", "name").ChildrenMap()
		for key, name := range names {
			fmt.Printf("Key: %s\n  Value: %s\n", key, name)
		}

		// Access embedded array
		urls, _ := child.Search("person", "avatars").Children()
		for _, url := range urls {
			urlObj, _ := url.ChildrenMap()
			for key, val := range urlObj {
				fmt.Printf("Key: %s\n  Value: %s\n", key, val)
			}
		}

		// For quikly take all elements of an array, you can use this method:
		links := child.Path("person.avatars.url").String()
		fmt.Println("All links on avatar images:", links)
	}
}
