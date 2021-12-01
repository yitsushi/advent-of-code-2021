package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/go-aoc/client"
)

func submitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit the solution",
		Run: func(cmd *cobra.Command, args []string) {
			solution, err := getSolution(cmd)
			if err != nil {
				logrus.Fatalf("Unable to solve the puzzle: %s\n\n", err)

				return
			}

			aocClient := client.NewClient(os.Getenv("AOC_SESSION"))

			dayNumber, _ := cmd.Flags().GetInt("day")
			partNumber, _ := cmd.Flags().GetInt("part")

			valid, err := aocClient.SubmitSolution(currentYear, dayNumber, partNumber, solution)
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())

				return
			}

			if valid {
				fmt.Fprintln(cmd.OutOrStdout(), "Done \\o/")
			} else {
				fmt.Fprintln(cmd.ErrOrStderr(), "Something is wrong :(")
			}
		},
	}

	cmd.Flags().Int("day", 1, "Day")
	cmd.Flags().Int("part", 1, "Part")
	cmd.Flags().String("input", "", "Input as value")
	cmd.Flags().String("input-file", "", "Input as file (path)")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
