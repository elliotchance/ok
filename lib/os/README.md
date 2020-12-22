# Package os

The `os` package contains operating and file system functions.

### Reading a File

```
import "os"

func main() {
    f = os.Open("foo.txt")
    for {
        line = f.ReadLine()
        if line == "" {
            break
        }

        print(line)
    }

    f.Close()
}
```

### Writing to a File

```
import "os"

func main() {
    f = Open("file.txt")
    f.WriteString("hello\n")
    f.WriteString("world\n")
    f.Close()
}
```


## Index

- [func CreateDirectory(path string)](#CreateDirectory)
- [func File(fd data) File](#File)
- [func FileInfo(Name string, Size number, Mode string, ModifiedTime time.Time, IsDir bool) FileInfo](#FileInfo)
- [func Info(path string) FileInfo](#Info)
- [func Open(path string) File](#Open)
- [func Remove(path string)](#Remove)
- [func Rename(old string, new string)](#Rename)
- [func TempPath() string](#TempPath)

### CreateDirectory

```
func CreateDirectory(path string)
```

CreateDirectory will create a directory at `path`, creating any necessary
directories along the way.

### File

```
func File(fd data) File
```

File represents a file handle.

### FileInfo

```
func FileInfo(Name string, Size number, Mode string, ModifiedTime time.Time, IsDir bool) FileInfo
```

No documentation.

### Info

```
func Info(path string) FileInfo
```

Info returns the FileInfo for a path, or raises an error if the path does not
exist.

### Open

```
func Open(path string) File
```

Open opens a file for reading or writing. If the file does not exist it will
be created. If the file does exist the cursor will begin at the start of the
file. You may use `f.Seek(0, 2)` immediately after opening the file if you
wish for writes to be appended at the end.

### Remove

```
func Remove(path string)
```

Remove will delete (unlink) the file located at `path`.

### Rename

```
func Rename(old string, new string)
```

Rename will move the file or directory located at `old` to `new`.

### TempPath

```
func TempPath() string
```

TempPath returns an absolute path that is safe to create a new file.

