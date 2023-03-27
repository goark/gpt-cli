package chat

import (
	"context"
	"strings"

	"github.com/goark/errs"
	openai "github.com/sashabaranov/go-openai"
)

// Request requesta OpenAI Chat completion, and returns response message. (REST access)
func (cctx *ChatContext) Request(ctx context.Context, msg string) (string, error) {
	resp, err := cctx.requestRaw(ctx, msg)
	if err != nil {
		return "", errs.Wrap(err)
	}
	if len(resp.Choices) == 0 {
		return "", nil
	}
	return resp.Choices[0].Message.Content, nil
}

func (cctx *ChatContext) requestRaw(ctx context.Context, msg string) (openai.ChatCompletionResponse, error) {
	msg = strings.TrimSpace(msg)
	if len(msg) > 0 {
		cctx.profile.Messages = append(cctx.profile.Messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: msg})
	}
	cctx.Logger().Info().Interface("request", cctx.profile).Send()
	resp, err := cctx.Client().CreateChatCompletion(ctx, cctx.profile)
	if err != nil {
		err = errs.Wrap(err, errs.WithContext("request", cctx.profile))
		cctx.Logger().Error().Interface("error", err).Send()
		return resp, err
	}
	cctx.Logger().Info().Interface("response", resp).Send()
	return resp, nil
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
