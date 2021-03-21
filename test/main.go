package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Args received: ", args)

	var in string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in += scanner.Text()
	}

	fmt.Println("Text: ", in)
	fmt.Println("Read everything from input. Thanks")
}
