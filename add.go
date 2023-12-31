package ldflags

import (
	"github.com/cdvelop/git"
	"github.com/cdvelop/strings"
)

const repository = `github.com/cdvelop`

// other_vars ej: main.encryption_key:123
func Add(other_vars ...map[string]string) (flags, err string) {
	const this = "ldflags Add error "

	var ldflags []string

	tag, err := git.GetLatestTag()
	if err != "" {
		return "", this + err
	}
	ldflags = append(ldflags, `-X '`+repository+`/model.app_version=`+tag+`'`)

	ldflags = append(ldflags, `-X '`+repository+`/model.go_version=`+GoVersion()+`'`)

	ldflags = append(ldflags, `-X '`+repository+`/model.tinygo_version=`+TinyGoVersion()+`'`)

	for _, others := range other_vars {
		for k, v := range others {
			ldflags = append(ldflags, `-X '`+k+`=`+v+`'`)
		}
	}

	return strings.Join(ldflags, " "), ""
}
