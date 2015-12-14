package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GoPkg struct {
	gopath string
	Path   string
	Name   string
}

func NewGoPkg(gopath, dep string) *GoPkg {
	return &GoPkg{
		gopath: gopath,
		Path:   filepath.Join(gopath, "src", dep),
		Name:   dep,
	}
}

func (p *GoPkg) RootPkg() (*GoPkg, error) {
	gitToplevel := exec.Command("git", "rev-parse", "--show-toplevel")
	gitToplevel.Dir = p.Path

	buf := new(bytes.Buffer)

	gitToplevel.Stdout = buf
	gitToplevel.Stderr = os.Stderr

	err := gitToplevel.Run()
	if err != nil {
		return nil, err
	}

	toplevel := strings.TrimRight(buf.String(), "\n")

	goSrcPath := filepath.Join(p.gopath, "src") + string(os.PathSeparator)
	name := strings.TrimPrefix(toplevel, goSrcPath)
	rootPkg := &GoPkg{
		gopath: p.gopath,
		Path:   toplevel,
		Name:   name,
	}

	return rootPkg, nil
}

func (p *GoPkg) HttpsOrigin() (string, error) {
	gitOriginURL := exec.Command("git", "config", "--get", "remote.origin.url")
	gitOriginURL.Dir = p.Path

	buf := new(bytes.Buffer)

	gitOriginURL.Stdout = buf
	gitOriginURL.Stderr = os.Stderr

	err := gitOriginURL.Run()
	if err != nil {
		return "", err
	}

	uri := strings.TrimRight(buf.String(), "\n")

	return httpsOrigin(uri), nil
}

func (p *GoPkg) Branch() (string, error) {
	gitRevParse := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	gitRevParse.Dir = p.Path

	buf := new(bytes.Buffer)

	gitRevParse.Stdout = buf
	gitRevParse.Stderr = os.Stderr

	err := gitRevParse.Run()
	if err != nil {
		return "", err
	}

	rev := strings.TrimRight(buf.String(), "\n")
	if rev == "HEAD" {
		return "", nil
	}

	return rev, nil
}
