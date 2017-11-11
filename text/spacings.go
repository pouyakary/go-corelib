
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

	package text

//
// ─── PAD LEFT ───────────────────────────────────────────────────────────────────
//

	// NormalizeLineBreaks `O(n)` &mdash; converts all line break conventions to
	// single `\n`. These line breaks are: `\n`, `\r`, `\v`, `\f`, `\u0085`,
	// `\u2028`, `\u2029` and `\r\n`
	func NormalizeLineBreaks( text string ) string {

		// Possible line breaks are: \n \r \v \f \u0085 \u2028 \u2029
		// but also the \r\n sequence is considered a line break

    	result := make( [ ]rune, len( text ) )
		iterableText := [ ]rune( text )
		previousChar := rune( 0 )
		index := 0

		for _, char := range iterableText {
			switch char {
			case '\n':
				previousChar = char
				if previousChar == '\r' {
					break
				} else {
					result[ index ] = '\n'
					index++
				}
				break

			case '\r':
			case '\v':
			case '\f':
			case '\u0085':
			case '\u2028':
			case '\u2029':
				previousChar = char
				result[ index ] = '\n'
				index++
				break

			default:
				previousChar = char
				result[ index ] = char
				index++
			}
		}

		return string( result )
	}

//
// ─── PAD LEFT ───────────────────────────────────────────────────────────────────
//

	// PadLeft appends space character to the _left_ side of the string `text`
	// till it's __total__ length becomes `size`. If the specified `size` be less
	// than the `text`'s length, it will return the original `text` value
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
