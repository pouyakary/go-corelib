
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

	package text

//
// ─── MAX LINE LENGTH ────────────────────────────────────────────────────────────
//

	// LongestLineLength `O(n)` &mdash; returns the length of longest line
	func LongestLineLength( text string ) int {
		max := 0
		currentLineSize := 0

		for _, char := range text {
			if char == '\n' {
				if ( currentLineSize > max ) {
					max = currentLineSize
				}
				currentLineSize = 0
			} else {
				currentLineSize++
			}
		}

		return max
	}

// ────────────────────────────────────────────────────────────────────────────────
