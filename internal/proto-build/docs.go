// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package proto_build

import (
	"fmt"
	"github.com/atomix/runtime-api/internal/exec"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetDocsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "docs",
		Aliases: []string{"markdown", "md"},
		Args:    cobra.ExactArgs(1),
		RunE:    runDocsCommand,
	}
	cmd.Flags().StringP("input", "i", ".", "the path to the root of the Protobuf sources")
	cmd.Flags().StringP("pattern", "f", "**/*.proto", "a pattern by which to filter Protobuf sources")
	cmd.Flags().StringP("output", "o", ".", "the path to which to write generated docs")
	return cmd
}

func runDocsCommand(cmd *cobra.Command, args []string) error {
	ctxDir, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}

	inputDir, err := cmd.Flags().GetString("input")
	if err != nil {
		return err
	}

	pattern, err := cmd.Flags().GetString("pattern")
	if err != nil {
		return err
	}

	outputDir, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	ctxInputDir := filepath.Join(ctxDir, inputDir)
	ctxOutputDir := filepath.Join(ctxDir, outputDir)
	err = doublestar.GlobWalk(os.DirFS(ctxInputDir), pattern, func(path string, info fs.DirEntry) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".proto" {
			return nil
		}

		var args []string
		var pathArgs []string
		pathArgs = append(pathArgs, ctxDir)
		pathArgs = append(pathArgs, filepath.Join(os.Getenv("GOPATH"), "src/github.com/gogo/protobuf"))
		protoPath := strings.Join(pathArgs, ":")
		args = append(args, "-I", protoPath)
		var specArgs []string
		specArgs = append(specArgs, filepath.Dir(filepath.Join(ctxOutputDir, path)))
		spec := strings.Join(specArgs, ",")
		args = append(args, fmt.Sprintf("--doc_out=%s", spec))
		var optArgs []string
		optArgs = append(optArgs, "markdown")
		optArgs = append(optArgs, fmt.Sprintf("%s.md", info.Name()[:len(info.Name())-len(filepath.Ext(info.Name()))]))
		opt := strings.Join(optArgs, ",")
		args = append(args, fmt.Sprintf("--doc_opt=%s", opt))
		ctxPath := filepath.Join(ctxDir, path)
		args = append(args, ctxPath)
		return exec.Run("protoc", exec.WithEnv(os.Environ()...), exec.WithDir(ctxDir), exec.WithArgs(args...))
	})
	if err != nil {
		return err
	}
	return nil
}
