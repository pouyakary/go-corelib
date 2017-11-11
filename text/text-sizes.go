
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

	package text

//
// ─── TEXT LONGEST LINE ──────────────────────────────────────────────────────────
//

	// LongestLineLength returns the length of the longest line
	func ( text Text ) LongestLineLength( ) int {
		max := 0
		for _, line := range text {
			lineLength := len( line )
			if lineLength > max {
				max = lineLength
			}
		}
		return max
	}

// ────────────────────────────────────────────────────────────────────────────────
