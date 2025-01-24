package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
func dash() {
	stdin := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Printf("%s", data)
}
func dash_b() {
	reader := bufio.NewReader(os.Stdin)
	lineCount := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			lineCount++
			fmt.Printf("%d %s\n", lineCount,line )
		} else {
			fmt.Printf("%s\n", line )
		}
	}
}

func dash_n() {
	reader := bufio.NewReader(os.Stdin)
	lineCount := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineCount++
		fmt.Printf("%d %s\n", lineCount,line )
	}
}

func concat() {
	var output io.Writer
	for _, file := range os.Args[1:] {
		input, err := os.Open(file)
		if err != nil {
			fmt.Errorf("Error openning file %s: %w", file, err)
		}
		
		output = os.Stdout

		_, err = io.Copy(output, input)
		if err != nil {
			input.Close()
			fmt.Errorf("Error copying file %s: %w", file, err)
		}
	}
}

func normal() {
	file := os.Args[1]

	open_file, err := os.Open(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer open_file.Close()

	data, err := io.ReadAll(open_file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Printf("%s", data)
}

func main() {
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Printf("%s\n", line)
		}
	}
	argType := os.Args[1]
	
	if strings.HasPrefix(argType, "-") {
		switch argType {
		case "-":
			dash()
			break
		case "-b":
			dash_b()
			break;
		case "-n":
			dash_n()
			break
		default:
			fmt.Printf("Invalid argument passed: %s\n try -b, -n\n", argType)
		
		}		
		
	}
	if !strings.HasPrefix(argType, "-") {
		if len(os.Args) > 1 && len(os.Args) == 2 {
			normal()
			return
		}else if len(os.Args) > 2 {
			concat()
			return
		} 
	}
}
	
