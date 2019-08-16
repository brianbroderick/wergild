package main

import "regexp"

// COLOR CODES and VT SEQUENCES
const modClear = "[0m" /* Resets Color (n) */
const modBold = "[1m"
const modFaint = "[2m"
const modUnderline = "[4m"
const modBlink = "[5m"
const modReverse = "[7m"

// Foreground Colors
const fgBlack = "[0;30m"   /* (l) */
const fgRed = "[0;31m"     /* (r) */
const fgGreen = "[0;32m"   /* (g) */
const fgYellow = "[0;33m"  /* (y) */
const fgBlue = "[0;34m"    /* (b) */
const fgMagenta = "[0;35m" /* (m) */
const fgCyan = "[0;36m"    /* (c) */
const fgWhite = "[0;37m"   /* (w) */

// Bold Foreground Colors
const fgBBlack = "[1;30m"   /* (L) */
const fgBRed = "[1;31m"     /* (R) */
const fgBGreen = "[1;32m"   /* (G) */
const fgBYellow = "[1;33m"  /* (Y) */
const fgBBlue = "[1;34m"    /* (B) */
const fgBMagenta = "[1;35m" /* (M) */
const fgBCyan = "[1;36m"    /* (C) */
const fgBWhite = "[1;37m"   /* (W) */

// Background Colors
const bgBlack = "[40m"
const bgRed = "[41m"
const bgGreen = "[42m"
const bgYellow = "[43m"
const bgBlue = "[44m"
const bgMagenta = "[45m"
const bgCyan = "[46m"
const bgWhite = "[47m"

const vtSaveCursor = "\x1b7"    // Save cursor and attributes
const vtRestCursor = "\x1b8"    // Restore cursor and attributes
const vtClearSet = "\x1b[r"     // Clear scrollable window size
const vtClearScreen = "\x1b[2J" // Clear screen
const vtClearLine = "\x1b[2K"   // Clear line
const vtTermReset = "\x1bc"     // Reset terminal completely

var colorMap = map[string]string{
	"{n": modClear,

	"{l": fgBlack,
	"{r": fgRed,
	"{g": fgGreen,
	"{y": fgYellow,
	"{b": fgBlue,
	"{m": fgMagenta,
	"{c": fgCyan,
	"{w": fgWhite,

	"{L": fgBBlack,
	"{R": fgBRed,
	"{G": fgBGreen,
	"{Y": fgBYellow,
	"{B": fgBBlue,
	"{M": fgBMagenta,
	"{C": fgBCyan,
	"{W": fgBWhite,
}

var colorExpression = regexp.MustCompile("\\{[nlrgybmcwLRGYBMCW]")

func colorCode(message string) string {
	return colorMap[message]
}

func colorize(message string) string {
	return colorExpression.ReplaceAllStringFunc(message, colorCode) + modClear
}
