package facade

import (
	"github.com/goark/gocli/rwi"
	"github.com/goark/gpt-cli/gpt/chat"
	"github.com/spf13/cobra"
)

// newVersionCmd returns cobra.Command instance for show sub-command
func newHistoryCmd(ui *rwi.RWI) *cobra.Command {
	historyCmd := &cobra.Command{
		Use:     "history",
		Aliases: []string{"hist", "h"},
		Short:   "Print chat history",
		Long:    "Print chat history.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// local options
			histPath, err := cmd.Flags().GetString("history-file")
			if err != nil {
				return debugPrint(ui, err)
			}
			userName, err := cmd.Flags().GetString("user-name")
			if err != nil {
				return debugPrint(ui, err)
			}
			assistantName, err := cmd.Flags().GetString("assistant-name")
			if err != nil {
				return debugPrint(ui, err)
			}

			// Output history
			if err := chat.OutputHistory(histPath, userName, assistantName, ui.Writer()); err != nil {
				return debugPrint(ui, err)
			}
			return nil
		},
	}
	historyCmd.Flags().StringP("history-file", "f", "", "Path of history file (JSON format)")
	historyCmd.Flags().StringP("user-name", "u", "", "User name (display name)")
	historyCmd.Flags().StringP("assistant-name", "a", "", "Assistant name (display name)")

	return historyCmd
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
