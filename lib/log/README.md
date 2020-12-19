# log

- [LevelDebug string](#constants)
- [LevelError string](#constants)
- [LevelFatal string](#constants)
- [LevelInfo string](#constants)
- [LevelWarn string](#constants)

- [func Log(level string, message string)](#Log)
- [func Logger(Log func(string, string)) Logger](#Logger)

## Constants

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

## Log

```
func Log(level string, message string)
```

Log is a very simple implementation for Logger. You may call this directly
for dynamic log levels, but it is recommened to put this into a Logger first.

## Logger

```
func Logger(Log func(string, string)) Logger
```

Logger provides the interface for a standard logger. You can use the Log
function provided to create a simple logger:

```
logger = log.Logger(log.Log)
```

