/*
Write a program which prompts the user to first
enter a name, and then enter an address.
Your program should create a map and add the name
and address to the map using the keys “name”
and “address”, respectively.
Your program should use Marshal() to create a
JSON object from the map,
 and then your program should print the JSON object.

1 Input Name
2 Input Address
3 Store both into Map
4 Convert Map to Json
5 Print Json

*/

package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	var name string
	var address string

	fmt.Printf("Enter your Name\n")
	fmt.Scan(&name)

	fmt.Printf("Enter your Address\n")
	fmt.Scan(&address)

	var m = make(map[string]string)
	m["name"] = name
	m["address"] = address

	var m_json, _ = json.Marshal(m)

	fmt.Printf("Map content (json): %v\n", string(m_json) )

}
