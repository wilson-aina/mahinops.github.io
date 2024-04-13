package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}

	currentDate := time.Now()

	// Extract year, month, and day
	year := currentDate.Year()
	month := int(currentDate.Month())
	day := currentDate.Day()

	// Get the argument provided in the terminal
	arg := os.Args[1]

	// Create the filename
	filename := fmt.Sprintf("_posts/%d-%02d-%02d-leetcode-%s.md", year, month, day, arg)

	// Check if file exists
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("File", filename, "already exists")
		return
	}

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully:", filename)

	writeToFile(arg, file, year, month, day)

}

func writeToFile(s string, file *os.File, year, month, day int) {
	// fmt.Fprintln(file, s)
	baseURL := fmt.Sprintf("https://leetcode.com/problems/%s/description/", s)

	fmt.Println(baseURL)

	// Write multiple lines to the file
	lines := []string{
		"---",
		"title: \"Leetcode - Dummy\"",
		fmt.Sprintf("date: %d-%02d-%02d", year, month, day),
		"categories: [Leetcode]",
		"tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm]",
		"---",
		"",
		"",
		fmt.Sprintf("# [Leetcode - Dummy](%s){:target=\"_blank\"}", baseURL),
		fmt.Sprintf("[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-%s.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-%s/)", s, s),
		"",
		"",
		"-",
		"",
		"",
		"Solution",
		"",
		"```",
		"```",
	}

	// [![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-arg.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-arg/)

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Multiple lines written to file:", file.Name())

}
