package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kinpoko/grvemu/rv32i"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grvemu [binary file] [instruction]",
	Short: "A toy RISC-V emulator for cli written in Go",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inst, err := cmd.Flags().GetString("inst")
		if err != nil {
			return err
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		switch inst {
		case "r32i":
			file, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer file.Close()

			binary, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}

			if err := rv32i.Run(binary, debug); err != nil {
				return err
			}

		// TODO support r64i
		// case "r64i":
		// fmt.Println(inst)

		default:
			fmt.Printf("%s is not supported\n", inst)
			fmt.Printf("this emulator supports r32i\n")
		}
		return nil

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("inst", "i", "r32i", "instruction")
	rootCmd.Flags().BoolP("debug", "d", false, "debug mode")
}
