package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Args received: ", args)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("Read everything from input. Thanks")
}
