
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

	package text
	import "strings"

//
// ─── REMOVE FROM END ────────────────────────────────────────────────────────────
//

	// RemoveSuffix removes a `suffix` from the end of string if that `suffix`
	// is present
	func RemoveSuffix( target, suffix string ) string {
		if strings.HasSuffix( target, suffix ) {
			lengthDifference := len( target ) - len( suffix )
			return target[ : lengthDifference ]
		}
		return target
	}

//
// ─── REMOVE FROM START ──────────────────────────────────────────────────────────
//

	// RemovePrefix removes a `prefix` from the end of string if that `prefix`
	// is present
	func RemovePrefix( target, prefix string ) string {
		if strings.HasPrefix( target, prefix ) {
			return target[ len( prefix ) : ]
		}
		return target
	}

// ────────────────────────────────────────────────────────────────────────────────
