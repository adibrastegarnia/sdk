// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package build

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "atomix-plugin-build",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
