
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

//
// ─── TYPINGS ────────────────────────────────────────────────────────────────────
//

    // ShapeCharset is a set of settings used to for templating the box format
    type ShapeCharset struct {
        LeftUpperCorner         string
        LeftLowerCorner         string
        RightUpperCorner        string
        RightLowerCorner        string
        LeftMiddleExtension     string
        LeftExtension           string
        RightMiddleExtension    string
        RightExtension          string
        TopExtension            string
        BottomExtension         string
    }

//
// ─── GENERATOR ──────────────────────────────────────────────────────────────────
//

    // GenerateShapeBox creates a shape box based on a text `input` _Lines_
    // with margins given as `verticalMargin` and `horizontalPadding` and a
    // shape character set
    func GenerateShapeBox( input lines.Lines, settings ShapeCharset ) lines.Lines {

        spacedBoxText  :=  input.CropToSpacedBox( )
        boxTextWidth   :=  spacedBoxText.LongestLineLength( )
        boxTextHeight  :=  len( spacedBoxText )
        result         :=  make( [ ]string, boxTextHeight + 2 )

        for lineNumber := 0; lineNumber < boxTextHeight + 2; lineNumber++ {
			result[ lineNumber ] =
				createLine(
					lineNumber, settings, spacedBoxText, boxTextWidth, boxTextHeight )
        }

        return result
    }

//
// ─── CREATE LINE ────────────────────────────────────────────────────────────────
//

	func createLine( lineNumber int, settings ShapeCharset, spacedBoxText lines.Lines, boxTextWidth, boxTextHeight int ) string {

		middleLineNumber :=
			( ( boxTextHeight + 2 ) / 2 ) - ( ( boxTextHeight % 2 ) + 1 )

		switch lineNumber {
		case 0:
			return settings.LeftUpperCorner + line.Repeat( settings.TopExtension, boxTextWidth ) + settings.RightUpperCorner

		case boxTextHeight + 1:
			return settings.LeftLowerCorner + line.Repeat( settings.BottomExtension, boxTextWidth ) + settings.RightLowerCorner

		case middleLineNumber:
			return  settings.LeftMiddleExtension + spacedBoxText[ lineNumber - 1 ] + settings.RightMiddleExtension

		default:
			return settings.LeftExtension + spacedBoxText[ lineNumber - 1 ] + settings.RightExtension
		}
	}

// ────────────────────────────────────────────────────────────────────────────────
