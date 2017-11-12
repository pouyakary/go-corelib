
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package line

//
// ─── PAD LEFT ───────────────────────────────────────────────────────────────────
//

    // PadLeft appends space character to the _left_ side of the string `text`
    // till it's __total__ length becomes `size`
    func PadLeft( text string, size int ) string {
        textLength := len( text )
        spacesNeeded := size - textLength
        if ( spacesNeeded > 0 ) {
            return Repeat( " ", spacesNeeded ) + text
        }
        return text
    }

//
// ─── PAD RIGHT ──────────────────────────────────────────────────────────────────
//

    // PadRight appends space character to the _right_ side of the string `text`
    // till it's __total__ length becomes `size`. If the specified `size` be less
    // than the `text`'s length, it will return the original `text` value
    func PadRight( text string, size int ) string {
        textLength := len( text )
        spacesNeeded := size - textLength
        if ( spacesNeeded > 0 ) {
            return text + Repeat( " ", spacesNeeded )
        }
        return text
    }

// ────────────────────────────────────────────────────────────────────────────────
