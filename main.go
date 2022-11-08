package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var sides []rune = []rune{'F', 'B', 'L', 'R', 'U', 'D'}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomChoice() string {
	var b strings.Builder
	side := sides[rand.Intn(len(sides))]
	b.WriteRune(side)
	if RandomBool() {
		b.WriteRune('\'')
	}
	return b.String()
}

func main() {
	rand.Seed(time.Now().UnixMicro())

	times := 20
	if len(os.Args) >= 2 {
		arg, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Argument must be a number")
		}
		times = arg
	}
	for i := 0; i < times; i++ {
		fmt.Print(RandomChoice(), " ")
	}
	fmt.Println()
}
