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
