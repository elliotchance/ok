package vm

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/elliotchance/ok/ast"
)

func Render(f io.Writer, x interface{}, indent string, vmPackage bool) {
	ty := reflect.TypeOf(x)

	// TODO(elliot): This is a special case for File and Reader.
	if ty == nil {
		fmt.Fprint(f, "nil")
		return
	}

	switch ty.Kind() {
	case reflect.Map:
		renderMap(f, x, indent, vmPackage)

	case reflect.Ptr:
		renderPtr(f, x, indent, vmPackage)

	case reflect.Slice:
		renderSlice(f, x, indent, vmPackage)

	default:
		fmt.Fprintf(f, "%#v", x)
	}
}

func localType(x interface{}, vmPackage bool) string {
	if vmPackage {
		return strings.Replace(fmt.Sprintf("%T", x), "vm.", "", -1)
	}

	return fmt.Sprintf("%T", x)
}

func renderMap(f io.Writer, x interface{}, indent string, vmPackage bool) {
	v := reflect.ValueOf(x)
	if len(v.MapKeys()) == 0 {
		fmt.Fprintf(f, "nil")
		return
	}

	// Maps keys need to be sorted.
	var mapKeys []reflect.Value
	for _, key := range v.MapKeys() {
		mapKeys = append(mapKeys, key)
	}

	sort.Slice(mapKeys, func(i, j int) bool {
		return mapKeys[i].String() < mapKeys[j].String()
	})

	fmt.Fprintf(f, "%s{\n", localType(x, vmPackage))
	for _, key := range mapKeys {
		fmt.Fprintf(f, "%s\t\"%s\": ", indent, key)
		Render(f, v.MapIndex(key).Interface(), indent+"\t", vmPackage)
		fmt.Fprintf(f, ",\n")
	}
	fmt.Fprintf(f, "%s}", indent)
}

func renderSlice(f io.Writer, x interface{}, indent string, vmPackage bool) {
	// Some types are handled as a single line. This is so the output generated
	// for vm/lib.go is far less verbose and more version control friendly.
	switch x.(type) {
	case Registers, []string:
		renderSliceOneLine(f, x, indent, vmPackage)
		return
	}

	v := reflect.ValueOf(x)
	if v.Len() == 0 {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "%s{\n", localType(x, vmPackage))
	for i := 0; i < v.Len(); i++ {
		value := v.Index(i)
		if !value.IsZero() {
			fmt.Fprintf(f, "%s\t", indent)
			Render(f, value.Interface(), indent+"\t", vmPackage)
			fmt.Fprintf(f, ",\n")
		}
	}
	fmt.Fprintf(f, "%s}", indent)
}

func renderSliceOneLine(f io.Writer, x interface{}, indent string, vmPackage bool) {
	v := reflect.ValueOf(x)
	if v.Len() == 0 {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "%s{", localType(x, vmPackage))
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			fmt.Fprintf(f, ",")
		}
		value := v.Index(i)
		Render(f, value.Interface(), indent+"\t", vmPackage)
	}
	fmt.Fprintf(f, "}")
}

func renderPtr(f io.Writer, x interface{}, indent string, vmPackage bool) {
	// Some types are handled as a single line. This is so the output generated
	// for vm/lib.go is far less verbose and more version control friendly.
	switch x.(type) {
	case Instruction, *ast.Literal, *ast.Argument:
		renderPtrWithoutKeys(f, x, indent, vmPackage)
		return
	}

	v := reflect.ValueOf(x).Elem()
	ty := reflect.TypeOf(x).Elem()

	if !v.IsValid() {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "&%s{\n", localType(x, vmPackage)[1:])
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsZero() {
			fmt.Fprintf(f, "%s\t%s: ", indent, ty.Field(i).Name)
			Render(f, field.Interface(), indent+"\t", vmPackage)
			fmt.Fprintf(f, ",\n")
		}
	}
	fmt.Fprintf(f, "%s}", indent)
}

func renderPtrWithoutKeys(f io.Writer, x interface{}, indent string, vmPackage bool) {
	v := reflect.ValueOf(x).Elem()

	if !v.IsValid() {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "&%s{", localType(x, vmPackage)[1:])
	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			fmt.Fprintf(f, ", ")
		}
		field := v.Field(i)
		Render(f, field.Interface(), indent+"\t", vmPackage)
	}
	fmt.Fprintf(f, "}")
}
