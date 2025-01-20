package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file := os.Args[1]

	if file == "-" {
		fmt.Println("Stdin block hit")
		stdin := bufio.NewReader(os.Stdin)
		data, err := io.ReadAll(stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Printf("%s", data)

	}else if file == "-n"{
		fmt.Println("Numbering block hit!")
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
	} else if len(os.Args) > 2  {
		fmt.Println("Concatenation block hit!")
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
		// fmt.Fprintln(os.Stdout, output)
	}else {

		fmt.Println("One arg block hit")
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
}

