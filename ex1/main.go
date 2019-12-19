package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Read csv file into array
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}
	r := csv.NewReader(csvFile)
	quiz, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Corrupt entries", err)
	}

	// Correct count
	count := 0

	// Parse the random flag -r and randomize the quiz if the flag is present
	randomPtr := flag.Bool("r", false, "a bool")
	flag.Parse()
	if *randomPtr {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(quiz), func(i, j int) { quiz[i], quiz[j] = quiz[j], quiz[i] })
		fmt.Println("random")
	}

	fmt.Print("Press enter to start the quiz")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t := time.NewTimer(30 * time.Second)
	go func() {
		<-t.C
		fmt.Println("\n", count, "out of", len(quiz), "correct")
		os.Exit(0)
	}()

	for _, entry := range quiz {
		fmt.Print(entry[0], ", ")

		for sc.Scan() {
			ans := strings.TrimSpace(sc.Text())
			if ans == entry[1] {
				count++
			}
			break
		}
	}
	fmt.Println(count, "out of", len(quiz), "correct")
}
