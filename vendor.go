package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/codegangsta/cli"
)

func vendor(c *cli.Context) {
	repo := c.String("repo")
	gopath := c.String("gopath")

	absRepo, err := filepath.Abs(repo)
	if err != nil {
		log.Fatal("could not resolve repo: " + err.Error())
	}

	rootPkgs := map[string]*GoPkg{}
	deps := c.Args()
	for _, dep := range deps {
		pkg := NewGoPkg(gopath, dep)
		rootPkg, err := pkg.RootPkg()
		if err != nil {
			log.Fatalf("Unable to get root package for '%s': %s", pkg.Name, err.Error())
		}
		rootPkgs[rootPkg.Name] = rootPkg
	}

	for _, pkg := range rootPkgs {
		pkgDestination := filepath.Join("vendor", pkg.Name)

		fi, err := os.Lstat(pkgDestination)
		if err == nil && fi.IsDir() {
			fmt.Printf("\x1b[33mskipping %s (already vendored)\x1b[0m\n", pkg.Name)
			continue
		}

		fmt.Println("\x1b[32madding " + pkg.Name + "\x1b[0m")

		origin, err := pkg.HttpsOrigin()
		if err != nil {
			log.Fatalf("error finding git origin of package '%s': %s", pkg.Name, err.Error())
		}

		addSubmoduleCmd := exec.Command("git", "submodule", "add", origin, pkgDestination)
		addSubmoduleCmd.Dir = absRepo

		err = addSubmoduleCmd.Run()
		if err != nil {
			log.Fatal("error adding submodule: " + err.Error())
		}
	}
}
