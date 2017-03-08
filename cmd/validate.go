// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/st3fan/xliff"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate an XLIFF file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runValidate,
}

func init() {
	RootCmd.AddCommand(validateCmd)
	validateCmd.Flags().BoolP("allow-incomplete", "", false, "Incomplete files do not fail validation")
	validateCmd.Flags().BoolP("quiet", "", false, "Do not print validation errors")
}

func runValidate(cmd *cobra.Command, args []string) {
	status := 0
	for _, path := range args {
		doc, err := xliff.FromFile(path)
		if err != nil {
			status = 1
			fmt.Printf("could not parse %s: %s\n", path, err.Error())
			continue
		}

		errors := doc.Validate()
		if len(errors) != 0 {
			status = 1
			if !cmd.Flag("quiet") {
				for _, err := range errors {
					fmt.Printf("%s: %s\n", path, err.Message)
				}
			}
		}
	}
	os.Exit(status)
}
