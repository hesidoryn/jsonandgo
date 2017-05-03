package main

import (
	"fmt"
	"strconv"

	"github.com/buger/jsonparser"
)

func jsonparsertest(data []byte) {
	fmt.Println("--------------------Uing jsonparser------------------")
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		result1, _, _, _ := jsonparser.Get(value, "company", "name")
		result2, _ := jsonparser.GetString(value, "company", "name")
		fmt.Println(string(result1))
		fmt.Println(string(result2))

		jsonparser.ObjectEach(value, func(key []byte, val []byte, dataType jsonparser.ValueType, offset int) error {
			fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n Offset: %d\n", string(key), string(val), dataType, offset)
			return nil
		}, "company", "name")

		// If you want to access only specifical keys in JSON-object, you can create object with paths to them
		paths := [][]string{
			[]string{"person", "name", "fullName"},
			[]string{"person", "github", "followers"},
			[]string{"person", "avatars", "[0]", "url"},
			[]string{"company", "name"},
		}
		jsonparser.EachKey(value, func(idx int, val []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0: // []string{"person", "name", "fullName"}
				fmt.Println(string(val))
			case 1: // []string{"person", "github", "followers"}
				followers, _ := strconv.Atoi(string(val))
				fmt.Println(followers)
			case 2: // []string{"person", "avatars", "[0]", "url"}
				fmt.Println(string(val))
			case 3: // []string{"company", "url"}
				fmt.Println(string(val))
			}
		}, paths...)
	})
}
