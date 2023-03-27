package gpt

import (
	"github.com/goark/errs"
	"github.com/goark/gpt-cli/ecode"
	"github.com/rs/zerolog"
)

type GPTContext struct {
	apiKey string
	logger *zerolog.Logger
}

// New function creates APIContext instance.
func New(apiKey string, logger *zerolog.Logger) (*GPTContext, error) {
	if len(apiKey) == 0 {
		return nil, errs.Wrap(ecode.ErrAPIKey)
	}
	return &GPTContext{apiKey: apiKey, logger: logger}, nil
}

// ApiKey method returns API key string.
func (gctx *GPTContext) APIKey() string {
	return gctx.apiKey
}

// Logger method returns logger.
func (gctx *GPTContext) Logger() *zerolog.Logger {
	return gctx.logger
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
