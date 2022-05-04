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

func GetGoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "go",
		Aliases: []string{"golang"},
		Args:    cobra.ExactArgs(1),
		RunE:    runGoCommand,
	}
	cmd.Flags().StringP("input", "i", ".", "the path to the root of the Protobuf sources")
	cmd.Flags().StringP("pattern", "f", "**/*.proto", "a pattern by which to filter Protobuf sources")
	cmd.Flags().StringP("output", "o", ".", "the path to which to write generated Go sources")
	cmd.Flags().StringP("package", "p", "", "the base Go path for generated sources")
	return cmd
}

func runGoCommand(cmd *cobra.Command, args []string) error {
	ctxDir, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}

	inputDir, err := cmd.Flags().GetString("input")
	if err != nil {
		return err
	}

	inputPattern, err := cmd.Flags().GetString("pattern")
	if err != nil {
		return err
	}

	outputDir, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	}

	outputPackage, err := cmd.Flags().GetString("package")
	if err != nil {
		return err
	}

	importOverrides := make(map[string]string)
	importOverrides["google/protobuf/any.proto"] = "github.com/gogo/protobuf/types"
	importOverrides["google/protobuf/timestamp.proto"] = "github.com/gogo/protobuf/types"
	importOverrides["google/protobuf/duration.proto"] = "github.com/gogo/protobuf/types"

	ctxInputDir := filepath.Join(ctxDir, inputDir)
	err = doublestar.GlobWalk(os.DirFS(ctxInputDir), inputPattern, func(path string, info fs.DirEntry) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".proto" {
			return nil
		}

		packagePath := filepath.Dir(path)
		goPath := filepath.Join(outputPackage, packagePath)
		importOverrides[path] = goPath
		return nil
	})

	ctxOutputDir := filepath.Join(ctxDir, outputDir)
	err = doublestar.GlobWalk(os.DirFS(ctxInputDir), inputPattern, func(path string, info fs.DirEntry) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".proto" {
			return nil
		}

		var args []string
		var pathArgs []string
		pathArgs = append(pathArgs, ctxDir)
		pathArgs = append(pathArgs, ctxInputDir)
		pathArgs = append(pathArgs, filepath.Join(os.Getenv("GOPATH"), "src/github.com/gogo/protobuf"))
		args = append(args, "-I", strings.Join(pathArgs, ":"))
		var specArgs []string
		var overrideArgs []string
		for protoPath, goPath := range importOverrides {
			overrideArgs = append(overrideArgs, fmt.Sprintf("M%s=%s", protoPath, goPath))
		}
		overrides := strings.Join(overrideArgs, ",")
		specArgs = append(specArgs, overrides)
		importPath := filepath.Join(outputPackage, filepath.Dir(path))
		specArgs = append(specArgs, fmt.Sprintf("import_path=%s", importPath))
		specArgs = append(specArgs, fmt.Sprintf("plugins=grpc:%s", ctxOutputDir))
		spec := strings.Join(specArgs, ",")
		args = append(args, fmt.Sprintf("--gogofaster_out=%s", spec))
		args = append(args, path)
		return exec.Run("protoc", exec.WithEnv(os.Environ()...), exec.WithDir(ctxDir), exec.WithArgs(args...))
	})
	if err != nil {
		return err
	}
	return nil
}
