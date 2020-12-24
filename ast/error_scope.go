package ast

// On a error handler for an ErrorScope.
type On struct {
	// Type is the type name, like "MyError" or "error.Error".
	Type string

	// Statement may be nil.
	Statements []Node

	Pos string
}

// Position returns the position.
func (node *On) Position() string {
	return node.Pos
}

// Finally will always be called on success or failure.
type Finally struct {
	// Index is the unique counter for each finally block in this function. It
	// is used to activate and deactivate finally blocks by the VM at runtime.
	Index int

	// Statement may be nil.
	Statements []Node

	Pos string
}

// Position returns the position.
func (node *Finally) Position() string {
	return node.Pos
}

// ErrorScope represents the try/on error scope.
type ErrorScope struct {
	// Statements is what will be run in this scope. It is allowed to be nil.
	Statements []Node

	// On can have zero or more elements.
	On []*On

	Pos string

	// Finally is optional. An empty finally block and no finally block are
	// treated the same way.
	Finally *Finally
}

// Position returns the position.
func (node *ErrorScope) Position() string {
	return node.Pos
}
