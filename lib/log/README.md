# Package log

The `log` package contains the standard log levels and interfaces for logging.

### Using the Standard Logger

```
import "log"

func main() {
    // Create a simple logger for stdout.
    l = log.Logger(log.Log)

    l.Info("ready")
}
```

### Implementing a Custom Logger

```
import "log"

func main() {
    // A buffered logger that can be inspected.
    messages = []string []
    logger = log.Logger(func (level, message string) {
        ^messages += ["{level} {message}"]
    })

    l.Info("ready")

    // messages = ["INFO ready"]
}
```

When creating a custom logger, be aware that `LogLevelFatal` should invoke
`runtime.Exit(1)` as clients can expect that `Fatal()` will stop the execution
of the program. One such case that you might not exist, however, is under
specific controlled conditions like a unit test.


## Index

- [LevelDebug string](#constants)
- [LevelError string](#constants)
- [LevelFatal string](#constants)
- [LevelInfo string](#constants)
- [LevelWarn string](#constants)

- [func Log(level string, message string)](#Log)
- [func Logger(Log func(string, string)) Logger](#Logger)

### Constants

```
LevelDebug = DEBUG
```

```
LevelError = ERROR
```

```
LevelFatal = FATAL
```

```
LevelInfo = INFO
```

```
LevelWarn = WARN
```

### Log

```
func Log(level string, message string)
```

Log is a very simple implementation for Logger. You may call this directly
for dynamic log levels, but it is recommened to put this into a Logger first.

### Logger

```
func Logger(Log func(string, string)) Logger
```

Logger provides the interface for a standard logger. You can use the Log
function provided to create a simple logger:

```
logger = log.Logger(log.Log)
```

