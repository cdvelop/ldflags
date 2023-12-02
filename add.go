package ldflags

import "github.com/cdvelop/git"

func Add() (flags, err string) {

	tag, err := git.GetLatestTag()
	if err != "" {
		return "", "ldflags error" + err
	}

	flags = `-X 'main.version=` + tag + `'`

	return
}
