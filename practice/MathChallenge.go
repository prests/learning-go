package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number 1: ")
	input1, _ := reader.ReadString('\n')

	float1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err1 != nil {
		panic("input1 was not a number")
	}

	fmt.Print("Enter number 2: ")
	input2, _ := reader.ReadString('\n')

	float2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err2 != nil {
		panic("input2 was not a number")
	}

	sum := float1 + float2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)
}
