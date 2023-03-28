package chat

import (
	"bytes"
	"strings"
	"testing"
)

var (
	history = `{
  "model": "gpt-3.5-turbo-0301",
  "max_tokens": 256,
  "temperature": 0.7,
  "top_p": 1,
  "n": 1,
  "messages": [
    {
    "role": "user",
    "content": "hello"
    },
    {
    "role": "assistant",
    "content": "Hi! How can I assist you today?"
    }
  ]
}`
	output = `# Chat with GPT

- ` + "`model`" + `: gpt-3.5-turbo-0301
- ` + "`max_tokens`" + `: 256
- ` + "`temperature`" + `: 0.7
- ` + "`top_p`" + `: 1
- ` + "`n`" + `: 1

## Spiegel

hello

## ChatGPT

Hi! How can I assist you today?
`
)

func TestOutputHistory(t *testing.T) {
	buf := &bytes.Buffer{}
	if err := OutputHistory(strings.NewReader(history), buf, "Spiegel", "ChatGPT"); err != nil {
		t.Errorf("OutputHistory() is [%v], want <nil>", err)
	} else {
		if got := buf.String(); got != output {
			t.Errorf("OutputHistory() is [%v], want [%v]", got, output)
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
