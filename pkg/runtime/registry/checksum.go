// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"github.com/atomix/runtime-api/pkg/errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func newChecksum(plugin *Plugin) *pluginChecksum {
	return &pluginChecksum{
		plugin: plugin,
	}
}

type pluginChecksum struct {
	plugin *Plugin
}

func (c *pluginChecksum) check(url string) error {
	localChecksum, err := c.compute()
	if err != nil {
		return err
	}
	remoteChecksum, err := c.fetch(url)
	if err != nil {
		return err
	}
	if localChecksum != remoteChecksum {
		return errors.NewFault("checksum for plugin %s did not match the checksum in %s", c.plugin, url)
	}
	return nil
}

func (c *pluginChecksum) fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		sum, file := split[0], split[1]
		if file == c.plugin.filename() {
			return sum, nil
		}
	}
	return "", errors.NewNotFound("could not retrieve checksum for %s from %s", c.plugin, url)
}

func (c *pluginChecksum) compute() (string, error) {
	reader, err := os.Open(c.plugin.Path)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	sha := sha256.New()
	buf := make([]byte, bufferSize)
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			return hex.EncodeToString(sha.Sum(nil)), nil
		}
		if err != nil {
			return "", err
		}

		_, err = sha.Write(buf)
		if err != nil {
			return "", err
		}
	}
}
