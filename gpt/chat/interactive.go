package chat

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gpt-cli/ecode"
	"github.com/nyaosorg/go-readline-ny"
	openai "github.com/sashabaranov/go-openai"
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
	cctx.profile.Stream = true
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
		cctx.profile.Messages = append(cctx.profile.Messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: text})
		resText, err := cctx.stream(ctx, client, w)
		if err != nil {
			return errs.Wrap(err)
		}
		cctx.profile.Messages = append(cctx.profile.Messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: resText})
	}
	return cctx.Save()
}

func (cctx *ChatContext) stream(ctx context.Context, client *openai.Client, w io.Writer) (string, error) {
	cctx.Logger().Info().Interface("request", cctx.profile).Send()
	stream, err := client.CreateChatCompletionStream(ctx, cctx.profile)
	if err != nil {
		err = errs.Wrap(err)
		cctx.Logger().Error().Interface("error", err).Send()
		return "", err
	}
	defer stream.Close()

	builder := strings.Builder{}
	fmt.Fprintln(w)
	for {
		resp, err := stream.Recv()
		if err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			err = errs.Wrap(ecode.ErrStream, errs.WithCause(err))
			cctx.Logger().Error().Interface("error", err).Send()
			return "", err
		}
		cctx.Logger().Info().Interface("response", resp).Send()
		if len(resp.Choices) > 0 {
			if delta := resp.Choices[0].Delta.Content; len(delta) > 0 {
				if _, err := builder.WriteString(delta); err != nil {
					err = errs.Wrap(err)
					cctx.Logger().Error().Interface("error", err).Send()
					return "", err
				}
				if _, err := io.WriteString(w, delta); err != nil {
					err = errs.Wrap(err)
					cctx.Logger().Error().Interface("error", err).Send()
					return "", err
				}
			}
		}
	}
	fmt.Fprintln(w)
	return builder.String(), nil
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
