// // Copyright 2024 The KeepShare Authors. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file.

package rapidgator

import (
	"context"
	"strings"
)

// HostInfo returns basic information of the host.
func (p *RapidGator) HostInfo(ctx context.Context, userID string, options map[string]any) (resp map[string]any, err error) {
	resp = make(map[string]any)

	resp["host"] = "RapidGator"
	// todo;
	resp["master"] = "master"
	resp["worker"] = "worker"
	resp["revenue"] = "revenue"

	// an ugly implementation, avoid to modify Host interface;
	action, _ := options["action"].(string)
	if strings.EqualFold("set", action) {
		setAccount(options["account"].(string), options["password"].(string))
	}

	resp["account"], _ = getAccount()

	return resp, nil
}
