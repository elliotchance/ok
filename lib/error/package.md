The `errors` package contains standard interfaces for basic error handling.

### Example

```
import "error"

func divide(a, b number) number {
    if b == 0 {
        raise error.Error("cannot divide by zero")
    }

    return a / b
}
```
