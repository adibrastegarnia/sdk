// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package build

var (
	info       string
	version    string
	commitHash string
	modHash    string
	modSum     string
)

//go:generate go run github.com/atomix/atomix-runtime/build/version

func Version() string {
	return version
}

func CommitHash() string {
	return commitHash
}

func ModHash() string {
	return modHash
}

func ModSum() string {
	return modSum
}

type Info struct {
	Version    string `yaml:"version"`
	CommitHash string `yaml:"commitHash"`
	ModHash    string `yaml:"modHash"`
	ModSum     string `yaml:"modSum"`
}
