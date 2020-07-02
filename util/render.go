package util

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

func Render(f io.Writer, x interface{}, indent string, vmPackage bool) {
	ty := reflect.TypeOf(x)
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

	fmt.Fprintf(f, "%s{\n", localType(x, vmPackage))
	for _, key := range v.MapKeys() {
		fmt.Fprintf(f, "%s\t\"%s\": ", indent, key)
		Render(f, v.MapIndex(key).Interface(), indent+"\t", vmPackage)
		fmt.Fprintf(f, ",\n")
	}
	fmt.Fprintf(f, "%s}", indent)
}

func renderSlice(f io.Writer, x interface{}, indent string, vmPackage bool) {
	v := reflect.ValueOf(x)
	if v.Len() == 0 {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "%s{\n", localType(x, vmPackage))
	for i := 0; i < v.Len(); i++ {
		fmt.Fprintf(f, "%s\t", indent)
		Render(f, v.Index(i).Interface(), indent+"\t", vmPackage)
		fmt.Fprintf(f, ",\n")
	}
	fmt.Fprintf(f, "%s}", indent)
}

func renderPtr(f io.Writer, x interface{}, indent string, vmPackage bool) {
	v := reflect.ValueOf(x).Elem()
	ty := reflect.TypeOf(x).Elem()

	if !v.IsValid() {
		fmt.Fprintf(f, "nil")
		return
	}

	fmt.Fprintf(f, "&%s{\n", localType(x, vmPackage)[1:])
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Fprintf(f, "%s\t%s: ", indent, ty.Field(i).Name)
		Render(f, field.Interface(), indent+"\t", vmPackage)
		fmt.Fprintf(f, ",\n")
	}
	fmt.Fprintf(f, "%s}", indent)
}
