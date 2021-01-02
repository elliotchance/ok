package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var f io.Writer

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	f, err := os.Create("fs/lib.go")
	check(err)

	fmt.Fprintf(f, "package fs\n\n")
	fmt.Fprintf(f, "import \"os\"\n")
	fmt.Fprintf(f, "import \"github.com/blang/vfs\"\n")
	fmt.Fprintf(f, "import \"github.com/blang/vfs/memfs\"\n")
	fmt.Fprintf(f, "import \"github.com/blang/vfs/mountfs\"\n")
	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\tFilesystem = mountfs.Create(vfs.OS())\n")
	fmt.Fprintf(f, "\tvar fs vfs.Filesystem\n")
	fmt.Fprintf(f, "\tvar f vfs.File\n")

	packages, err := ioutil.ReadDir("lib")
	check(err)

	for _, packagePathInfo := range packages {
		pkgName := packagePathInfo.Name()

		// lang is not importable - it only contains tests.
		if !packagePathInfo.IsDir() || pkgName == "lang" {
			continue
		}

		fmt.Fprintf(f, "\tfs = memfs.Create()\n")

		fileNames, err := ioutil.ReadDir("lib/" + pkgName)
		check(err)

		for _, filePathInfo := range fileNames {
			if !strings.HasSuffix(filePathInfo.Name(), ".ok") {
				continue
			}

			fileData, err := ioutil.ReadFile(fmt.Sprintf("lib/%s/%s", pkgName, filePathInfo.Name()))
			check(err)

			encodedFile := string(fileData)
			encodedFile = strings.ReplaceAll(encodedFile, "\\", "\\\\")
			encodedFile = strings.ReplaceAll(encodedFile, "\"", "\\\"")
			encodedFile = strings.ReplaceAll(encodedFile, "\n", "\\n\" +\n\t\t\t\"")

			fmt.Fprintf(f, "\tf, _ = fs.OpenFile(\"%s\", os.O_RDWR|os.O_CREATE, 0777)\n",
				filePathInfo.Name())
			fmt.Fprintf(f, "\tf.Write([]byte(\"%s\"))\n", encodedFile)
		}

		fmt.Fprintf(f, "\tFilesystem.Mount(fs, \"/%s\")\n", pkgName)
	}

	fmt.Fprintf(f, "}\n")
	fmt.Fprintf(f, "\n")
}
