package mql

// Statement represents a single command
type Statement interface {
	String() string
	KeyTok() Token
}
