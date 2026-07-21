// Command create_mod packages a directory into a proxy.golang.org-format
// module zip using golang.org/x/mod/zip.CreateFromDir.
//
// Usage: create_mod <module-path> <version> <src-dir> <out.zip>
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("usage: create_mod <module-path> <version> <src-dir> <out.zip>")
	}
	modPath, version, dir, out := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(out)
	if err != nil {
		log.Fatalf("create %s: %v", out, err)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, dir); err != nil {
		log.Fatalf("CreateFromDir: %v", err)
	}
	log.Printf("wrote module zip %s (%s %s)", out, modPath, version)
}
