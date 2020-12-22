The `runtime` package contains functionality for the current process and
runtime environment.

### Environment Variables

```
import "runtime"

func main() {
    print("$HOME = {runtime.Env("HOME")}")
}
```

### Stack Traces

```
import "runtime"

func main() {
    print(runtime.Stack().String())
}
```
