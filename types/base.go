package types

var (
	Any    = TypeFromString("any")
	Bool   = TypeFromString("bool")
	Char   = TypeFromString("char")
	Data   = TypeFromString("data")
	Number = TypeFromString("number")
	String = TypeFromString("string")

	AnyArray    = NewArray(Any)
	BoolArray   = NewArray(Bool)
	NumberArray = NewArray(Number)
	StringArray = NewArray(String)

	AnyMap    = NewMap(Any)
	NumberMap = NewMap(Number)
	StringMap = NewMap(String)

	ErrorInterface = NewInterface("error.Error", map[string]*Type{
		"Error": String,
	})
)
