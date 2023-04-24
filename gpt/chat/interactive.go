package chat

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gpt-cli/ecode"
	"github.com/hymkor/go-multiline-ny"
	"github.com/nyaosorg/go-readline-ny"
)

// Interactive method is chatting in interactive mode (stream access).
func (cctx *ChatContext) Interactive(ctx context.Context, w io.Writer) error {
	if cctx == nil {
		return errs.Wrap(ecode.ErrNullPointer)
	}
	client := cctx.Client()
	editor := readline.Editor{
		Prompt: func() (int, error) { return fmt.Print("\nChat>") },
	}
	fmt.Fprintln(w, "Input 'q' or 'quit' to stop")
	cctx.prepare.Stream = true
	for {
		text, err := editor.ReadLine(ctx)
		if err != nil {
			return errs.Wrap(err)
		}
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			continue
		}
		if strings.EqualFold(text, "q") || strings.EqualFold(text, "quit") {
			break
		}
		_ = cctx.AppendUserMessages([]string{text})
		resText, err := cctx.stream(ctx, client, w)
		if err != nil {
			return errs.Wrap(err)
		}
		_ = cctx.AppendAssistantMessages([]string{resText})
	}
	return cctx.Save()
}

// InteractiveMulti method is chatting in interactive mode with multiline editing (stream access).
func (cctx *ChatContext) InteractiveMulti(ctx context.Context, w io.Writer) error {
	if cctx == nil {
		return errs.Wrap(ecode.ErrNullPointer)
	}
	client := cctx.Client()
	var editor multiline.Editor
	editor.SetPrompt(func(w io.Writer, lnum int) (int, error) {
		return fmt.Fprintf(w, "Chat:%2d>", lnum+1)
	})

	fmt.Fprintln(w, "Input 'Ctrl+J' or 'Ctrl+Enter' to submit message")
	fmt.Fprintln(w, "Input 'Ctrl+D' with no chars to stop")
	fmt.Fprintln(w, "      or input text \"q\" or \"quit\" and submit to stop")
	cctx.prepare.Stream = true
	for {
		lines, err := editor.Read(ctx)
		if err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return errs.Wrap(err)
		}
		if len(lines) == 0 {
			fmt.Fprintln(w)
			continue
		}
		text := strings.TrimSpace(strings.Join(lines, "\n"))
		if len(text) == 0 {
			fmt.Fprintln(w)
			continue
		}
		if strings.EqualFold(text, "q") || strings.EqualFold(text, "quit") {
			break
		}

		_ = cctx.AppendUserMessages([]string{text})
		resText, err := cctx.stream(ctx, client, w)
		if err != nil {
			return errs.Wrap(err)
		}
		fmt.Fprintln(w)
		_ = cctx.AppendAssistantMessages([]string{resText})
	}
	return cctx.Save()
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
