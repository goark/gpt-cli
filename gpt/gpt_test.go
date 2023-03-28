package gpt

import (
	"errors"
	"testing"

	"github.com/goark/gpt-cli/ecode"
	"github.com/rs/zerolog"
)

var notLogger = zerolog.Nop()

func TestNew(t *testing.T) {
	testCases := []struct {
		apiKey       string
		cacheDir     string
		logger       *zerolog.Logger
		cacheDirWant string
		err          error
	}{
		{apiKey: "", cacheDir: "", logger: &notLogger, cacheDirWant: ".", err: ecode.ErrAPIKey},
		{apiKey: "foo", cacheDir: "", logger: &notLogger, cacheDirWant: ".", err: nil},
		{apiKey: "foo", cacheDir: "bar", logger: &notLogger, cacheDirWant: "bar", err: nil},
	}
	for _, tc := range testCases {
		if cctx, err := New(tc.apiKey, tc.cacheDir, tc.logger); !errors.Is(err, tc.err) {
			t.Errorf("New() is [%v], want [%v]", err, tc.err)
		} else if err == nil {
			if got := cctx.CacheDir(); got != tc.cacheDirWant {
				t.Errorf("CacheDir() is [%v], want [%v]", got, tc.cacheDirWant)
			}
		}
	}
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
