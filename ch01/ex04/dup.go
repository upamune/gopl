package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DupCount struct {
	Count     int
	FileNames map[string]struct{}
}

func NewDupCount(fileName string) *DupCount {
	return &DupCount{
		Count:     0,
		FileNames: make(map[string]struct{}),
	}
}

func (c *DupCount) Increment(fileName string) {
	c.Count += 1
	c.FileNames[fileName] = struct{}{}
}

func main() {
	counts := make(map[string]*DupCount)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, count := range counts {
		if count.Count > 1 {
			fileNames := []string{}
			for name := range count.FileNames {
				fileNames = append(fileNames, name)
			}

			fmt.Printf("%d\t%s\t%s\n", count.Count, line, strings.Join(fileNames, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]*DupCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if _, ok := counts[text]; !ok {
			dupCount := NewDupCount(f.Name())
			counts[text] = dupCount
		}
		counts[text].Increment(f.Name())
	}
}
