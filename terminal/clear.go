
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package terminal

//
// ─── CLEAR TERMINAL SCREEN ──────────────────────────────────────────────────────
//

    // Clear cleans the Terminal screen
    func Clear( ) {
        print( "\033[H\033[2J" )
    }

// ────────────────────────────────────────────────────────────────────────────────
