package ldflags

import (
	"github.com/cdvelop/git"
	"github.com/cdvelop/strings"
)

// other_vars ej: main.encryption_key:123
func Add(other_vars map[string]string) (flags, err string) {
	const this = "ldflags Add error "

	var ldflags []string

	tag, err := git.GetLatestTag()
	if err != "" {
		return "", this + err
	}
	ldflags = append(ldflags, `-X 'main.app_version=`+tag+`'`)

	go_version, err := GoVersion()
	if err != "" {
		return "", this + err
	}
	ldflags = append(ldflags, `-X 'main.go_version=`+go_version+`'`)

	for k, v := range other_vars {
		ldflags = append(ldflags, `-X '`+k+`=`+v+`'`)
	}

	return strings.Join(ldflags, " "), ""
}
