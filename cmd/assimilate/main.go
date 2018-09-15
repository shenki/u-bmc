// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/u-root/u-bmc/pkg/ast2400"
)

func main() {
	a := ast2400.Open()
	defer a.Close()

	a.Mem().MustWrite32(ast2400.SCU_BASE+0x0, 0x1688A8A8)
	a.Mem().MustWrite32(ast2400.SCU_BASE+0x84, 0xfffff000)
	a.Mem().MustWrite32(ast2400.SCU_BASE+0x0, 0x0)
}