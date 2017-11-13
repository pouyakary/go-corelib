
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

	import "kary/core/line"

//
// ─── MAKE SPACE BOX ─────────────────────────────────────────────────────────────
//

	// GetSpacedBox makes a white space space box of the given text
	func GetSpacedBox( text Lines ) Lines {
		longestLine := LongestLineLength( text )
		for index, singleLine := range text {
			text[ index ] =
				line.PadRight( singleLine, longestLine )
		}
		return text
	}

// ────────────────────────────────────────────────────────────────────────────────
