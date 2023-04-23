package facade

import (
	"github.com/goark/errs"
	"github.com/goark/gocli/rwi"
	"github.com/goark/gpt-cli/gpt/chat"
	"github.com/spf13/cobra"
)

// newVersionCmd returns cobra.Command instance for show sub-command
func newInteractiveCmd(ui *rwi.RWI) *cobra.Command {
	interactiveCmd := &cobra.Command{
		Use:     "interactive",
		Aliases: []string{"i"},
		Short:   "Interactive mode",
		Long:    "Interactive mode in chat.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Global options
			opts, err := getOptions()
			if err != nil {
				return debugPrint(ui, err)
			}
			prepPath, err := cmd.Flags().GetString("prepare-file")
			if err != nil {
				return debugPrint(ui, err)
			}
			savePath, err := cmd.Flags().GetString("output-file")
			if err != nil {
				return debugPrint(ui, err)
			}
			multiLine, err := cmd.Flags().GetBool("multi-line")
			if err != nil {
				return debugPrint(ui, err)
			}

			// create Chat context
			cctx, err := chat.New(opts.APIKey, opts.CacheDir, opts.Logger, prepPath, savePath)
			if err != nil {
				opts.Logger.Error().Interface("error", errs.Wrap(err)).Send()
				return debugPrint(ui, err)
			}

			// kicking interactive mode
			if multiLine {
				if err := cctx.InteractiveMulti(cmd.Context(), ui.Writer()); err != nil {
					return debugPrint(ui, err)
				}
			} else {
				if err := cctx.Interactive(cmd.Context(), ui.Writer()); err != nil {
					return debugPrint(ui, err)
				}
			}
			if len(cctx.SavePath()) > 0 {
				return ui.Outputln("\nsave to", cctx.SavePath())
			}
			return nil
		},
	}
	interactiveCmd.Flags().StringP("prepare-file", "p", "", "Path of prepare file (JSON format)")
	interactiveCmd.Flags().StringP("output-file", "o", "", "Path of save file (JSON format)")
	interactiveCmd.Flags().BoolP("multi-line", "m", false, "Editing with multi-line mode")

	return interactiveCmd
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
