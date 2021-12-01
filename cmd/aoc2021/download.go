package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yitsushi/go-aoc/client"
)

func downloadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download a puzzle input",
		Run: func(cmd *cobra.Command, args []string) {
			dayNumber, _ := cmd.Flags().GetInt("day")

			targetDir := fmt.Sprintf("input/day%02d", dayNumber)
			targetFile := fmt.Sprintf("%s/input", targetDir)

			ensurePath(targetDir)

			aocClient := client.NewClient(os.Getenv("AOC_SESSION"))

			err := aocClient.DownloadAndSaveInput(currentYear, dayNumber, targetFile)
			if err != nil {
				logrus.Fatal(err.Error())

				return
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Done... %s\n", targetFile)
		},
	}

	cmd.Flags().Int("day", 1, "Day")

	_ = cmd.MarkFlagRequired("day")

	return cmd
}
