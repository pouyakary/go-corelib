
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package shapes

    import "kary/core/line"

//
// ─── CREATE LINE ────────────────────────────────────────────────────────────────
//

    // CreateLine generates a line shape with a given `width`
    func CreateLine( width int ) string {
        return line.Repeat( "─", width )
    }

// ────────────────────────────────────────────────────────────────────────────────
