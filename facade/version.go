package facade

import (
	"strings"

	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

var versionStrings = []string{ //output message of version
	Name + " " + Version,
	"repository: https://github.com/goark/gpt-cli",
}

func getVersion() string {
	return strings.Join(versionStrings, "\n")
}

// newVersionCmd returns cobra.Command instance for show sub-command
func newVersionCmd(ui *rwi.RWI) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver", "v"},
		Short:   "Print the version number",
		Long:    "Print the version number of " + Name,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.OutputErrln(getVersion())
		},
	}

	return versionCmd
}

/* MIT License
 *
 * Copyright 2023 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
