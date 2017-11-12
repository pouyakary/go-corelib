
//
// Copyright (c) 2017-present Kary Foundation, Inc. All rights reserved.
//   Author: Pouya Kary <k@karyfoundation.org>
//

//
// ─── NOTE ───────────────────────────────────────────────────────────────────────
//

	// this code is used from:
	// https://github.com/wayneashleyberry/terminal-dimensions
	// Copyright (c) 2016 Wayne Ashley Berry
	// published under MIT license

	// thanks a lot...

//
// ─── SETUP ──────────────────────────────────────────────────────────────────────
//

	package terminal

	import "os"
	import "os/exec"
	import "strconv"
	import "strings"

//
// ─── GET WIDTH ──────────────────────────────────────────────────────────────────
//

	// Width return the width of the terminal.
	func Width( ) ( int, error ) {
		output, err := size( )
		if err != nil {
			return 0, err
		}
		_, width, err := parse( output )

		return int( width ), err
	}

//
// ─── HEIGHT ─────────────────────────────────────────────────────────────────────
//

	// Height returns the height of the terminal.
	func Height( ) ( int, error ) {
		output, err := size( )
		if err != nil {
			return 0, err
		}
		height, _, err := parse( output )
		return int( height ), err
	}

//
// ─── SIZE ───────────────────────────────────────────────────────────────────────
//

	func size( ) ( string, error ) {
		cmd := exec.Command( "stty", "size" )
		cmd.Stdin = os.Stdin
		out, err := cmd.Output( )
		return string( out ), err
	}

//
// ─── PARSE ──────────────────────────────────────────────────────────────────────
//

	func parse( input string ) ( uint, uint, error ) {
		parts := strings.Split( input, " " )
		x, err := strconv.Atoi( parts[ 0 ] )
		if err != nil {
			return 0, 0, err
		}
		y, err := strconv.Atoi(
			strings.Replace( parts[ 1 ], "\n", "", 1 ) )
		if err != nil {
			return 0, 0, err
		}
		return uint( x ), uint( y ), nil
	}

// ────────────────────────────────────────────────────────────────────────────────
