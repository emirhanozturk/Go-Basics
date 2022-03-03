package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[English Words] -> [Turkish Words]")
		return
	}

	query := args[0]

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
	}
	//to delete key in map
	delete(dict, "awesome")
	delete(dict, "awesome")
	delete(dict, "notexisting") //if we want to delete nonexisting word,nothing happens
	//we don't need to chech whether there is or not.We can use delete function directly

	for k := range dict { // #9
		delete(dict, k)
	} //This loop look like inefficient but behind the scenes go convert it to only one single operation which called mapClear()

	turkish := make(map[string]string, len(dict))
	//create new map with new memory adress
	for k, v := range dict {
		turkish[v] = k
	}

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found.", query)

	fmt.Printf("# of keys: %d\n", len(dict))
}
