// Copyright 2023 OWASP Core Rule Set Project
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Traceableai/crs-toolchain/v2/internal/updater"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of crs-toolchain",
	Long:  `All software has versions. This is crs-toolchain's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crs-toolchain", rootCmd.Version)
		// do not run when in CI (e.g. GitHub Actions)
		if os.Getenv("CI") != "true" {
			latest, err := updater.LatestVersion()
			if err != nil {
				logger.Error().Err(err).Msg("Failed to check for updates")
			} else if latest != "" {
				fmt.Println("Latest version is:", latest)
				fmt.Println("Run 'crs-toolchain self-update' to update")
			}
		}
	},
}
