package logger

import (
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestDate(t *testing.T) {
	want := strings.Join([]string{"nop", "error", "warn", "info", "debug", "trace"}, "|")
	if got := strings.Join(LevelList(), "|"); got != want {
		t.Errorf("LevelList() is \"%v\", want \"%v\"", got, want)
	}
}

func TestLevelFrom(t *testing.T) {
	testCases := []struct {
		s    string
		want zerolog.Level
	}{
		{s: "", want: zerolog.NoLevel},
		{s: "error", want: zerolog.ErrorLevel},
		{s: "warn", want: zerolog.WarnLevel},
		{s: "info", want: zerolog.InfoLevel},
		{s: "debug", want: zerolog.DebugLevel},
		{s: "trace", want: zerolog.TraceLevel},
		{s: "foo", want: zerolog.NoLevel},
	}
	for _, tc := range testCases {
		if got := LevelFrom(tc.s); got != tc.want {
			t.Errorf("LevelFrom(\"%s\") is [%v], want [%v]", tc.s, got, tc.want)
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
