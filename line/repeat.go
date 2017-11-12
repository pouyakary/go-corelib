
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package line

//
// ─── REPEAT TEXT ────────────────────────────────────────────────────────────────
//

    // Repeat contacts a string `text` to itself by the `times` specified.
    func Repeat( text string, times int ) string {
        result := make( [ ]rune, len( text ) * times )
        globalIndex := 0

        for index := 0; index < times; index++ {
            for _, char := range text {
                result[ globalIndex ] = char
                globalIndex++
            }
        }

        return string( result )
    }

// ────────────────────────────────────────────────────────────────────────────────
