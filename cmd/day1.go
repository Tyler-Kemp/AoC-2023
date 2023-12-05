package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var day1Cmd = &cobra.Command{
	Use:   "1",
	Short: "Run the solution for day 1",
	Long: `Run the solution for day 1.
		For example:
		aoc2023 1 -f input.txt
		will run the solution for day 1 using the input.txt file as input.`,
	Run: runDay,
}

func init() {
	rootCmd.AddCommand(day1Cmd)
	day1Cmd.Flags().StringP("file", "f", "", "Specify an input file")
}

func runDay(cmd *cobra.Command, args []string) {
	fileName, _ := cmd.Flags().GetString("file")
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	var lines []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		lines = append(lines, line)
	}

	total := 0

	for _, line := range lines {
		var firstNumberIndex int
		var lastNumberIndex int

		for i, char := range line {
			if char >= '0' && char <= '9' {
				firstNumberIndex = i
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if char >= '0' && char <= '9' {
				lastNumberIndex = i
				break
			}
		}

		newNumString := line[firstNumberIndex:firstNumberIndex+1] + line[lastNumberIndex:lastNumberIndex+1]

		newNum, err := strconv.Atoi(newNumString)
		if err != nil {
			log.Fatal(err)
		}

		total += newNum
		fmt.Printf("Total: %d\n", total)
	}
}
