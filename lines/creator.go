
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

    package lines

//
// ─── CREATE TEXT OBJECT ─────────────────────────────────────────────────────────
//

    // Create `O(n)` &mdash; Creates a Kary Framework `Text` from plain text `string`
    func Create( input string ) Lines {

        result        :=  [ ]string { }
        iterableText  :=  [ ]rune( input )
        previousChar  :=  rune( 0 )
        currentLine   :=  ""
        index         :=  0

        for _, char := range iterableText {
            switch char {
            case '\n':
                if previousChar != '\r' {
                    result = append( result, currentLine )
                    currentLine = ""
                    index++
                }
                previousChar = char
                break

            case '\r', '\v', '\f', '\u0085', '\u2028', '\u2029':
                previousChar = char
                result = append( result, currentLine )
                currentLine = ""
                index++
                break

            default:
                previousChar = char
                currentLine += string( char )
                index++
            }
        }

        // adding currentline
        result = append( result, currentLine )

        // done
        return result
    }

// ────────────────────────────────────────────────────────────────────────────────
