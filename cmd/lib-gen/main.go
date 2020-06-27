package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/vm"
)

var f io.Writer

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func render(x interface{}, indent string) {
	ty := reflect.TypeOf(x)
	switch ty.Kind() {
	case reflect.Map:
		renderMap(x, indent)

	case reflect.Ptr:
		renderPtr(x, indent)

	case reflect.Slice:
		renderSlice(x, indent)

	default:
		fmt.Fprintf(f, "%#v", x)
	}
}

func localType(x interface{}) string {
	return strings.Replace(fmt.Sprintf("%T", x), "vm.", "", -1)
}

func renderMap(x interface{}, indent string) {
	v := reflect.ValueOf(x)
	if len(v.MapKeys()) == 0 {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "%s{\n", localType(x))
	for _, key := range v.MapKeys() {
		fmt.Fprintf(f, "%s\t\"%s\": ", indent, key)
		render(v.MapIndex(key).Interface(), indent+"\t")
		fmt.Fprintf(f, ",\n")
	}
	fmt.Fprintf(f, "%s}", indent)
}

func renderSlice(x interface{}, indent string) {
	v := reflect.ValueOf(x)
	if v.Len() == 0 {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "%s{\n", localType(x))
	for i := 0; i < v.Len(); i++ {
		fmt.Fprintf(f, "%s\t", indent)
		render(v.Index(i).Interface(), indent+"\t")
		fmt.Fprintf(f, ",\n")
	}
	fmt.Fprintf(f, "%s}", indent)
}

func renderPtr(x interface{}, indent string) {
	v := reflect.ValueOf(x).Elem()
	ty := reflect.TypeOf(x).Elem()
	fmt.Fprintf(f, "&%s{\n", localType(x)[1:])
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Fprintf(f, "%s\t%s: ", indent, ty.Field(i).Name)
		render(field.Interface(), indent+"\t")
		fmt.Fprintf(f, ",\n")
	}
	fmt.Fprintf(f, "%s}", indent)
}

func main() {
	pkg, err := compiler.CompilePackage("lib/math", false)
	check(err)

	f, err = os.Create("vm/lib.go")
	check(err)

	funcs := map[string]*vm.InternalDefinition{}
	for name, fn := range pkg.Funcs {
		funcDef := pkg.FuncDefs[name]

		// We don't need to serialize this.
		funcDef.Statements = nil

		funcs["math."+name] = &vm.InternalDefinition{
			CompiledFunc: fn,
			FuncDef:      funcDef,
		}
	}

	fmt.Fprintf(f, "package vm\n\n")
	fmt.Fprintf(f, "import \"github.com/elliotchance/ok/ast\"\n\n")
	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\tLib = ")
	render(funcs, "\t")
	fmt.Fprintf(f, "\n}\n")
	fmt.Fprintf(f, "\n")
}
