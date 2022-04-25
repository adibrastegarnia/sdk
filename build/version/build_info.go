// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/atomix/atomix-runtime/internal/exec"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	version, err := getVersion()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commitHash, err := getCommitHash()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	modHash, err := getModHash()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	modSum, err := getModSum()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := os.Create("vars.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	out.WriteString("// SPDX-FileCopyrightText: 2022-present Intel Corporation\n//\n// SPDX-License-Identifier: Apache-2.0\n\n")
	out.WriteString("package build\n\n")
	out.WriteString("func init() {\n")
	out.WriteString(fmt.Sprintf("\tversion = \"%s\"\n", version))
	out.WriteString(fmt.Sprintf("\tcommitHash = \"%s\"\n", commitHash))
	out.WriteString(fmt.Sprintf("\tmodHash = \"%s\"\n", modHash))
	out.WriteString(fmt.Sprintf("\tmodSum = \"%s\"\n", modSum))
	out.WriteString("}\n")
}

func getVersion() (string, error) {
	output, err := exec.Output("git", exec.WithArgs("describe", "--tags"))
	if err != nil {
		return "", err
	}
	return strings.Trim(string(output), "\n "), nil
}

func getCommitHash() (string, error) {
	output, err := exec.Output("git", exec.WithArgs("rev-list", "-1", "HEAD"))
	if err != nil {
		return "", err
	}
	return strings.Trim(string(output), "\n "), nil
}

func getModHash() (string, error) {
	goModBytes, err := ioutil.ReadFile("../../go.mod")
	if err != nil {
		return "", err
	}

	goModSha := sha256.New()
	reader := ioutil.NopCloser(bytes.NewReader(goModBytes))
	defer reader.Close()

	_, err = io.Copy(goModSha, reader)
	if err != nil {
		return "", err
	}

	shaBytes := goModSha.Sum(nil)
	hash := fmt.Sprintf("%x", shaBytes)
	return hash, nil
}

func getModSum() (string, error) {
	goModBytes, err := ioutil.ReadFile("../../go.mod")
	if err != nil {
		return "", err
	}

	goModSha := sha256.New()
	reader := ioutil.NopCloser(bytes.NewReader(goModBytes))
	defer reader.Close()

	_, err = io.Copy(goModSha, reader)
	if err != nil {
		return "", err
	}

	shaBytes := goModSha.Sum(nil)
	modSha := sha256.New()
	fmt.Fprintf(modSha, "%x  %s\n", shaBytes, "go.mod")
	checksum := fmt.Sprintf("h1:%s", base64.StdEncoding.EncodeToString(modSha.Sum(nil)))
	return checksum, nil
}
