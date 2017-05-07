// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package hello

import (
	"fmt"
	"net/http"

	"github.com/mholt/caddy/caddyhttp/httpserver"
)

type Hello struct {
	Next httpserver.Handler
	Path string
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.URL.Path == h.Path {
		fmt.Fprintf(w, "Hello, world!")
		return 0, nil
	}
	return h.Next.ServeHTTP(w, r)
}
