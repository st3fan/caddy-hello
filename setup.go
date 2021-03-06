// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package hello

import (
	"fmt"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("hello", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func parse(c *caddy.Controller) (string, error) {
	if !c.NextArg() {
		return "", c.ArgErr()
	}
	c.Next()
	return c.Val(), nil
}

func setup(c *caddy.Controller) error {
	c.OnStartup(func() error {
		fmt.Println("startup: caddy-hello")
		return nil
	})

	path, err := parse(c)
	if err != nil {
		return err
	}

	httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return &Hello{Next: next, Path: path}
	})

	return nil
}
