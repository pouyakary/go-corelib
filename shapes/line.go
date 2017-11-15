
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
        return createHorizontalLine( charset.BoxDrawingsLightHorizontal, width )
    }

    // CreateHorizontalHeavyLine generates a heavy horizontal line shape
    // with a given `width`
    func CreateHorizontalHeavyLine( width int ) string {
        return createHorizontalLine( charset.BoxDrawingsHeavyHorizontal, width )
    }

    func createHorizontalLine( char string, width int ) string {
        return line.Repeat( char, width )
    }

//
// ─── VERTICAL LINE ──────────────────────────────────────────────────────────────
//

    // CreateVerticalLightLine generates a light vertical line shape with a
    // given `height`
    func CreateVerticalLightLine( height int ) lines.Lines {
        return createVerticalLine( charset.BoxDrawingsLightVertical, height )
    }

    // CreateVerticalHeavyLine generates a heavy vertical line shape with a
    // given `height`
    func CreateVerticalHeavyLine( height int ) lines.Lines {
        return createVerticalLine( charset.BoxDrawingsHeavyVertical, height )
    }

    func createVerticalLine( char string, height int ) lines.Lines {
        result := make( lines.Lines, height )
        for index := 0; index < height; index++ {
            result[ index ] = char
        }
        return result
    }

//
// ─── HORIZONTAL 2D LINE ─────────────────────────────────────────────────────────
//

    // CreateHorizontalUpLeftToBottomRightLine is a line object from top left to
    // bottom right specified by it's `width` and `height`
    func CreateHorizontalUpLeftToBottomRightLine( width, height int ) lines.Lines {
        // in special case of height == 1
        if height == 1 {
            return lines.Lines{ createHorizontalLine(
                charset.BoxDrawingsLightHorizontal, width ) }
        }

        // in the normal case which is done within three phases: top line,
        // middle vertical line, bottom line.

        leftSize, rightSize := computeLineLeftRightSpaces( width )
        result := make( lines.Lines, height )

        // top line
        topLineLeftSide  := line.Repeat( charset.BoxDrawingsLightHorizontal, leftSize )
        topLineRightSide := line.Repeat( " ", rightSize )
        topLine          := ( topLineLeftSide + charset.BoxDrawingsLightDownAndLeft +
                              topLineRightSide )

        result[ 0 ] = topLine


        // middle parts
        middleLineLeftSide  := line.Repeat( " ", leftSize )
        middleLineRightSize := line.Repeat( " ", rightSize )
        middleLine          := ( middleLineLeftSide + charset.BoxDrawingsLightVertical +                           middleLineRightSize )

        for index := 1; index < height - 1; index++ {
            result[ index ] = middleLine
        }

        // ending bottom line
        bottomLineLeftSide  := line.Repeat( " ", leftSize )
        bottomLineRightSide := line.Repeat( charset.BoxDrawingsLightHorizontal,
                                            rightSize )
        bottomLine          := ( bottomLineLeftSide + charset.BoxDrawingsLightUpAndRight +
                                 bottomLineRightSide )

        result[ height - 1 ] = bottomLine


        // done
        return result
    }

//
// ──────────── COMPUTE LEFT AND RIGHT SIDE SPACES FOR THE HORIZONTAL 2D LINE ─────
//

    func computeLineLeftRightSpaces( width int ) ( left, right int ) {
        if width % 2 == 1 {
            left = width / 2
            right = width / 2
        } else {
            left = width / 2
            right = left - 1
        }
        return
    }

// ────────────────────────────────────────────────────────────────────────────────
