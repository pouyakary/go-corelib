
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package shapes

    import "kary/core/line"
	import "kary/core/lines"
	import "kary/core/shapes/charset"

//
// ─── HORIZONTAL LINE ────────────────────────────────────────────────────────────
//

    // CreateHorizontalLightLine generates a light horizontal line shape
	// with a given `width`
    func CreateHorizontalLightLine( width int ) string {
        return line.Repeat( charset.BoxDrawingsLightHorizontal, width )
    }

	// CreateHorizontalHeavyLine generates a heavy horizontal line shape
	// with a given `width`
	func CreateHorizontalHeavyLine( width int ) string {
		return line.Repeat( charset.BoxDrawingsHeavyHorizontal, width )
	}

//
// ─── VERTICAL LINE ──────────────────────────────────────────────────────────────
//

	// CreateVerticalLightLine generates a light vertical line shape with a
	// given `height`
	func CreateVerticalLightLine( height int ) lines.Lines {
		result := make( lines.Lines, height )
		for index := 0; index < height; index++ {
			result[ index ] = charset.BoxDrawingsLightVertical
		}
		return result
	}

	// CreateVerticalHeavyLine generates a heavy vertical line shape with a
	// given `height`
	func CreateVerticalHeavyLine( height int ) lines.Lines {
		result := make( lines.Lines, height )
		for index := 0; index < height; index++ {
			result[ index ] = charset.BoxDrawingsHeavyVertical
		}
		return result
	}

// ────────────────────────────────────────────────────────────────────────────────
