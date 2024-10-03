package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var notes []string
	var noteLimit int

	fmt.Println("Enter the maximum number of notes: ")
	_, err := fmt.Scanf("%d", &noteLimit)
	if err != nil {
		fmt.Println(err)
	}

	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command and data: ")
		cmd, args, err := ParseCmd(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch cmd {
		case "create":
			// First, check if there are any empty slots in the slice
			noteCreated := false
			for i := range notes {
				if notes[i] == "" { // Reuse empty slot (from deleted notes)
					notes[i] = strings.Join(args, " ")
					noteCreated = true
					fmt.Print("[OK] The note was successfully created\n")
					break
				}
			}

			// If no empty slot was found, check if there's space to append
			if !noteCreated {
				if len(notes) < noteLimit {
					notes = append(notes, strings.Join(args, " "))
					fmt.Print("[OK] The note was successfully created\n")
				} else {
					fmt.Print("[Error] Notepad is full\n")
				}
			}

		case "list":
			hasNotes := false
			displayIndex := 1
			for _, n := range notes {
				if n != "" { // Skip empty entries
					fmt.Printf("[Info] %d: %s\n", displayIndex, n)
					displayIndex++
					hasNotes = true
				}
			}

			if !hasNotes {
				fmt.Print("[Info] Notepad is empty\n")
			}

		case "update":
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("[Error] Invalid position: %v\n", args[0])
			} else if index-1 < 0 || index > noteLimit {
				// Ensure we check against noteLimit, not len(notes)
				fmt.Printf("[Error] Position %v is out of the boundaries [1, %v]\n", args[0], noteLimit)
			} else {
				// Check if the position exists in the notes array, even if it's empty
				if index-1 >= len(notes) || notes[index-1] == "" {
					fmt.Printf("[Error] There is nothing to update\n")
				} else {
					// If there's a note to update
					notes[index-1] = strings.Join(args[1:], " ")
					fmt.Printf("[OK] The note at position %v was successfully updated\n", args[0])
				}
			}

		case "delete":
			index, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("[Error] Invalid position: %v\n", args[0])
			} else if index-1 < 0 || index > noteLimit {
				fmt.Printf("[Error] Position %v is out of the boundaries [1, %v]\n", args[0], noteLimit)
			} else {
				if index-1 >= len(notes) || notes[index-1] == "" {
					fmt.Printf("[Error] There is nothing to delete\n")
				} else {
					notes[index-1] = ""
					fmt.Printf("[OK] The note at position %v was successfully deleted\n", args[0])
				}
			}

		case "clear":
			notes = []string{}
			fmt.Print("[OK] All notes were successfully deleted\n")

		case "exit":
			fmt.Println("[Info] Bye!")
			return

		default:
			fmt.Print("[Error] Unknown command\n")
		}
	}
}

func ParseCmd(input *bufio.Scanner) (string, []string, error) {
	var cmd string
	var args []string
	if input.Scan() {
		input := input.Text()
		parts := strings.Split(input, " ")
		cmd = parts[0]

		switch cmd {
		case "create":
			if len(parts) < 2 {
				return "", nil, fmt.Errorf("[Error] Missing note argument\n")
			}

		case "update":
			if len(parts) < 2 {
				return "", nil, fmt.Errorf("[Error] Missing position argument\n")
			}
			if len(parts) < 3 {
				return "", nil, fmt.Errorf("[Error] Missing note argument\n")
			}

		case "delete":
			if len(parts) < 2 {
				return "", nil, fmt.Errorf("[Error] Missing position argument\n")
			}
		}

		if len(parts) > 1 {
			args = parts[1:]
		}

	}
	return cmd, args, nil
}
