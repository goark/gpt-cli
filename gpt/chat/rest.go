package chat

import (
	"context"
	"io"

	"github.com/goark/errs"
	"github.com/sashabaranov/go-openai"
)

func (cctx *ChatContext) rest(ctx context.Context, client *openai.Client, w io.Writer) (string, error) {
	cctx.prepare.Stream = false
	cctx.Logger().Info().Interface("request", cctx.prepare).Send()
	resp, err := client.CreateChatCompletion(ctx, cctx.prepare)
	if err != nil {
		err = errs.Wrap(err, errs.WithContext("request", cctx.prepare))
		cctx.Logger().Error().Interface("error", err).Send()
		return "", err
	}
	cctx.Logger().Info().Interface("response", resp).Send()

	if len(resp.Choices) == 0 {
		return "", nil
	}
	resText := resp.Choices[0].Message.Content
	if _, err := io.WriteString(w, resText); err != nil {
		err = errs.Wrap(err)
		cctx.Logger().Error().Interface("error", err).Send()
		return "", err
	}
	return resText, nil
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
