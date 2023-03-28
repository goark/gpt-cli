package chat

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/goark/errs"
	"github.com/sashabaranov/go-openai"
)

func OutputHistory(r io.Reader, w io.Writer, userName, assistantName string) error {
	hist := openai.ChatCompletionRequest{}
	if err := json.NewDecoder(r).Decode(&hist); err != nil {
		return errs.Wrap(err)
	}

	// Output
	fmt.Fprintln(w, "# Chat with GPT")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "- `model`:", hist.Model)
	if hist.MaxTokens != 0 {
		fmt.Fprintln(w, "- `max_tokens`:", hist.MaxTokens)
	}
	if hist.Temperature != 0 {
		fmt.Fprintln(w, "- `temperature`:", hist.Temperature)
	}
	if hist.TopP != 0 {
		fmt.Fprintln(w, "- `top_p`:", hist.TopP)
	}
	if hist.N != 0 {
		fmt.Fprintln(w, "- `n`:", hist.N)
	}
	for _, msg := range hist.Messages {
		role := msg.Role
		switch {
		case role == openai.ChatMessageRoleUser && len(userName) > 0:
			role = userName
		case role == openai.ChatMessageRoleAssistant && len(assistantName) > 0:
			role = assistantName
		}
		fmt.Fprintln(w)
		fmt.Fprintln(w, "##", role)
		fmt.Fprintln(w)
		fmt.Fprintln(w, msg.Content)
	}
	return nil
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
