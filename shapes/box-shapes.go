
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package shapes

    import "kary/core/lines"
    import "kary/core/shapes/charset"

//
// ─── PARENTHESIS ────────────────────────────────────────────────────────────────
//

    // CreateParesthesis create a __Paresthesis__ box out of a given `text` and
    // margin settings.
    func CreateParesthesis( text lines.Lines, margin Margin ) lines.Lines {
        parenthesisCharSet := ShapeCharset{
            LeftUpperCorner:       charset.LeftParenthesisUpperHook,
            LeftLowerCorner:       charset.LeftParenthesisLowerHook,
            RightUpperCorner:      charset.RightParenthesisUpperHook,
            RightLowerCorner:      charset.RightParenthesisLowerHook,
            LeftMiddleExtension:   charset.LeftParenthesisExtension,
            LeftExtension:         charset.LeftParenthesisExtension,
            RightMiddleExtension:  charset.RightParenthesisExtension,
            RightExtension:        charset.RightParenthesisExtension,
            TopExtension:          " ",
            BottomExtension:       " ",
        }

        return GenerateShapeBox( text, parenthesisCharSet, margin )
    }

//
// ─── SQUARE BRACKET ─────────────────────────────────────────────────────────────
//

    // CreateSquareBracket create a __square bracket__ box out of a given `text`
    // and margin settings.
    func CreateSquareBracket( text lines.Lines, margin Margin ) lines.Lines {
        squareBracketCharSet := ShapeCharset{
            LeftUpperCorner:       charset.LeftSquareBracketUpperCorner,
            LeftLowerCorner:       charset.LeftSquareBracketLowerCorner,
            RightUpperCorner:      charset.RightSquareBracketUpperCorner,
            RightLowerCorner:      charset.RightSquareBracketLowerCorner,
            LeftMiddleExtension:   charset.LeftSquareBracketExtension,
            LeftExtension:         charset.LeftSquareBracketExtension,
            RightMiddleExtension:  charset.RightSquareBracketExtension,
            RightExtension:        charset.RightSquareBracketExtension,
            TopExtension:          " ",
            BottomExtension:       " ",
        }

        return GenerateShapeBox( text, squareBracketCharSet, margin )
    }

//
// ─── CURLEY BRACKET ─────────────────────────────────────────────────────────────
//

    // CreateCurleyBracket create a __curley bracket__ box out of a given `text`
    // and  margin settings.
    func CreateCurleyBracket( text lines.Lines, margin Margin ) lines.Lines {
        curleyBracketCharSet := ShapeCharset{
            LeftUpperCorner:       charset.LeftCurleyBracketUpperHook,
            LeftLowerCorner:       charset.LeftCurleyBracketLowerHook,
            RightUpperCorner:      charset.RightCurleyBracketUpperHook,
            RightLowerCorner:      charset.RightCurleyBracketLowerHook,
            LeftMiddleExtension:   charset.LeftCurleyBracketMiddlePiece,
            LeftExtension:         charset.LeftCurleyBracketExtension,
            RightMiddleExtension:  charset.RightCurleyBracketMiddlePiece,
            RightExtension:        charset.RightCurleyBracketExtension,
            TopExtension:          " ",
            BottomExtension:       " ",
        }

        return GenerateShapeBox( text, curleyBracketCharSet, margin )
    }

//
// ─── LIGHT BOX ──────────────────────────────────────────────────────────────────
//

	// CreateLightBox creates a __light__ box out of a given `text`
    // and  margin settings.
    func CreateLightBox( text lines.Lines, margin Margin ) lines.Lines {
        curleyBracketCharSet := ShapeCharset{
            LeftUpperCorner:       charset.BoxDrawingsLightDownAndRight,
            LeftLowerCorner:       charset.BoxDrawingsLightUpAndRight,
            RightUpperCorner:      charset.BoxDrawingsLightDownAndLeft,
            RightLowerCorner:      charset.BoxDrawingsLightUpAndLeft,
            LeftMiddleExtension:   charset.BoxDrawingsLightVertical,
            LeftExtension:         charset.BoxDrawingsLightVertical,
            RightMiddleExtension:  charset.BoxDrawingsLightVertical,
            RightExtension:        charset.BoxDrawingsLightVertical,
            TopExtension:          charset.BoxDrawingsLightHorizontal,
            BottomExtension:       charset.BoxDrawingsLightHorizontal,
        }

        return GenerateShapeBox( text, curleyBracketCharSet, margin )
    }

//
// ─── HEAVY BOX ──────────────────────────────────────────────────────────────────
//

	// CreateHeavyBox creates a __heavy__ box out of a given `text`
    // and  margin settings.
    func CreateHeavyBox( text lines.Lines, margin Margin ) lines.Lines {
        curleyBracketCharSet := ShapeCharset{
            LeftUpperCorner:       charset.BoxDrawingsHeavyDownAndRight,
            LeftLowerCorner:       charset.BoxDrawingsHeavyUpAndRight,
            RightUpperCorner:      charset.BoxDrawingsHeavyDownAndLeft,
            RightLowerCorner:      charset.BoxDrawingsHeavyUpAndLeft,
            LeftMiddleExtension:   charset.BoxDrawingsHeavyVertical,
            LeftExtension:         charset.BoxDrawingsHeavyVertical,
            RightMiddleExtension:  charset.BoxDrawingsHeavyVertical,
            RightExtension:        charset.BoxDrawingsHeavyVertical,
            TopExtension:          charset.BoxDrawingsHeavyHorizontal,
            BottomExtension:       charset.BoxDrawingsHeavyHorizontal,
        }

        return GenerateShapeBox( text, curleyBracketCharSet, margin )
    }

// ────────────────────────────────────────────────────────────────────────────────
