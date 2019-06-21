package main

// V is a single value
type V struct {
	Symbol             string
	VerboseDescription string
	Index              int
}

// T is the struct holding the type information
type T struct {
	Name               string
	FileName           string
	PackageName        string
	TypeDescription    string
	UnknownDescription string
	Values             []V
}
