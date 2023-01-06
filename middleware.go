package metricmware

/*
 * Copyright (c) 2023. Seth Osher.  All Rights Reserved.
 * MIT License
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
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

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/zserge/metric"
	"strings"
)

var (
	metricHandler = fasthttpadaptor.NewFastHTTPHandlerFunc(metric.Handler(metric.Exposed).ServeHTTP)
)

// New creates a new metrics middleware under [prefix]/debug/metrics using the prefix from config if specified
func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		path := c.Path()

		// We are only interested in /debug/metrics routes
		if len(path) < len("/debug/metrics") || !strings.HasPrefix(path, cfg.Prefix+"/debug/metrics") {
			return c.Next()
		}
		switch path {
		case cfg.Prefix + "/debug/metrics", cfg.Prefix + "/debug/metrics/":
			metricHandler(c.Context())
		default:
			return c.Redirect("/debug/metrics", fiber.StatusFound)
		}
		return nil
	}
}
