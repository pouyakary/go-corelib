
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

	package text

	import "strings"
	import "kary/core/str"

//
// ─── MULTI LINE TRIM ────────────────────────────────────────────────────────────
//

	// CropToSpacedBox removes text all around the string and returns a boxed string
	func ( text Text ) CropToSpacedBox( ) Text {
		longestLine := text.LongestLineLength( )
		resultingLines := make( [ ]string, len( text ) )

		for index, line := range text {
			trimmedLine := strings.TrimSpace( line )
			spacedLine := str.PadRight( trimmedLine, longestLine )
			resultingLines[ index ] = spacedLine
		}

		return resultingLines
	}

// ────────────────────────────────────────────────────────────────────────────────
