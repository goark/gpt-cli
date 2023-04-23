package facade

import (
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
			clipboardFlag, err := cmd.Flags().GetBool("clipboard")
			if err != nil {
				return debugPrint(ui, err)
			}
			prepPath, err := cmd.Flags().GetString("prepare-file")
			if err != nil {
				return debugPrint(ui, err)
			}
			attaches, err := cmd.Flags().GetStringSlice("attach-file")
			if err != nil {
				return debugPrint(ui, err)
			}
			paths, err := getFiles(attaches)
			if err != nil {
				return debugPrint(ui, err)
			}
			msg, err := cmd.Flags().GetString("message")
			if err != nil {
				return debugPrint(ui, err)
			}
			savePath, err := cmd.Flags().GetString("output-file")
			if err != nil {
				return debugPrint(ui, err)
			}
			streamMode, err := cmd.Flags().GetBool("stream-mode")
			if err != nil {
				return debugPrint(ui, err)
			}

			// create Chat context
			cctx, err := chat.New(opts.APIKey, opts.CacheDir, opts.Logger, prepPath, savePath)
			if err != nil {
				opts.Logger.Error().Interface("error", errs.Wrap(err)).Send()
				return debugPrint(ui, err)
			}

			var msgs []string = []string{}
			// message from command-line
			if len(msg) > 0 {
				msgs = append(msgs, msg)
			}
			// message from clipboard
			if clipboardFlag {
				text, err := clipboard.ReadAll()
				if err != nil {
					return debugPrint(ui, err)
				}
				if len(text) > 0 {
					msgs = append(msgs, text)
				}
			}
			// messages from attached files
			for _, path := range paths {
				msg, err := chat.AttachFile(path)
				if err != nil {
					return debugPrint(ui, err)
				}
				msgs = append(msgs, msg)
			}

			// kicking single mode
			if err := cctx.Request(cmd.Context(), streamMode, msgs, ui.Writer()); err != nil {
				return debugPrint(ui, err)
			}
			if len(cctx.SavePath()) > 0 {
				return ui.Outputln("\nsave to", cctx.SavePath())
			}
			return nil
		},
	}
	chatCmd.Flags().BoolP("clipboard", "c", false, "Input message from clipboard")
	chatCmd.Flags().BoolP("stream-mode", "", false, "Output from GPT by Stream mode")
	chatCmd.Flags().StringP("message", "m", "", "Chat message")
	chatCmd.Flags().StringP("prepare-file", "p", "", "Path of prepare file (JSON format)")
	chatCmd.Flags().StringSliceP("attach-file", "a", nil, "Path of attach files (text file only)")
	chatCmd.Flags().StringP("output-file", "o", "", "Path of save file (JSON format)")

	chatCmd.AddCommand(
		newHistoryCmd(ui),
		newInteractiveCmd(ui),
	)

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
