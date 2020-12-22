The `reflect` package contains runtime checking and manipulating of types and
values.

### Objects

```
import "reflect"

func MyObject(Foo, Bar number) MyObject {}

func main() {
    obj = MyObject(123, 456)

    reflect.Set(obj, "Bar", 789)
    reflect.Get(obj, "Bar")       // 789

    reflect.Properties(obj)       // ["Bar", "Foo"]
}
```

### Iterating Arrays or Maps

`Properties` works the same way with arrays or maps, but the key type will be a
`number` for array and `string` for a map.

```
import "reflect"

func iterateMap(x any) {
    for key in reflect.Properties(x) {
        print("{key} = {reflect.Get(x, key)}")
    }
}
```