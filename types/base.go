package types

var (
	Any    = TypeFromString("any")
	Bool   = TypeFromString("bool")
	Char   = TypeFromString("char")
	Data   = TypeFromString("data")
	Number = TypeFromString("number")
	String = TypeFromString("string")

	AnyArray    = TypeFromString("[]any")
	BoolArray   = TypeFromString("[]bool")
	NumberArray = TypeFromString("[]number")
	StringArray = TypeFromString("[]string")

	AnyMap    = TypeFromString("{}any")
	NumberMap = TypeFromString("{}number")
	StringMap = TypeFromString("{}string")

	ErrorInterface = NewInterface("error.Error", map[string]*Type{
		"Error": String,
	})
)
