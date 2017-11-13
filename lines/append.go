
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

	package lines

//
// ─── APPEND ─────────────────────────────────────────────────────────────────────
//

	// Append adds a single line string to the start of all the lines within
	// a `Lines` text
	func Append( singleLineText string, text Lines ) Lines {
		for index, textLine := range text {
			text[ index ] = singleLineText + textLine
		}
		return text
	}

// ────────────────────────────────────────────────────────────────────────────────
