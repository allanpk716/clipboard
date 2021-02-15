// Copyright 2021 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed
// by a GNU GPL-3 license that can be found in the LICENSE file.
//
// Written by Changkun Ou <changkun.de>

package clipboard // import "golang.design/x/clipboard"

import (
	"sync"
)

// MIMEType represents the MIME type of clipboard data.
type MIMEType int

// All sorts of supported clipboard data
const (
	// MIMEText indicates plain text MIME format
	MIMEText MIMEType = iota
	// MIMEImage indicates image/png MIME format
	MIMEImage
)

// Due to the limitation on operating systems (such as darwin),
// concurrent read can even cause panic, use a global lock to
// guarantee one read at a time.
var lock = sync.Mutex{}

// Read reads and returns the clipboard data.
func Read(t MIMEType) []byte {
	lock.Lock()
	defer lock.Unlock()

	return read(t)
}

// Write writes the given buffer to the clipboard.
//
// If the MIME type indicates an image, then the given buf assumes
// the image data is PNG encoded.
func Write(t MIMEType, buf []byte) {
	lock.Lock()
	defer lock.Unlock()

	write(t, buf)
}