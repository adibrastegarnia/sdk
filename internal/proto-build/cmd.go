// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package proto_build

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "atomix-proto-build",
	}
	cmd.AddCommand(GetGoCommand())
	cmd.AddCommand(GetDocsCommand())
	return cmd
}
