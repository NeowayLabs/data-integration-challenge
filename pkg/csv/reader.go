package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// ReadAll opens a channel and returns all records as an array of strings
func ReadAll(path string, fieldSeparator rune, chanSize int) <-chan []string {
	cs := make(chan []string, chanSize)

	go func() {
		defer close(cs)

		fr, err := os.Open(path)
		if err != nil {
			log.Fatalf("Failed to open file: %s. Error: %v", path, err)
			return
		}
		defer fr.Close()

		r := csv.NewReader(bufio.NewReader(fr))
		r.Comma = fieldSeparator

		for {
			record, err := r.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				close(cs)
				log.Fatal(err)
			}

			cs <- record
		}
	}()

	return cs
}
