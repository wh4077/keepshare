// Copyright 2023 The KeepShare Authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/KeepShareOrg/keepshare/cmd"

	// _ "github.com/KeepShareOrg/keepshare/hosts/pikpak"     // register PikPak host
	_ "github.com/KeepShareOrg/keepshare/hosts/rapidgator" // register RapidGator host
)

func main() {
	cmd.Execute()
}
