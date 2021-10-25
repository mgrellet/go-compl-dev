package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}
	printMap(colors)

	myMap := make(map[int]string)
	myMap[10] = "a value"
	fmt.Println(myMap)
	delete(myMap, 10)
	fmt.Println(myMap)

}

func printMap(m map[string]string){
	for key, val := range m {
		fmt.Println(key+" - "+val)
	}
}
