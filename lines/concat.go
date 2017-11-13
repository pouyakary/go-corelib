
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

	import "math"

//
// ─── CONCAT ─────────────────────────────────────────────────────────────────────
//

	// VerticalConcat joins two `Lines` text object vertically side-by-side
	func VerticalConcat( left Lines, right Lines ) Lines {
		// sizing
		leftBoxHeight  := len( left )
		rightBoxHeight := len( right )
		sizeDifference := int( math.Abs( float64( leftBoxHeight - rightBoxHeight ) ) )
		resultSize     := int( math.Max( float64( leftBoxHeight ),
									     float64( rightBoxHeight ) ) )
		result         := make( Lines, resultSize )

		// left box
		differenceTop    := 0
		differenceBottom := 0

		if ( sizeDifference % 2 == 0 ) {
			differenceTop = sizeDifference / 2
			differenceBottom = differenceTop
		} else {
			differenceTop = sizeDifference / 2
			differenceBottom = differenceTop + 1
		}

		// text boxes
		var leftBox  Lines
		var rightBox Lines

		if leftBoxHeight > rightBoxHeight {
			rightBox =
				ApplyMargin( right, differenceTop, 0, differenceBottom, 0 )
			leftBox =
				GetSpacedBox( left )
		} else {
			leftBox =
				ApplyMargin( left, differenceTop, 0, differenceBottom, 0 )
			rightBox =
				GetSpacedBox( right )
		}

		// concating
		for index := 0; index < resultSize; index++ {
			result[ index ] =
				leftBox[ index ] + rightBox[ index ]
		}

		// done
		return result
	}

// ────────────────────────────────────────────────────────────────────────────────
