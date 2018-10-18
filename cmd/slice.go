package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var sliceCmd = &cobra.Command{
	Use:   "slice file start end [outfile]",
	Short: "extract a part of a binary file",
	Long:  "extract some bytes out of a binary file",
	Args:  cobra.RangeArgs(3, 4),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file: %v", err)
			os.Exit(1)
		}
		start, serr := strconv.ParseInt(args[1], 0, 64)
		end, eerr := strconv.ParseInt(args[2], 0, 64)
		if serr != nil || eerr != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse offsets: %v, %v", serr, eerr)
			os.Exit(1)
		}
		buf := make([]byte, end-start)
		if l, err := f.ReadAt(buf, start); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file: %v", err)
			os.Exit(1)
		} else if l != len(buf) {
			fmt.Fprintf(os.Stderr, "Partial read: %v/%v", l, len(buf))
			os.Exit(1)
		}
		if len(args) > 3 {
			err := ioutil.WriteFile(args[3], buf, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create output file: %v", err)
				os.Exit(2)
			}
		} else {
			os.Stdout.Write(buf)
		}
	},
}
