// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains extra hooks for testing the go command.

// +build testgo

package work

import "os"

func init() {
	if v := os.Getenv("TESTGO_VERSION"); v != "" {
		runtimeVersion = v
	}
}
