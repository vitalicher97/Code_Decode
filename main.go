package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func encode(cypherMap map[string]string) {

	fmt.Println("Print text to encode:")

	reader := bufio.NewReader(os.Stdin)
	text, _:= reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	//length := len(text)

	resultText := ""

	for _, letter := range text {
		element, ok := cypherMap[string(letter)]
		if ok {
			resultText = resultText + element
			continue
		}
		resultText = resultText + string(letter)
	}

	fmt.Println(resultText)

}

func decode(decypherMap map[string]string) {

	fmt.Println("Print text to deencode:")

	reader := bufio.NewReader(os.Stdin)
	text, _:= reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	//length := len(text)

	resultText := ""

	for _, letter := range text {
		element, ok := decypherMap[string(letter)]
		if ok {
			resultText = resultText + element
			continue
		}
		resultText = resultText + string(letter)
	}

	fmt.Println(resultText)

}

func main() {

	cypherMap := make(map[string]string)

	cypherMap["A"] = "1"
	cypherMap["a"] = "1"
	cypherMap["E"] = "2"
	cypherMap["e"] = "2"
	cypherMap["I"] = "3"
	cypherMap["i"] = "3"
	cypherMap["O"] = "4"
	cypherMap["o"] = "4"
	cypherMap["U"] = "5"
	cypherMap["u"] = "5"

	decypherMap := map[string]string {
		"1": "a",
		"2": "e",
		"3": "i",
		"4": "o",
		"5": "u",
	}




	fmt.Println("Input 'encode' or 'decode':")

	reader := bufio.NewReader(os.Stdin)
	text, _:= reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if text == "encode" {
		encode(cypherMap)
	} else if text == "decode" {
		decode(decypherMap)
	} else {
		fmt.Println("Invalid input")
	}

}
