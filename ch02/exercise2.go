package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func exercise2() {
	fmt.Println("exercise1")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter racer name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)

	fmt.Print("Enter racer rank: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	rank, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	medal := "participant"
	// YOUR CODE HERE: Write some code that looks at the
	// "rank" variable and sets the "medal" variable to an
	// appropriate string. If "rank" is 1, "medal" should
	// be set to "gold". If "rank" is 2, "medal" should be
	// "silver". A rank of 3 should get a "bronze" medal,
	// and any other rank should get a "participant" medal.
	if rank == 3 {
		medal = "bronze"
	} else if rank == 2 {
		medal = "silver"
	} else if rank == 1 {
		medal = "gold"
	}

	fmt.Println(name, "gets a", medal, "medal!")
}
