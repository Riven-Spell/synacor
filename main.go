package main

import (
	"fmt"
	"os"
	"synacor/interpreter"
)

func main() {
	if len(os.Args) > 1 {
		if fi, err := os.Stat(os.Args[1]); err == nil {
			buffer := make([]byte, fi.Size())
			fmt.Println("Loaded", len(buffer)/2, "words (16 bits) into program memory.")

			f, err := os.Open(os.Args[1])
			if err != nil {
				fmt.Println("Could not open file: " + err.Error())
			}

			_, err = f.Read(buffer)
			if err != nil {
				fmt.Println("Could not read file: " + err.Error())
			}

			system := interpreter.NewSystem()
			system.LoadProgram(buffer)
			if err := system.StartSystem(); err != nil {
				fmt.Println("Encountered an error during execution: " + err.Error())
			}
		} else {
			fmt.Println("Encountered an error loading file: " + err.Error())
		}
	} else {
		fmt.Println("Usage: synacor [binary]")
	}
}
