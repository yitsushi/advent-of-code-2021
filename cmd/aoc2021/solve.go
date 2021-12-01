package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/go-aoc/puzzle"
)

func getSolution(cmd *cobra.Command) (string, error) {
	dayNumber, _ := cmd.Flags().GetInt("day")
	partNumber, _ := cmd.Flags().GetInt("part")
	inputValue, _ := cmd.Flags().GetString("input")
	inputFile, _ := cmd.Flags().GetString("input-file")

	month := withDays(puzzle.NewMonth())

	day, err := month.Get(dayNumber)
	if err != nil {
		return "", err
	}

	solver := puzzle.NewSolver(day, partNumber)

	return solver.Solve(inputValue, inputFile)
}

func solveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "solve",
		Short: "Solve a puzzle",
		Run: func(cmd *cobra.Command, args []string) {
			solution, err := getSolution(cmd)
			if err != nil {
				logrus.Fatalf("Unable to solve the puzzle: %s\n\n", err)

				return
			}

			fmt.Fprintln(cmd.OutOrStderr(), solution)
		},
	}

	cmd.Flags().Int("day", 1, "Day")
	cmd.Flags().Int("part", 1, "Part (0 => both)")
	cmd.Flags().String("input", "", "Input as value")
	cmd.Flags().String("input-file", "", "Input as file (path)")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
