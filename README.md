ok
==

*ok* is a strongly-duck-typed language.


Quick Start
-----------

Since ok is so early in its development you will need the go compiler to run ok
programs:

```bash
go get -u github.com/elliotchance/ok
```

Programs in ok are directories containing one or more `.ok` files. You can run a
program by specifying the directory (or see the included tests/):

```bash
ok run my/program
```


Testing
-------

To run the full suite:

```bash
make test
```
