package facade

import (
	"io"

	"github.com/atotto/clipboard"
	"github.com/goark/errs"
	"github.com/goark/gocli/rwi"
	"github.com/goark/gpt-cli/gpt/chat"
	"github.com/spf13/cobra"
)

// newVersionCmd returns cobra.Command instance for show sub-command
func newChatCmd(ui *rwi.RWI) *cobra.Command {
	chatCmd := &cobra.Command{
		Use:     "chat",
		Aliases: []string{"c"},
		Short:   "Chat with GPT-x",
		Long:    "Chat with GPT-x, input from standard input.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Global options
			opts, err := getOptions()
			if err != nil {
				return debugPrint(ui, err)
			}
			// local options
			interactivMode, err := cmd.Flags().GetBool("interactive")
			if err != nil {
				return debugPrint(ui, err)
			}
			clipboardFlag, err := cmd.Flags().GetBool("clipboard")
			if err != nil {
				return debugPrint(ui, err)
			}
			profilePath, err := cmd.Flags().GetString("profile")
			if err != nil {
				return debugPrint(ui, err)
			}

			// create Chat context
			cctx, err := chat.New(opts.APIKey, opts.Logger, profilePath)
			if err != nil {
				opts.Logger.Error().Interface("error", errs.Wrap(err)).Send()
				return debugPrint(ui, err)
			}

			// interactive mode
			if interactivMode {
				return debugPrint(ui, cctx.Interactive(cmd.Context(), ui.Writer()))
			}

			// single mode
			var text string
			if clipboardFlag {
				text, err = clipboard.ReadAll()
				if err != nil {
					return debugPrint(ui, err)
				}
			} else {
				b, err := io.ReadAll(ui.Reader())
				if err != nil {
					return debugPrint(ui, err)
				}
				text = string(b)
			}
			respMsg, err := cctx.Request(cmd.Context(), text)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.Outputln(respMsg))
		},
	}
	chatCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")
	chatCmd.Flags().BoolP("clipboard", "c", false, "Input message from clipboard")
	chatCmd.Flags().StringP("profile", "p", "", "Path of Profile file (JSON format)")

	return chatCmd
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
