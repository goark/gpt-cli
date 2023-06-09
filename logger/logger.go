package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/goark/errs"
	"github.com/goark/gocli/cache"
	"github.com/rs/zerolog"
)

// New function returns new zerolog.Logger instance.
func New(lvl zerolog.Level, dir string) (*zerolog.Logger, error) {
	logger := zerolog.Nop()
	if lvl == zerolog.NoLevel {
		return &logger, nil
	}
	logpath := getPath(dir)
	if file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600); err != nil {
		return &logger, errs.Wrap(err, errs.WithContext("logpath", logpath))
	} else {
		logger = zerolog.New(file).Level(lvl).With().Timestamp().Logger()
	}
	return &logger, nil
}

// DefaultLogDir function returns default log directory ($XDG_CACHE_HOME/appName/)
func DefaultLogDir(appName string) string {
	return cache.Dir(appName)
}

func getPath(dir string) string {
	if len(dir) == 0 {
		dir = "."
	}
	_ = os.MkdirAll(dir, 0700)
	return filepath.Join(dir, fmt.Sprintf("access.%s.log", time.Now().Local().Format("20060102")))
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
