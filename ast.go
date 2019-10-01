package main

// Statement represents a single command
type Statement interface {
	Node
	// stmt is unexported to ensure implementations of Statement
	// can only originate in this package.
	stmt()
	// RequiredPrivileges() (ExecutionPrivileges, error)
}

// Node represents a node in the abstract syntax tree.
type Node interface {
	// node is unexported to ensure implementations of Node
	// can only originate in this package.
	node()
	String() string
}
