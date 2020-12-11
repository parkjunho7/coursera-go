package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var name string
	var addr string
	var hashMap map[string]string = make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter your Name : ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Printf("Please enter your Address : ")
	scanner.Scan()
	addr = scanner.Text()
	hashMap[name] = addr
	b, err := json.MarshalIndent(hashMap, "", "    ")
	if err == nil {
		fmt.Printf("%s\n", b)
	}

}
