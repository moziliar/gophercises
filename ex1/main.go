package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
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

	for _, entry := range quiz {
		fmt.Print(entry[0], ", ")

		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			ans := sc.Text()
			if ans == entry[1] {
				count++
			}
			break
		}
	}
	fmt.Println(count, "out of", len(quiz), "correct")
}
