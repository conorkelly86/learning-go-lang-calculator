package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	
	reader :=  bufio.NewReader(os.Stdin)
	// fmt.Print("Enter Value 1: ")
	// value1, _ := reader.ReadString('\n')
	// fmt.Println("You entered: ", value1)


	fmt.Print("Enter Value 1: ")
	value1, _ := reader.ReadString('\n')
	value1Float, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", value1Float)
	}
	fmt.Print("Enter A Number: ")
	value2, _ := reader.ReadString('\n')
	value2Float, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", value2Float)
	}


	result := value1Float + value2Float
	fmt.Println(result)

}
