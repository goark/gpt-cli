package facade

import (
	"path/filepath"
	"sort"

	"github.com/goark/errs"
	"github.com/goark/gocli/cache"
	"github.com/goark/gpt-cli/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type options struct {
	APIKey   string
	CacheDir string
	Logger   *zerolog.Logger
}

func getOptions() (*options, error) {
	logger, err := logger.New(
		logger.LevelFrom(viper.GetString("log-level")),
		viper.GetString("log-dir"),
	)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &options{
		APIKey:   viper.GetString("api-key"),
		CacheDir: cache.Dir(Name),
		Logger:   logger,
	}, nil
}

func getFiles(ss []string) ([]string, error) {
	paths := map[string]bool{}
	for _, s := range ss {
		pp, err := filepath.Glob(s)
		if err != nil {
			return nil, errs.Wrap(err, errs.WithContext("path", s))
		}
		for _, p := range pp {
			paths[p] = true
		}
	}
	if len(paths) > 0 {
		plist := make([]string, 0, len(paths))
		for k := range paths {
			plist = append(plist, k)
		}
		sort.Strings(plist)
		return plist, nil
	}
	return []string{}, nil
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
