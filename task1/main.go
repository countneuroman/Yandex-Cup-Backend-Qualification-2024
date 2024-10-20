package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		stat     = make(map[int]float64)
		operator string
		id       int
		temp     float64
		sumTemp  float64
		parts    []string
	)
	reader := bufio.NewReader(os.Stdin)

loop:
	for loop := true; loop; {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			loop = false
			break loop
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Fprintln(os.Stderr, "Input is empty string")
			continue
		}
		parts = strings.Split(input, " ")

		operator = parts[0]

		switch operator {
		case "!":
			loop = false
			break loop
		case "?":
			if len(stat) == 0 {
				fmt.Printf("0.000000000\n")
			} else {
				average := sumTemp / float64(len(stat))
				fmt.Printf("%.9f\n", average)
			}
			continue
		case "-":
			id, err = strconv.Atoi(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if val, exists := stat[id]; exists {
				sumTemp -= val
				delete(stat, id)
			}
			continue
		case "+":
			id, err = strconv.Atoi(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			temp, err = strconv.ParseFloat(parts[2], 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if _, exists := stat[id]; exists {
				sumTemp -= stat[id]
			}
			stat[id] = temp
			sumTemp += temp
			continue
		case "~":
			id, err = strconv.Atoi(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			temp, err = strconv.ParseFloat(parts[2], 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if val, exists := stat[id]; exists {
				sumTemp -= val
			}
			stat[id] = temp
			sumTemp += temp
			continue
		default:
			fmt.Fprintln(os.Stderr, "operator not found")
			fmt.Fprintln(os.Stderr, operator)
			loop = false
			break loop
		}
	}
}
