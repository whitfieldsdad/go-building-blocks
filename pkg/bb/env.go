package bb

import (
	"os"
	"strings"
)

func GetEnvironmentVariables() map[string]string {
	m := map[string]string{}
	for _, e := range os.Environ() {
		p := strings.SplitN(e, "=", 2)
		k := p[0]
		v := p[1]
		m[k] = v
	}
	return m
}
