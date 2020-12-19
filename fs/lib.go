package fs

import "os"
import "github.com/blang/vfs"
import "github.com/blang/vfs/memfs"
import "github.com/blang/vfs/mountfs"

func init() {
	Filesystem = mountfs.Create(vfs.OS())
	var fs vfs.Filesystem
	var f vfs.File
	fs = memfs.Create()
	f, _ = fs.OpenFile("error.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Error is a basic type to carry an error message.\n" +
		"func Error(Error string) Error {\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/error")
	fs = memfs.Create()
	f, _ = fs.OpenFile("decode-object.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"reflect\"\n" +
		"\n" +
		"func DecodeObject(json string, dest any) {\n" +
		"    map = Decode(json)\n" +
		"    if map is {}any {\n" +
		"        for value, key in map {\n" +
		"            // We must permit the first letter being upper cased since we can\n" +
		"            // only modify public properties. However, snake_case and camelCase\n" +
		"            // will not work.\n" +
		"            if key[0]\n" +
		"\n" +
		"            reflect.Set(dest, key, value)\n" +
		"        }\n" +
		"\n" +
		"        return\n" +
		"    }\n" +
		"\n" +
		"    //raise error.Error(\"json is not an object\")\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("decode.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"error\"\n" +
		"\n" +
		"func Decode(json string) any {\n" +
		"    tokens = tokenize(json)\n" +
		"    pos = 0\n" +
		"    \n" +
		"    // There can only be one root element, so we should not see any more tokens.\n" +
		"    next = consumeNext()\n" +
		"    if pos < len(tokens) {\n" +
		"        raise error.Error(\"invalid json\")\n" +
		"    }\n" +
		"\n" +
		"    return next\n" +
		"\n" +
		"    func consumeNext() any {\n" +
		"        token = ^consume()\n" +
		"\n" +
		"        if token == \"[\" {\n" +
		"            return ^consumeArray()\n" +
		"        }\n" +
		"\n" +
		"        if token == \"\\{\" {\n" +
		"            return ^consumeMap()\n" +
		"        }\n" +
		"\n" +
		"        return ^decodeToken(token)\n" +
		"    }\n" +
		"\n" +
		"    func consumeArray() []any {\n" +
		"        arr = []any []\n" +
		"        for {\n" +
		"            token = ^consume()\n" +
		"            if token == \"]\" {\n" +
		"                break\n" +
		"            }\n" +
		"\n" +
		"            // TODO(elliot): This doesn't ensure that the JSON is correctly\n" +
		"            //  formatted.\n" +
		"            if token == \",\" {\n" +
		"                continue\n" +
		"            }\n" +
		"\n" +
		"            arr += []any [^decodeToken(token)]\n" +
		"        }\n" +
		"\n" +
		"        return arr\n" +
		"    }\n" +
		"\n" +
		"    // TODO(elliot): Cannot set return type as \"{}any\" without syntax\n" +
		"    func consumeMap() any {\n" +
		"        map = {}any {}\n" +
		"        for {\n" +
		"            keyToken = ^consume()\n" +
		"            if keyToken == \"}\" {\n" +
		"                break\n" +
		"            }\n" +
		"\n" +
		"            // TODO(elliot): This doesn't ensure that the JSON is correctly\n" +
		"            //  formatted.\n" +
		"            if keyToken == \",\" {\n" +
		"                continue\n" +
		"            }\n" +
		"\n" +
		"            // TODO(elliot): This doesn't ensure that the JSON is correctly\n" +
		"            //  formatted.\n" +
		"            ^consume() // ignore \":\"\n" +
		"\n" +
		"            valueToken = ^consume()\n" +
		"            map[strings.Substr(keyToken, 1, len(keyToken))] = ^decodeToken(valueToken)\n" +
		"        }\n" +
		"\n" +
		"        return map\n" +
		"    }\n" +
		"\n" +
		"    func consume() string {\n" +
		"        token = ^peek()\n" +
		"        ++^pos\n" +
		"\n" +
		"        return token\n" +
		"    }\n" +
		"\n" +
		"    func peek() string {\n" +
		"        if ^pos >= len(^tokens) {\n" +
		"            raise error.Error(\"invalid json\")\n" +
		"        }\n" +
		"\n" +
		"        return ^tokens[^pos]\n" +
		"    }\n" +
		"\n" +
		"    func decodeToken(token string) any {\n" +
		"        if token == \"true\" {\n" +
		"            return true\n" +
		"        }\n" +
		"\n" +
		"        if token == \"false\" {\n" +
		"            return false\n" +
		"        }\n" +
		"\n" +
		"        if strings.HasPrefix(token, \"S\") {\n" +
		"            return strings.Substr(token, 1, len(token))\n" +
		"        }\n" +
		"\n" +
		"        if strings.HasPrefix(token, \"N\") {\n" +
		"            return number strings.Substr(token, 1, len(token))\n" +
		"        }\n" +
		"\n" +
		"        return Null()\n" +
		"    }\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("encode.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"strings\"\n" +
		"import \"reflect\"\n" +
		"\n" +
		"// Encode transforms a value into a JSON string.\n" +
		"func Encode(value any) string {\n" +
		"    if value is bool {\n" +
		"        return string value\n" +
		"    }\n" +
		"    \n" +
		"    if value is number {\n" +
		"        return string value\n" +
		"    }\n" +
		"\n" +
		"    if value is char {\n" +
		"        return str(string value)\n" +
		"    }\n" +
		"\n" +
		"    if value is data {\n" +
		"        return str(string value)\n" +
		"    }\n" +
		"\n" +
		"    if value is string {\n" +
		"        return str(value)\n" +
		"    }\n" +
		"\n" +
		"    if value is Null {\n" +
		"        return \"null\"\n" +
		"    }\n" +
		"\n" +
		"    if reflect.Kind(value) == \"array\" {\n" +
		"        s = \"[\"\n" +
		"        for i = 0; i < reflect.Len(value); ++i {\n" +
		"            if s != \"[\" {\n" +
		"                s += \",\"\n" +
		"            }\n" +
		"\n" +
		"            s += Encode(reflect.Get(value, i))\n" +
		"        }\n" +
		"\n" +
		"        return s + \"]\"\n" +
		"    }\n" +
		"\n" +
		"    // This code works for both objects and maps. Keys are sorted by Properties.\n" +
		"    s = \"\\{\"\n" +
		"    for key in reflect.Properties(value) {\n" +
		"        if s != \"\\{\" {\n" +
		"            s += \",\"\n" +
		"        }\n" +
		"\n" +
		"        s += \"\\\"{key}\\\":{Encode(reflect.Get(value, key))}\"\n" +
		"    }\n" +
		"\n" +
		"    return s + \"}\"\n" +
		"}\n" +
		"\n" +
		"func str(s string) string {\n" +
		"    escaped = strings.ReplaceAll(s, \"\\\"\", \"\\\\\\\"\")\n" +
		"\n" +
		"    return \"\\\"{escaped}\\\"\"\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("null.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Null provides a value and interface for JSON \"null\". It can be used as a\n" +
		"// literal:\n" +
		"//\n" +
		"// ```\n" +
		"// assert(json.Encode(json.Null()) == \"null\")\n" +
		"// ```\n" +
		"//\n" +
		"// Or, it can be used as an interface for existing types:\n" +
		"//\n" +
		"// ```\n" +
		"// func Person(name string) {\n" +
		"//     func IsJSONNull() bool {\n" +
		"//         return ^name == \"\"\n" +
		"//     }\n" +
		"// }\n" +
		"// ```\n" +
		"func Null() Null {\n" +
		"    func IsJSONNull() bool {\n" +
		"        return true\n" +
		"    }\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("tokenize.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"error\"\n" +
		"\n" +
		"func tokenize(json string) []string {\n" +
		"    tokens = []string []\n" +
		"    word = \"\"\n" +
		"\n" +
		"    for i = 0; i < len(json); ++i {\n" +
		"        // Ignore whitespace\n" +
		"        if json[i] == ' ' or json[i] == '\\n' or json[i] == '\\t' or json[i] == '\\r' {\n" +
		"            continue\n" +
		"        }\n" +
		"\n" +
		"        // Numbers\n" +
		"        if json[i] >= '0' and json[i] <= '9' {\n" +
		"            // N is the token prefix that describes the type for the parser.\n" +
		"            word = \"N\"\n" +
		"\n" +
		"            for i < len(json) {\n" +
		"                if json[i] != '.' and (json[i] < '0' or json[i] > '9') {\n" +
		"                    break\n" +
		"                }\n" +
		"\n" +
		"                word += string json[i]\n" +
		"                ++i\n" +
		"            }\n" +
		"\n" +
		"            tokens += [word]\n" +
		"            word = \"\"\n" +
		"            --i\n" +
		"            continue\n" +
		"        }\n" +
		"\n" +
		"        // Strings\n" +
		"        if json[i] == '\"' {\n" +
		"            ++i // skip \"\n" +
		"\n" +
		"            // S is the token prefix that describes the type for the parser.\n" +
		"            word = \"S\"\n" +
		"\n" +
		"            for {\n" +
		"                if i >= len(json) {\n" +
		"                    raise error.Error(\"unterminated string\")\n" +
		"                }\n" +
		"\n" +
		"                if json[i] == '\"' {\n" +
		"                    break\n" +
		"                }\n" +
		"\n" +
		"                word += string json[i]\n" +
		"                ++i\n" +
		"            }\n" +
		"\n" +
		"            tokens += [word]\n" +
		"            word = \"\"\n" +
		"            continue\n" +
		"        }\n" +
		"\n" +
		"        // Operators\n" +
		"        // TODO(elliot): https://github.com/elliotchance/ok/issues/108\n" +
		"        if json[i] == '{' or json[i] == '}' or json[i] == '[' or json[i] == ']' or json[i] == ':' or json[i] == ',' {\n" +
		"            if word != \"\" {\n" +
		"                tokens += [word]\n" +
		"                word = \"\"\n" +
		"            }\n" +
		"\n" +
		"            tokens += [string json[i]]\n" +
		"            continue\n" +
		"        }\n" +
		"        \n" +
		"        word += string json[i]\n" +
		"    }\n" +
		"\n" +
		"    if word != \"\" {\n" +
		"        tokens += [word]\n" +
		"    }\n" +
		"\n" +
		"    return tokens\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/json")
	fs = memfs.Create()
	f, _ = fs.OpenFile("log.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"runtime\"\n" +
		"import \"strings\"\n" +
		"import \"time\"\n" +
		"\n" +
		"// These log levels are described in similarly named methods of the Logger\n" +
		"// interface.\n" +
		"LevelDebug = \"DEBUG\"\n" +
		"LevelInfo = \"INFO\"\n" +
		"LevelWarn = \"WARN\"\n" +
		"LevelError = \"ERROR\"\n" +
		"LevelFatal = \"FATAL\"\n" +
		"\n" +
		"// Logger provides the interface for a standard logger. You can use the Log\n" +
		"// function provided to create a simple logger:\n" +
		"//\n" +
		"// ```\n" +
		"// logger = log.Logger(log.Log)\n" +
		"// ```\n" +
		"func Logger(Log func(string, string)) Logger {\n" +
		"    // Debug messages in this level contain extensive contextual information.\n" +
		"    // They are mostly used for problem diagnosis. Information on this Level are\n" +
		"    // for Developers and not for the Users.\n" +
		"    func Debug(message string) {\n" +
		"        ^Log(LevelDebug, message)\n" +
		"    }\n" +
		"\n" +
		"    // Info messages contain some contextual information to help trace execution\n" +
		"    // (at a coarse-grained level) in a production environment.\n" +
		"    func Info(message string) {\n" +
		"        ^Log(LevelInfo, message)\n" +
		"    }\n" +
		"\n" +
		"    // Warn messages indicate a potential problem in the system. The System is\n" +
		"    // able to handle the problem by themself or to proccede with this problem\n" +
		"    // anyway.\n" +
		"    func Warn(message string) {\n" +
		"        ^Log(LevelWarn, message)\n" +
		"    }\n" +
		"\n" +
		"    // Error messages indicate a serious problem in the system. The problem is\n" +
		"    // usually non-recoverable and requires manual intervention.\n" +
		"    func Error(message string) {\n" +
		"        ^Log(LevelError, message)\n" +
		"    }\n" +
		"\n" +
		"    // Fatal indicates that the entire system is unhealthy and must exit now.\n" +
		"    // Fatal is different from other levels in that after outputting its log\n" +
		"    // message it will exit the process with a non-success status code.\n" +
		"    func Fatal(message string) {\n" +
		"        ^Log(LevelFatal, message)\n" +
		"    }\n" +
		"}\n" +
		"\n" +
		"// Log is a very simple implementation for Logger. You may call this directly\n" +
		"// for dynamic log levels, but it is recommended to put this into a Logger\n" +
		"// first.\n" +
		"func Log(level, message string) {\n" +
		"    // https://github.com/elliotchance/ok/issues/99\n" +
		"    level = strings.PadLeft(level, \" \", 5)\n" +
		"    print(\"{time.Now().String()} {level} {message}\")\n" +
		"\n" +
		"    if level == LevelFatal {\n" +
		"        runtime.Exit(1)\n" +
		"    }\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/log")
	fs = memfs.Create()
	f, _ = fs.OpenFile("abs.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Abs returns the absolute (positive number).\n" +
		"func Abs(x number) number {\n" +
		"    if x < 0 {\n" +
		"        return -x\n" +
		"    }\n" +
		"\n" +
		"    return x\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("constants.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("E   = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113\n" +
		"Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796\n" +
		"Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // https://oeis.org/A001622\n" +
		"\n" +
		"Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193\n" +
		"SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774\n" +
		"SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161\n" +
		"SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339\n" +
		"\n" +
		"Ln2  = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162\n" +
		"Ln10 = 2.30258509299404568401799145468436420760110148862877297603332790  // https://oeis.org/A002392\n" +
		""))
	f, _ = fs.OpenFile("log.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// LogE is the natural logarithm (base e).\n" +
		"func LogE(x number) number {\n" +
		"    return __log(x)\n" +
		"}\n" +
		"\n" +
		"// Log10 is a logarithm with base 10.\n" +
		"func Log10(x number) number {\n" +
		"    return __log(x) / __log(10)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("powers.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Exp returns e to the power of x.\n" +
		"func Exp(x number) number {\n" +
		"    return __pow(E, x)\n" +
		"}\n" +
		"\n" +
		"// Pow returns base to power.\n" +
		"func Pow(base, power number) number {\n" +
		"    return __pow(base, power)\n" +
		"}\n" +
		"\n" +
		"// Sqrt returns the square root of x.\n" +
		"func Sqrt(x number) number {\n" +
		"    return __pow(x, 0.5)\n" +
		"}\n" +
		"\n" +
		"// Cbrt returns the cube root of x.\n" +
		"func Cbrt(x number) number {\n" +
		"    return __pow(x, 1/3)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("rand.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Rand returns a random number between 0 and 1.\n" +
		"func Rand() number {\n" +
		"    return __rand()\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("rounding.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Ceil will round x up to the nearest integer.\n" +
		"func Ceil(x number) number {\n" +
		"    frac = x % 1\n" +
		"    if frac == 0 {\n" +
		"        return x\n" +
		"    }\n" +
		"\n" +
		"    if x < 0 {\n" +
		"        return x - frac\n" +
		"    }\n" +
		"\n" +
		"    return x + (1 - frac)\n" +
		"}\n" +
		"\n" +
		"// Floor will round x down to the nearest integer.\n" +
		"func Floor(x number) number {\n" +
		"    frac = x % 1\n" +
		"    if frac == 0 {\n" +
		"        return x\n" +
		"    }\n" +
		"\n" +
		"    if x < 0 {\n" +
		"        return x - frac + 1\n" +
		"    }\n" +
		"\n" +
		"    return x - frac\n" +
		"}\n" +
		"\n" +
		"// Round will return a new number rounded to prec number of digits after the\n" +
		"// decimal point. Prec must be at least 0.\n" +
		"func Round(x, prec number) number {\n" +
		"    p = __pow(10, prec)\n" +
		"    y = x * p\n" +
		"\n" +
		"    diff = y % 1\n" +
		"    if diff >= 0.5 {\n" +
		"        return (y + (1 - diff)) / p\n" +
		"    }\n" +
		"\n" +
		"    return (y - diff) / p\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/math")
	fs = memfs.Create()
	f, _ = fs.OpenFile("file.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// File represents a file handle.\n" +
		"func File(fd data) File {\n" +
		"    // WriteString writes a string to the current position of the file.\n" +
		"    func WriteString(s string) {\n" +
		"        __write(^fd, data s)\n" +
		"    }\n" +
		"\n" +
		"    // WriteData writes data to the current position of the file.\n" +
		"    func WriteData(d data) {\n" +
		"        __write(^fd, d)\n" +
		"    }\n" +
		"\n" +
		"    // Close closes the file handle. After Close is invoked you can no longer\n" +
		"    // interact with the file.\n" +
		"    func Close() {\n" +
		"        __close(^fd)\n" +
		"    }\n" +
		"\n" +
		"    // Seek sets the offset for the next read or write on file to offset,\n" +
		"    // interpreted according to whence:\n" +
		"    //\n" +
		"    // - 0 means relative to the origin of the file\n" +
		"    // - 1 means relative to the current offset\n" +
		"    // - 2 means relative to the end.\n" +
		"    //\n" +
		"    func Seek(offset, whence number) number {\n" +
		"        __seek(^fd, offset, whence)\n" +
		"    }\n" +
		"\n" +
		"    // ReadBytes will read up to `bytes` in length. The actual number of bytes\n" +
		"    // read will be returned. If the number of bytes read is less than `bytes`\n" +
		"    // the end of the file has been reached. Trying to call ReadBytes will\n" +
		"    // always result in 0 bytes being read.\n" +
		"    //\n" +
		"    // It's important to understand this reads data (raw bytes) which may not be\n" +
		"    // compatible with multibyte text encodings like UTF-8. For reading strings\n" +
		"    // use ReadString instead.\n" +
		"    func ReadData(bytes number) data {\n" +
		"        return __read_data(^fd, bytes)\n" +
		"    }\n" +
		"\n" +
		"    // ReadString follows the same rules as ReadBytes, but is designed for\n" +
		"    // reading UTF-8 strings where each character may be more than one byte. The\n" +
		"    // number returned will be the number of characters, not bytes.\n" +
		"    func ReadString(chars number) string {\n" +
		"        return __read_string(^fd, chars)\n" +
		"    }\n" +
		"\n" +
		"    // ReadLine will read a string until a new line or carriage return is hit.\n" +
		"    func ReadLine() string {\n" +
		"        line = \"\"\n" +
		"\n" +
		"        for {\n" +
		"            chars = __read_string(^fd, 1)\n" +
		"            if len(chars) == 0 {\n" +
		"                break\n" +
		"            }\n" +
		"\n" +
		"            line += chars\n" +
		"\n" +
		"            if chars == \"\\n\" {\n" +
		"                break\n" +
		"            }\n" +
		"        }\n" +
		"\n" +
		"        return line\n" +
		"    }\n" +
		"}\n" +
		"\n" +
		"// Open opens a file for reading or writing. If the file does not exist it will\n" +
		"// be created. If the file does exist the cursor will begin at the start of the\n" +
		"// file. You may use `f.Seek(0, 2)` immediately after opening the file if you\n" +
		"// wish for writes to be appended at the end.\n" +
		"func Open(path string) File {\n" +
		"    return File(__open(path))\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("filesystem.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Remove will delete (unlink) the file located at `path`.\n" +
		"func Remove(path string) {\n" +
		"    __remove(path)\n" +
		"}\n" +
		"\n" +
		"// Rename will move the file or directory located at `old` to `new`.\n" +
		"func Rename(old, new string) {\n" +
		"    __rename(old, new)\n" +
		"}\n" +
		"\n" +
		"// CreateDirectory will create a directory at `path`, creating any necessary\n" +
		"// directories along the way.\n" +
		"func CreateDirectory(path string) {\n" +
		"    __mkdir(path)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("info.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"time\"\n" +
		"\n" +
		"func FileInfo(\n" +
		"    Name string,            // base name of the file\n" +
		"    Size number,            // length in bytes for regular files; system-dependent for others\n" +
		"    Mode string,            // file mode bits\n" +
		"    ModifiedTime time.Time, // modification time\n" +
		"    IsDir bool              // easier than decoding from Mode\n" +
		") FileInfo {}\n" +
		"\n" +
		"// Info returns the FileInfo for a path, or raises an error if the path does not\n" +
		"// exist.\n" +
		"func Info(path string) FileInfo {\n" +
		"    name, size, mode, modTimeYear, modTimeMonth, modTimeDay, modTimeHour, modTimeMinute, modTimeSecond, isDir = __info(path)\n" +
		"    modTime = time.Time(modTimeYear, modTimeMonth, modTimeDay, modTimeHour, modTimeMinute, modTimeSecond)\n" +
		"    \n" +
		"    return FileInfo(name, size, mode, modTime, isDir)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("temp.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// TempPath returns an absolute path that is safe to create a new file.\n" +
		"func TempPath() string {\n" +
		"    // TODO(elliot): \"/tmp\" is not always the right choice. Also, maybe we\n" +
		"    //  should incorporate the date?\n" +
		"    return \"/tmp/ok.\" + string (__rand() * 1000000000)\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/os")
	fs = memfs.Create()
	f, _ = fs.OpenFile("call.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Call can be used to call a function variable without knowing the type.\n" +
		"//\n" +
		"// `fn` must be a callable function and args must have the correct number of\n" +
		"// arguments for the function. The number of return values will depend on the\n" +
		"// function being called.\n" +
		"//\n" +
		"// Example:\n" +
		"//\n" +
		"// ```\n" +
		"// fn = func(a, b number) number {\n" +
		"//    return a + b\n" +
		"// }\n" +
		"//\n" +
		"// print(reflect.Call(fn, []any [1.2, 3.4]))\n" +
		"// ```\n" +
		"func Call(fn any, args []any) []any {\n" +
		"    return __call(fn, args)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("get.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Get performs one of the following (depending on the input type):\n" +
		"//\n" +
		"// - Arrays: `prop` must be an number that represents the index in the array. If\n" +
		"// it is not a number or is out of bound then an error is raised.\n" +
		"//\n" +
		"// - Maps: `prop` must be a string that represents the key in the map. If prop\n" +
		"// is not a string of the key does not exist in the map an error will be raised.\n" +
		"//\n" +
		"// - Objects: `prop` must be a string representing the property name. You may\n" +
		"// only access public properties (starting with a capital letter). Trying to\n" +
		"// access a property that does not exist or is not public will result in an\n" +
		"// error.\n" +
		"//\n" +
		"// Any other type will always result in an error.\n" +
		"func Get(obj, prop any) any {\n" +
		"    return __get(obj, prop)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("interface.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Interface returns the canonical interface as a string.\n" +
		"//\n" +
		"// Examples:\n" +
		"//\n" +
		"// ```\n" +
		"// { }\n" +
		"// { Name string }\n" +
		"// { Greet(string) bool; Name string }\n" +
		"// ```\n" +
		"func Interface(value any) string {\n" +
		"    return __interface(value)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("kind.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"strings\"\n" +
		"\n" +
		"// Kind returns the runtime type of a value. One of; \"bool\", \"char\", \"data\",\n" +
		"// \"number\", \"string\", \"array\", \"map\" or \"func\".\n" +
		"func Kind(value any) string {\n" +
		"    type = Type(value)\n" +
		"\n" +
		"    switch {\n" +
		"        case strings.HasPrefix(type, \"[]\") {\n" +
		"            return \"array\"\n" +
		"        }\n" +
		"\n" +
		"        case strings.HasPrefix(type, \"\\{}\") {\n" +
		"            return \"map\"\n" +
		"        }\n" +
		"\n" +
		"        case strings.HasPrefix(type, \"func(\") {\n" +
		"            return \"func\"\n" +
		"        }\n" +
		"    }\n" +
		"\n" +
		"    return type\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("len.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Len returns the number of elements in an array or map. If the value is not an\n" +
		"// array or map then an error is raised.\n" +
		"func Len(value any) number {\n" +
		"    return __len(value)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("props.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Properties returns the public properties of an object, or the keys in a map.\n" +
		"// If the input is not an object or map, an error is raised. The values returned\n" +
		"// will be sorted in either case.\n" +
		"func Properties(obj any) []string {\n" +
		"    return __props(obj)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("set.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Set performs one of the following (depending on the input type):\n" +
		"//\n" +
		"// - Arrays: `prop` must be an number that represents the index in the array. If\n" +
		"// it is not a number or is out of bound then an error is raised.\n" +
		"//\n" +
		"// - Maps: `prop` must be a string that represents the key in the map. If prop\n" +
		"// is not a string an error will be raised. However, if the key does not exist\n" +
		"// it will be created. If the key already exists its value will be replaced.\n" +
		"//\n" +
		"// - Objects: `prop` must be a string representing the property name. You may\n" +
		"// only access public properties (starting with a capital letter). Trying to\n" +
		"// access a property that does not exist or is not public will result in an\n" +
		"// error.\n" +
		"//\n" +
		"// Any other type will always result in an error.\n" +
		"func Set(obj, prop, value any) any {\n" +
		"    return __set(obj, prop, value)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("type.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Type returns the runtime type of a value, some examples would be:\n" +
		"//\n" +
		"// ```\n" +
		"// reflect.Type(3)                      == \"number\"\n" +
		"// reflect.Type([\"foo\", \"bar\"])         == \"[]string\"\n" +
		"// reflect.Type(func hello(ok bool) {}) == \"func(bool)\"\n" +
		"// ```\n" +
		"func Type(value any) string {\n" +
		"    return __type(value)\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/reflect")
	fs = memfs.Create()
	f, _ = fs.OpenFile("env.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Env returns the value for a environment variable. If the variable does not\n" +
		"// exist an empty string is returned. To check if an environment variable is set\n" +
		"// but empty you should use LookupEnv.\n" +
		"func Env(name string) string {\n" +
		"    value, _ = LookupEnv(name)\n" +
		"\n" +
		"    return value\n" +
		"}\n" +
		"\n" +
		"// LookupEnv return the value and the existence of a environment variable.\n" +
		"// LookupEnv is only required to determine cases where the environment variable\n" +
		"// is set but empty. See Env.\n" +
		"func LookupEnv(name string) (string, bool) {\n" +
		"    return __env_get(name)\n" +
		"}\n" +
		"\n" +
		"// SetEnv is used to set or replace or existing environment variable. Setting an\n" +
		"// environment variable to an empty value is not the same as unsetting it. See\n" +
		"// UnsetEnv.\n" +
		"func SetEnv(name, value string) {\n" +
		"    return __env_set(name, value)\n" +
		"}\n" +
		"\n" +
		"// UnsetEnv will remove an environment variable and return if the environment\n" +
		"// variable previously existed. To check for the existence first, use LookupEnv.\n" +
		"func UnsetEnv(name string) {\n" +
		"    __env_unset(name)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("exit.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Exit will kill the current process and return the `status` code to the parent\n" +
		"// process or system.\n" +
		"func Exit(status number) {\n" +
		"    __exit(status)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("stack.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"strings\"\n" +
		"\n" +
		"func StackElement(\n" +
		"    File string,\n" +
		"    LineNumber, LineOffset number,\n" +
		"    FunctionName string\n" +
		") StackElement {}\n" +
		"\n" +
		"// Stack returns the current stack in reverse order so that the deepest call\n" +
		"// will be at the top.\n" +
		"func Stack() StackTrace {\n" +
		"    elements = strings.Split(__stack(), \"\\n\")\n" +
		"    stack = []StackElement []\n" +
		"    for element in elements {\n" +
		"        parts = strings.Split(element, \"|\")\n" +
		"        locationParts = strings.Split(parts[0], \":\")\n" +
		"        stack += [StackElement(\n" +
		"            locationParts[0],\n" +
		"            number locationParts[1],\n" +
		"            number locationParts[2],\n" +
		"            parts[1])]\n" +
		"    }\n" +
		"\n" +
		"    return StackTrace(stack)\n" +
		"}\n" +
		"\n" +
		"func StackTrace(Elements []StackElement) StackTrace {\n" +
		"    func String() string {\n" +
		"        s = \"\"\n" +
		"        for i = 1; element in ^Elements; ++i {\n" +
		"            s += \"{i} {element.FunctionName} {element.File}:{element.LineNumber}\\n\"\n" +
		"        }\n" +
		"\n" +
		"        // Remove the final newline.\n" +
		"        return strings.TrimRight(s, \"\\n\")\n" +
		"    }\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/runtime")
	fs = memfs.Create()
	f, _ = fs.OpenFile("case.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// ToLower returns a lower case version of s.\n" +
		"//\n" +
		"// TODO(elliot): This only works for ASCII characters.\n" +
		"func ToLower(s string) string {\n" +
		"    result = \"\"\n" +
		"    for c in s {\n" +
		"        n = number c\n" +
		"        if n >= number 'A' and n <= number 'Z' {\n" +
		"            result += string char (n + 32)\n" +
		"        } else {\n" +
		"            result += string c\n" +
		"        }\n" +
		"    }\n" +
		"\n" +
		"    return result\n" +
		"}\n" +
		"\n" +
		"// ToUpper returns a upper case version of s.\n" +
		"//\n" +
		"// TODO(elliot): This only works for ASCII characters.\n" +
		"func ToUpper(s string) string {\n" +
		"    result = \"\"\n" +
		"    for c in s {\n" +
		"        n = number c\n" +
		"        if n >= number 'a' and n <= number 'z' {\n" +
		"            result += string char (n - 32)\n" +
		"        } else {\n" +
		"            result += string c\n" +
		"        }\n" +
		"    }\n" +
		"\n" +
		"    return result\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("contains.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Contains checks whether substr exists in s.\n" +
		"func Contains(s, substr string) bool {\n" +
		"    return Index(s, substr) != -1\n" +
		"}\n" +
		"\n" +
		"// HasPrefix will return true if s starts with prefix.\n" +
		"func HasPrefix(s, prefix string) bool {\n" +
		"    if len(s) < len(prefix) {\n" +
		"        return false\n" +
		"    }\n" +
		"\n" +
		"    for i = 0; i < len(prefix); ++i {\n" +
		"        if s[i] != prefix[i] {\n" +
		"            return false\n" +
		"        }\n" +
		"    }\n" +
		"\n" +
		"    return true\n" +
		"}\n" +
		"\n" +
		"// HasSuffix will return true if s ends with suffix.\n" +
		"func HasSuffix(s, suffix string) bool {\n" +
		"    if len(s) < len(suffix) {\n" +
		"        return false\n" +
		"    }\n" +
		"\n" +
		"    j = len(s) - 1\n" +
		"    for i = len(suffix) - 1; i >= 0; --i {\n" +
		"        if s[j] != suffix[i] {\n" +
		"            return false\n" +
		"        }\n" +
		"\n" +
		"        --j\n" +
		"    }\n" +
		"\n" +
		"    return true\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("index.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Index returns the index of the first occurance of substr; or -1 if not found.\n" +
		"func Index(s, substr string) number {\n" +
		"    return IndexAfter(s, substr, -1)\n" +
		"}\n" +
		"\n" +
		"// IndexAfter returns the index of the first occurance of substr at or after\n" +
		"// offset; or -1 if not found. If the offset is zero or negative then this\n" +
		"// function will perform as Index.\n" +
		"//\n" +
		"// Example: Iterating through all matches\n" +
		"//\n" +
		"//     substr = \"foo\"\n" +
		"//     for i = Index(s, substr); i != -1; i = IndexAfter(s, substr, i) {\n" +
		"//         print(i)\n" +
		"//     }\n" +
		"//\n" +
		"func IndexAfter(s, substr string, offset number) number {\n" +
		"    offset = max(offset, -1)\n" +
		"\n" +
		"    for i = offset + 1; i <= len(s) - len(substr); ++i {\n" +
		"        found = true\n" +
		"\n" +
		"        for j = 0; j < len(substr); ++j {\n" +
		"            if s[i + j] != substr[j] {\n" +
		"                found = false\n" +
		"                break\n" +
		"            }\n" +
		"        }\n" +
		"\n" +
		"        if found {\n" +
		"            return i\n" +
		"        }\n" +
		"    }\n" +
		"\n" +
		"    return -1\n" +
		"}\n" +
		"\n" +
		"// TODO(elliot): This can be moved into the math library.\n" +
		"func min(a, b number) number {\n" +
		"    if a < b {\n" +
		"        return a\n" +
		"    }\n" +
		"\n" +
		"    return b\n" +
		"}\n" +
		"\n" +
		"// TODO(elliot): This can be moved into the math library.\n" +
		"func max(a, b number) number {\n" +
		"    if a > b {\n" +
		"        return a\n" +
		"    }\n" +
		"\n" +
		"    return b\n" +
		"}\n" +
		"\n" +
		"// LastIndex returns the start of the last occurence of substr in s.\n" +
		"func LastIndex(s, substr string) number {\n" +
		"    index = Index(Reverse(s), Reverse(substr))\n" +
		"    if index == -1 {\n" +
		"        return -1\n" +
		"    }\n" +
		"\n" +
		"    return len(s) - index + len(substr)\n" +
		"}\n" +
		"\n" +
		"// LastIndexBefore returns the start of the last occurence of substr that starts\n" +
		"// before offset.\n" +
		"//\n" +
		"// Example: Iterating through all matches\n" +
		"//\n" +
		"//     substr = \"foo\"\n" +
		"//     for i = LastIndex(s, substr); i != -1; i = LastIndexBefore(s, substr, i) {\n" +
		"//         print(i)\n" +
		"//     }\n" +
		"//\n" +
		"func LastIndexBefore(s, substr string, offset number) number {\n" +
		"    // Since we are reversing the string we need to make sure an offset larger\n" +
		"    // than the string does not eat into the prefix.\n" +
		"    offset = len(s) - min(offset, len(s)) + 1\n" +
		"\n" +
		"    index = IndexAfter(Reverse(s), Reverse(substr), offset)\n" +
		"    if index == -1 {\n" +
		"        return -1\n" +
		"    }\n" +
		"\n" +
		"    return len(s) - index + len(substr)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("join.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Join returns a string containing all elements joined with glue. The glue will\n" +
		"// only appear between elements, even if those elements are empty strings\n" +
		"// themselves. Glue is allowed to be empty.\n" +
		"func Join(strings []string, glue string) string {\n" +
		"    result = \"\"\n" +
		"    for s, i in strings {\n" +
		"        if i > 0 {\n" +
		"            result += glue\n" +
		"        }\n" +
		"\n" +
		"        result += s\n" +
		"    }\n" +
		"\n" +
		"    return result\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("pad.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// PadLeft will return a string with a length of at least `toLen` in length.\n" +
		"// `s` will not be truncated if it's longer than `toLen`.\n" +
		"//\n" +
		"// If `pad` is more than one character, the whole string is repeated, except if\n" +
		"// `pad` does not entirely fit, then `pad` will be truncated on the last\n" +
		"// occurrence.\n" +
		"//\n" +
		"// If `pad` is empty, the original string will always be returned.\n" +
		"func PadLeft(s, pad string, toLen number) string {\n" +
		"    if len(s) >= toLen or pad == \"\" {\n" +
		"        return s\n" +
		"    }\n" +
		"\n" +
		"    return createPad(pad, toLen - len(s)) + s\n" +
		"}\n" +
		"\n" +
		"// PadRight follows all the same rules as PadLeft, but will place padding (if\n" +
		"// any) on the right side of the string.\n" +
		"func PadRight(s, pad string, toLen number) string {\n" +
		"    if len(s) >= toLen or pad == \"\" {\n" +
		"        return s\n" +
		"    }\n" +
		"\n" +
		"    return s + createPad(pad, toLen - len(s))\n" +
		"}\n" +
		"\n" +
		"func createPad(pad string, toLen number) string {\n" +
		"    return Substr(Repeat(pad, toLen / len(pad)), 0, toLen)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("repeat.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Repeat returns str repeated times. Any values for times that is 0 or below\n" +
		"// will result in an empty string.\n" +
		"func Repeat(str string, times number) string {\n" +
		"    result = \"\"\n" +
		"    for i = 0; i < times; ++i {\n" +
		"        result += str\n" +
		"    }\n" +
		"\n" +
		"    return result\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("replace.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// ReplaceAll return a string with each of find substituted with replace.\n" +
		"//\n" +
		"// If `find` is empty, then `replace` will be inserted between every character.\n" +
		"// It will not be insert before the first character or after the last.\n" +
		"func ReplaceAll(s, find, replace string) string {\n" +
		"    return Join(Split(s, find), replace)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("reverse.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Reverse creates a string with all characters in the opposite order.\n" +
		"func Reverse(s string) string {\n" +
		"    result = \"\"\n" +
		"    for i = len(s) - 1; i >= 0; --i {\n" +
		"        result += string(s[i])\n" +
		"    }\n" +
		"\n" +
		"    return result\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("split.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Split will always returns one or more elements. If the glue is empty the\n" +
		"// string will be split into characters.\n" +
		"//\n" +
		"// TODO(elliot): This is a horribly inefficient algorithm. This was very early\n" +
		"//  on in the language when there we're barely any features, please clean this\n" +
		"//  up if you see it.\n" +
		"func Split(s, delimiter string) []string {\n" +
		"    elements = []string []\n" +
		"    \n" +
		"    if delimiter == \"\" {\n" +
		"        // TODO(elliot): In the future this should be possible by casting s to\n" +
		"        // []char.\n" +
		"        for i = 0; i < len(s); ++i {\n" +
		"            elements += [string s[i]]\n" +
		"        }\n" +
		"    } else {\n" +
		"        element = \"\"\n" +
		"        for i = 0; i < len(s); ++i {\n" +
		"            if IndexAfter(s, delimiter, i - 1) - i == 0 {\n" +
		"                elements += [element]\n" +
		"                element = \"\"\n" +
		"                i += len(delimiter) - 1\n" +
		"            } else {\n" +
		"                // TODO(elliot): We shouldn't need to have this extra cast here.\n" +
		"                element += string s[i]\n" +
		"            }\n" +
		"        }\n" +
		"\n" +
		"        elements += [element]\n" +
		"    }\n" +
		"\n" +
		"    return elements\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("substr.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Substr returns a portion of the string. The `fromIndex` and `toIndex` must be\n" +
		"// within the bounds of the string.\n" +
		"func Substr(s string, fromIndex, toIndex number) string {\n" +
		"    r = \"\"\n" +
		"    for i = fromIndex; i < toIndex; ++i {\n" +
		"        r += string(s[i])\n" +
		"    }\n" +
		"\n" +
		"    return r\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("trim.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// TrimLeft returns string with any of the cutset characters removed from the\n" +
		"// left (start).\n" +
		"func TrimLeft(s, cutset string) string {\n" +
		"    for offset = 0; offset < len(s); ++offset {\n" +
		"        if Index(cutset, string s[offset]) == -1 {\n" +
		"            return substrFrom(s, offset)\n" +
		"        }\n" +
		"    }\n" +
		"\n" +
		"    return s\n" +
		"}\n" +
		"\n" +
		"// TrimRight returns string with any of the cutset characters removed from the\n" +
		"// right (end).\n" +
		"func TrimRight(s, cutset string) string {\n" +
		"    return Reverse(TrimLeft(Reverse(s), cutset))\n" +
		"}\n" +
		"\n" +
		"// Trim returns string with any of the cutset characters removed from the left\n" +
		"// (start) and right (end).\n" +
		"func Trim(s, cutset string) string {\n" +
		"    return TrimRight(TrimLeft(s, cutset), cutset)\n" +
		"}\n" +
		"\n" +
		"// TrimPrefix will trim the prefix if it exists at the start of the string, or\n" +
		"// return the original string.\n" +
		"//\n" +
		"// If prefix appears more then once at the start of the string only the first\n" +
		"// prefix will be removed.\n" +
		"//\n" +
		"// If prefix is equal to s then an empty result will be returned.\n" +
		"func TrimPrefix(s, prefix string) string {\n" +
		"    if HasPrefix(s, prefix) {\n" +
		"        return substrFrom(s, len(prefix))\n" +
		"    }\n" +
		"\n" +
		"    return s\n" +
		"}\n" +
		"\n" +
		"// TrimSuffix will trim the suffix if it exists at the end of the string, or\n" +
		"// return the original string.\n" +
		"//\n" +
		"// If suffix appears more then once at the end of the string only the last\n" +
		"// suffix will be removed.\n" +
		"//\n" +
		"// If suffix is equal to s then an empty result will be returned.\n" +
		"func TrimSuffix(s, suffix string) string {\n" +
		"    return Reverse(TrimPrefix(Reverse(s), Reverse(suffix)))\n" +
		"}\n" +
		"\n" +
		"// TODO(elliot): This can be removed when we support array ranges.\n" +
		"func substrFrom(s string, index number) string {\n" +
		"    result = \"\"\n" +
		"    for index < len(s) {\n" +
		"        result += string s[index]\n" +
		"        ++index\n" +
		"    }\n" +
		"\n" +
		"    return result\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/strings")
	fs = memfs.Create()
	f, _ = fs.OpenFile("add.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Add returns a new time after applying a duration. You may use a negative\n" +
		"// duration to subtract.\n" +
		"func Add(t Time, duration Duration) Time {\n" +
		"    return FromUnix(Unix(t) + duration.Seconds())\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("compare.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Sub returns the duration between two time. The result will be negative if `b`\n" +
		"// is before `a`, positive if `b` is after `a` and `0` if the two times are\n" +
		"// equal.\n" +
		"func Sub(a, b Time) Duration {\n" +
		"    return Duration(Unix(a) - Unix(b))\n" +
		"}\n" +
		"\n" +
		"// Equal returns true only if both instances represent the exact same time\n" +
		"// (including fractional seconds).\n" +
		"func Equal(a, b Time) bool {\n" +
		"    duration = Sub(a, b)\n" +
		"\n" +
		"    return duration.Seconds() == 0\n" +
		"}\n" +
		"\n" +
		"// Before returns true if `a` is before `b`.\n" +
		"func Before(a, b Time) bool {\n" +
		"    duration = Sub(a, b)\n" +
		"    \n" +
		"    return duration.Seconds() < 0\n" +
		"}\n" +
		"\n" +
		"// After returns true if `a` is after `b`.\n" +
		"func After(a, b Time) bool {\n" +
		"    duration = Sub(a, b)\n" +
		"    \n" +
		"    return duration.Seconds() > 0\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("duration.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("import \"math\"\n" +
		"\n" +
		"// These constants only go to Hour because after then the durations get more\n" +
		"// ambiguous.\n" +
		"Nanosecond  = 0.000000001\n" +
		"Microsecond = 0.000001\n" +
		"Millisecond = 0.001\n" +
		"Second      = 1\n" +
		"Minute      = 60\n" +
		"Hour        = 3600\n" +
		"\n" +
		"// A Duration represents a length of time.\n" +
		"func Duration(seconds number) Duration {\n" +
		"    func Nanoseconds() number {\n" +
		"        return ^seconds / Nanosecond\n" +
		"    }\n" +
		"\n" +
		"    func Microseconds() number {\n" +
		"        return ^seconds / Microsecond\n" +
		"    }\n" +
		"\n" +
		"    func Milliseconds() number {\n" +
		"        return ^seconds / Millisecond\n" +
		"    }\n" +
		"\n" +
		"    func Seconds() number {\n" +
		"        return ^seconds\n" +
		"    }\n" +
		"\n" +
		"    func Minutes() number {\n" +
		"        return ^seconds / Minute\n" +
		"    }\n" +
		"\n" +
		"    func Hours() number {\n" +
		"        return ^seconds / Hour\n" +
		"    }\n" +
		"\n" +
		"    func String() string {\n" +
		"        s = \"\"\n" +
		"        seconds = ^seconds\n" +
		"\n" +
		"        if seconds >= 3600 {\n" +
		"            hours = math.Floor(seconds / 3600)\n" +
		"            seconds -= hours * 3600\n" +
		"            s += \"{ hours }h\"\n" +
		"        }\n" +
		"\n" +
		"        if seconds >= 60 {\n" +
		"            minutes = math.Floor(seconds / 60)\n" +
		"            seconds -= minutes * 60\n" +
		"            s += \"{ minutes }m\"\n" +
		"        }\n" +
		"\n" +
		"        if seconds != 0 {\n" +
		"            s += \"{ seconds }s\"\n" +
		"        }\n" +
		"\n" +
		"        return s\n" +
		"    }\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("sleep.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Sleep will pause the execution for a specific duration of time.\n" +
		"func Sleep(duration Duration) {\n" +
		"    __sleep(duration.Seconds())\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("time.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("January = 1\n" +
		"February = 2\n" +
		"March = 3\n" +
		"April = 4\n" +
		"May = 5\n" +
		"June = 6\n" +
		"July = 7\n" +
		"August = 8\n" +
		"September = 9\n" +
		"October = 10\n" +
		"November = 11\n" +
		"December = 12\n" +
		"\n" +
		"// A Time represents a single point in time. The Second may be fractional.\n" +
		"func Time(Year, Month, Day, Hour, Minute, Second number) Time {\n" +
		"    func String() string {\n" +
		"        // TODO(elliot): This is a known bug with how out-of-scope variables are\n" +
		"        //  currently handled. See the todo in compiler/expr.go for\n" +
		"        //  ast.Indentifier. We need to copy the values into this scope as a\n" +
		"        //  work-around for now.\n" +
		"        month = ^Month\n" +
		"        day = ^Day\n" +
		"        hour = ^Hour\n" +
		"        minute = ^Minute\n" +
		"        second = ^Second\n" +
		"\n" +
		"        return \"{^Year}-{zeroPad(month)}-{zeroPad(day)} {zeroPad(hour)}:{zeroPad(minute)}:{zeroPad(second)}\"\n" +
		"    }\n" +
		"}\n" +
		"\n" +
		"// Now returns the current time.\n" +
		"func Now() Time {\n" +
		"    year, month, day, hour, minute, second = __now()\n" +
		"\n" +
		"    return Time(year, month, day, hour, minute, second)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("unix.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// Unix returns the unix timestamp, the number of seconds elapsed since January\n" +
		"// 1, 1970 UTC. The value returned may be fractional.\n" +
		"func Unix(t Time) number {\n" +
		"    return __unix(t)\n" +
		"}\n" +
		"\n" +
		"// FromUnix performs the opposite operation as Unix. It receives the number of\n" +
		"// seconds elapsed since January 1, 1970 UTC and returns the Time. The input\n" +
		"// seconds may be fractional.\n" +
		"func FromUnix(seconds number) Time {\n" +
		"    year, month, day, hour, minute, second = __fromunix(seconds)\n" +
		"\n" +
		"    return Time(year, month, day, hour, minute, second)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("util.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("func zeroPad(n number) string {\n" +
		"    if n < 10 {\n" +
		"        return \"0\" + string n\n" +
		"    }\n" +
		"\n" +
		"    return string n\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/time")
	fs = memfs.Create()
	f, _ = fs.OpenFile("is.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// IsControl reports whether the character is a control character.\n" +
		"func IsControl(c char) bool {\n" +
		"    return __unicode_is(\"Control\", c)\n" +
		"}\n" +
		"\n" +
		"// IsDigit reports whether the character is a decimal digit.\n" +
		"func IsDigit(c char) bool {\n" +
		"    return __unicode_is(\"Digit\", c)\n" +
		"}\n" +
		"\n" +
		"// IsGraphic reports whether the character is defined as a Graphic by Unicode.\n" +
		"// Such characters include letters, marks, numbers, punctuation, symbols, and\n" +
		"// spaces, from categories L, M, N, P, S, Zs.\n" +
		"func IsGraphic(c char) bool {\n" +
		"    return __unicode_is(\"Graphic\", c)\n" +
		"}\n" +
		"\n" +
		"// IsLetter reports whether the character is a letter (category L).\n" +
		"func IsLetter(c char) bool {\n" +
		"    return __unicode_is(\"Letter\", c)\n" +
		"}\n" +
		"\n" +
		"// IsLower reports whether the character is a lower case letter.\n" +
		"func IsLower(c char) bool {\n" +
		"    return __unicode_is(\"Lower\", c)\n" +
		"}\n" +
		"\n" +
		"// IsMark reports whether the character is a mark character (category M).\n" +
		"func IsMark(c char) bool {\n" +
		"    return __unicode_is(\"Mark\", c)\n" +
		"}\n" +
		"\n" +
		"// IsNumber reports whether the character is a number (category N).\n" +
		"func IsNumber(c char) bool {\n" +
		"    return __unicode_is(\"Number\", c)\n" +
		"}\n" +
		"\n" +
		"// IsPrint reports whether the character is defined as printable. Such\n" +
		"// characters include letters, marks, numbers, punctuation, symbols, and the\n" +
		"// ASCII space character, from categories L, M, N, P, S and the ASCII space\n" +
		"// character.\n" +
		"//\n" +
		"// This categorization is the same as IsGraphic except that the only spacing\n" +
		"// character is ASCII space, U+0020.\n" +
		"func IsPrint(c char) bool {\n" +
		"    return __unicode_is(\"Print\", c)\n" +
		"}\n" +
		"\n" +
		"// IsPunct reports whether the character is a Unicode punctuation character\n" +
		"// (category P).\n" +
		"func IsPunct(c char) bool {\n" +
		"    return __unicode_is(\"Punct\", c)\n" +
		"}\n" +
		"\n" +
		"// IsSpace reports whether the character is a space character as defined by\n" +
		"// Unicode's White Space property; in the Latin-1 space this is '\\t', '\\n',\n" +
		"// '\\v', '\\f', '\\r', ' ', U+0085 (NEL), U+00A0 (NBSP).\n" +
		"//\n" +
		"// Other definitions of spacing characters are set by category Z and property\n" +
		"// Pattern_White_Space.\n" +
		"func IsSpace(c char) bool {\n" +
		"    return __unicode_is(\"Space\", c)\n" +
		"}\n" +
		"\n" +
		"// IsSymbol reports whether the character is a symbolic character.\n" +
		"func IsSymbol(c char) bool {\n" +
		"    return __unicode_is(\"Symbol\", c)\n" +
		"}\n" +
		"\n" +
		"// IsTitle reports whether the character is a title case letter.\n" +
		"func IsTitle(c char) bool {\n" +
		"    return __unicode_is(\"Title\", c)\n" +
		"}\n" +
		"\n" +
		"// IsUpper reports whether the character is an upper case letter.\n" +
		"func IsUpper(c char) bool {\n" +
		"    return __unicode_is(\"Upper\", c)\n" +
		"}\n" +
		""))
	f, _ = fs.OpenFile("to.ok", os.O_RDWR|os.O_CREATE, 0777)
	f.Write([]byte("// ToLower maps the character to lower case.\n" +
		"func ToLower(c char) char {\n" +
		"    return __unicode_to(\"Lower\", c)\n" +
		"}\n" +
		"\n" +
		"// ToUpper maps the character to upper case.\n" +
		"func ToUpper(c char) char {\n" +
		"    return __unicode_to(\"Upper\", c)\n" +
		"}\n" +
		"\n" +
		"// ToTitle maps the character to title case.\n" +
		"func ToTitle(c char) char {\n" +
		"    return __unicode_to(\"Title\", c)\n" +
		"}\n" +
		""))
	Filesystem.Mount(fs, "/unicode")
}
