# Package runtime

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


## Index

- [func Env(name string) string](#Env)
- [func Exit(status number)](#Exit)
- [func LookupEnv(name string) (string, bool)](#LookupEnv)
- [func SetEnv(name string, value string)](#SetEnv)
- [func Stack() StackTrace](#Stack)
- [func StackElement(File string, LineNumber number, LineOffset number, FunctionName string) StackElement](#StackElement)
- [func StackTrace(Elements []StackElement) StackTrace](#StackTrace)
- [func UnsetEnv(name string)](#UnsetEnv)

### Env

```
func Env(name string) string
```

Env returns the value for a environment variable. If the variable does not
exist an empty string is returned. To check if an environment variable is set
but empty you should use LookupEnv.

### Exit

```
func Exit(status number)
```

Exit will kill the current process and return the `status` code to the parent
process or system.

### LookupEnv

```
func LookupEnv(name string) (string, bool)
```

LookupEnv return the value and the existence of a environment variable.
LookupEnv is only required to determine cases where the environment variable
is set but empty. See Env.

### SetEnv

```
func SetEnv(name string, value string)
```

SetEnv is used to set or replace or existing environment variable. Setting an
environment variable to an empty value is not the same as unsetting it. See
UnsetEnv.

### Stack

```
func Stack() StackTrace
```

Stack returns the current stack in reverse order so that the deepest call
will be at the top.

### StackElement

```
func StackElement(File string, LineNumber number, LineOffset number, FunctionName string) StackElement
```

No documentation.

### StackTrace

```
func StackTrace(Elements []StackElement) StackTrace
```

No documentation.

### UnsetEnv

```
func UnsetEnv(name string)
```

UnsetEnv will remove an environment variable and return if the environment
variable previously existed. To check for the existence first, use LookupEnv.

