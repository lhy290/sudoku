package main

import (
	"sudoku/pkg/eliminater"
	"sudoku/pkg/logger"

	"github.com/spf13/cobra"
)

func main() {
	var (
		sudokufile string
		logLevel   int
		LogFile    string
		result     string
	)

	parseCmd := &cobra.Command{
		Use:   "parse [-f sudokufile -O resultFile]",
		Short: "parse sudoku",
		Run: func(cmd *cobra.Command, args []string) {
			logger.InitLog(logLevel, LogFile)
			eliminater.
				New(sudokufile, result).
				Run()
		},
	}

	parseCmd.Flags().StringVarP(&sudokufile, "file", "f", "", "sudoku file")
	parseCmd.Flags().StringVarP(&result, "O", "", "result", "result file")
	parseCmd.Flags().IntVarP(&logLevel, "log-level", "", 5, "5 - debug, 4 - info, 2 - error")
	parseCmd.Flags().StringVarP(&LogFile, "log-file", "", "", "log file")
	parseCmd.Execute()
}
