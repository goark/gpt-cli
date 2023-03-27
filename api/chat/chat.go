package chat

import (
	"context"
	"encoding/json"
	"os"

	"github.com/goark/errs"
	"github.com/goark/gpt-cli/api"
	openai "github.com/sashabaranov/go-openai"
)

// ChatContext is context data for chat
type ChatContext struct {
	*api.APIContext
	profile openai.ChatCompletionRequest
}

// New function create new ChatContext instance.
func New(apiKey string, profilePath string) (*ChatContext, error) {
	actx, err := api.New(apiKey)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	profile := openai.ChatCompletionRequest{}
	if len(profilePath) > 0 {
		file, err := os.Open(profilePath)
		if err != nil {
			return nil, errs.Wrap(err, errs.WithContext("profilePath", profilePath))
		}
		defer file.Close()
		if err := json.NewDecoder(file).Decode(&profile); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	if len(profile.Model) == 0 {
		profile.Model = openai.GPT3Dot5Turbo0301
	}
	if profile.Messages == nil {
		profile.Messages = []openai.ChatCompletionMessage{}
	}
	return &ChatContext{APIContext: actx, profile: profile}, nil
}

// RequestRaw requesta OpenAI Chat completion, and returns raw response. (REST access)
func (cctx *ChatContext) RequestRaw(ctx context.Context, msg string) (openai.ChatCompletionResponse, error) {
	cctx.profile.Messages = append(cctx.profile.Messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: msg})
	resp, err := openai.NewClient(cctx.APIKey()).CreateChatCompletion(ctx, cctx.profile)
	return resp, errs.Wrap(err, errs.WithContext("request", cctx.profile))
}

// Request requesta OpenAI Chat completion, and returns response message. (REST access)
func (cctx *ChatContext) Request(ctx context.Context, msg string) (string, error) {
	cctx.profile.Messages = append(cctx.profile.Messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: msg})
	resp, err := openai.NewClient(cctx.APIKey()).CreateChatCompletion(ctx, cctx.profile)
	if err != nil {
		return "", errs.Wrap(err, errs.WithContext("request", cctx.profile))
	}
	if len(resp.Choices) == 0 {
		return "", nil
	}
	return resp.Choices[0].Message.Content, nil
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
