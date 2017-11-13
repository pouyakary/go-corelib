
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
// ─── APPLY MARGIN ───────────────────────────────────────────────────────────────
//

	// ApplyMargin applies a white space margin to the given input `text`
	func ApplyMargin( text Lines, top, right, bottom, left int ) Lines {
		// input
		spacedBoxText       := CropToSpacedBox( text )

		// sizes
		spacedBoxTextHeight := len( spacedBoxText )
		spacedBoxTextWidth	:= len( spacedBoxText[ 0 ] )
		boxHeight			:= spacedBoxTextHeight + bottom + top
		boxWidth			:= spacedBoxTextWidth + left + right

		// result
		result 		        := make( [ ]string, boxHeight )
		resultLineNumber    := 0

		// template strings
		templateLeftSpace   := line.Repeat( " ", left )
		templateRightSpace  := line.Repeat( " ", right )
		templateEmptyLine   := line.Repeat( " ", boxWidth )

		// top part
		for index := 0; index < top; index++ {
			result[ resultLineNumber ] = templateEmptyLine
			resultLineNumber++
		}

		// middle part
		for index := 0; index < spacedBoxTextHeight; index++ {
			result[ resultLineNumber ] =
				templateLeftSpace + spacedBoxText[ index ] + templateRightSpace
			resultLineNumber++
		}

		// bottom part
		for index := 0; index < bottom; index++ {
			result[ resultLineNumber ] =
				templateEmptyLine
			resultLineNumber++
		}

		// done
		return result
	}

// ────────────────────────────────────────────────────────────────────────────────
