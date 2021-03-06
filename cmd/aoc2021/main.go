package main

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	currentYear = 2021
)

func rootCommand() cobra.Command {
	cmd := cobra.Command{
		Use:   "aoc2021",
		Short: "Advent of Code 2021",
		Long:  "All my solutions \\o/",
	}

	allLogLevels := []string{}

	for _, level := range logrus.AllLevels {
		allLogLevels = append(allLogLevels, level.String())
	}

	cmd.PersistentFlags().String(
		"log-level",
		logrus.WarnLevel.String(),
		fmt.Sprintf("Log level (%s)", strings.Join(allLogLevels, ", ")),
	)

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		logrus.SetOutput(cmd.OutOrStdout())

		level, _ := cmd.Flags().GetString("log-level")

		lvl, err := logrus.ParseLevel(level)
		if err != nil {
			return err
		}

		logrus.SetLevel(lvl)

		return nil
	}

	cmd.AddCommand(solveCommand())
	cmd.AddCommand(downloadCommand())
	cmd.AddCommand(submitCommand())
	cmd.AddCommand(scaffoldCommand())

	return cmd
}

func main() {
	root := rootCommand()
	_ = root.Execute()
}
