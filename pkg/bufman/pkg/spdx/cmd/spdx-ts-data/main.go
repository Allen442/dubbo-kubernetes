// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"regexp"
	"strings"
)

import (
	"github.com/spf13/cobra"
)

import (
	"github.com/apache/dubbo-kubernetes/pkg/bufman/pkg/app"
	"github.com/apache/dubbo-kubernetes/pkg/bufman/pkg/app/appcmd"
	"github.com/apache/dubbo-kubernetes/pkg/bufman/pkg/spdx"
)

const (
	programName = "spdx-ts-data"
)

var (
	numberRegexp     = regexp.MustCompile(`^(?:\d+|\d+\.\d+)$`)
	identifierRegexp = regexp.MustCompile(`^[a-zA-Z$_][a-zA-Z0-9$_]*$`)
)

func main() {
	appcmd.Main(context.Background(), newCommand())
}

func newCommand() *appcmd.Command {
	return &appcmd.Command{
		Use:  programName,
		Args: cobra.NoArgs,
		Run:  run,
	}
}

func run(ctx context.Context, container app.Container) error {
	licenseInfos, err := spdx.GetLicenseInfos(ctx)
	if err != nil {
		return err
	}
	data, err := getTSFileData(licenseInfos)
	if err != nil {
		return err
	}
	_, err = container.Stdout().Write(data)
	return err
}

func getTSFileData(licenseInfos []*spdx.LicenseInfo) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	_, _ = buffer.WriteString(`// Code generated by ` + programName + `. DO NOT EDIT.

export const LICENSES = {
`)

	for _, licenseInfo := range licenseInfos {
		id, err := formatString(licenseInfo.ID)
		if err != nil {
			return nil, err
		}
		name, err := formatString(licenseInfo.Name)
		if err != nil {
			return nil, err
		}
		propName := strings.ToLower(licenseInfo.ID)
		if shouldQuotePropertyName(propName) {
			quoted, err := json.Marshal(propName)
			if err != nil {
				return nil, err
			}
			propName = string(quoted)
		}
		_, _ = buffer.WriteString(`  ` + propName + `: {
    ID: ` + id + ` as const,
    name: ` + name + ` as const,
  },
`)
	}

	_, _ = buffer.WriteString(`};
`)
	return buffer.Bytes(), nil
}

// formatsString matches https://prettier.io/docs/en/options.html#quotes default behavior.
func formatString(str string) (string, error) {
	var numDoubleQuotes, numSingleQuotes int
	for _, r := range str {
		switch r {
		case '"':
			numDoubleQuotes++
		case '\'':
			numSingleQuotes++
		}
	}
	b, err := json.Marshal(str)
	if err != nil {
		return "", err
	}
	if numDoubleQuotes <= numSingleQuotes {
		return string(b), nil
	}
	var inEscape bool
	var sb strings.Builder
	sb.WriteByte('\'')
	for _, r := range str {
		if inEscape {
			inEscape = false
			if r != '"' {
				sb.WriteRune('\\')
			}
			sb.WriteRune(r)
		} else if r == '\\' {
			inEscape = true
		} else {
			sb.WriteRune(r)
		}
	}
	sb.WriteByte('\'')
	return sb.String(), nil
}

// shouldQuotePropertyName returns where the name should be quoted as a property key.
// This follows the rules for 'as-needed' here: https://prettier.io/docs/en/options.html#quote-props
func shouldQuotePropertyName(name string) bool {
	if numberRegexp.MatchString(name) || identifierRegexp.MatchString(name) {
		return false
	}
	return true
}
