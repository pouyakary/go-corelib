
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package lines

    import "strings"
    import "kary/core/line"

//
// ─── MULTI LINE TRIM ────────────────────────────────────────────────────────────
//

    // CropToSpacedBox removes text all around the string and returns a boxed string
    func CropToSpacedBox( text Lines ) Lines {
        longestLine    := LongestLineLength( text )
        resultingLines := make( [ ]string, len( text ) )

        for index, singleLine := range text {
            trimmedLine := strings.TrimSpace( singleLine )
            spacedLine  := line.PadRight( trimmedLine, longestLine )
            resultingLines[ index ] = spacedLine
        }

        return resultingLines
    }

// ────────────────────────────────────────────────────────────────────────────────
