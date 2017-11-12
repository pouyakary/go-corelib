
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package lines

//
// ─── TEXT LONGEST LINE ──────────────────────────────────────────────────────────
//

    // LongestLineLength returns the length of the longest line
    func ( text Lines ) LongestLineLength( ) int {
        max := 0
        for _, singleLine := range text {
            lineLength := len( singleLine )
            if lineLength > max {
                max = lineLength
            }
        }
        return max
    }

// ────────────────────────────────────────────────────────────────────────────────
