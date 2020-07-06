package ctx

import (
	"go.uber.org/fx"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type ResolveFileName struct {
}

func (self *ResolveFileName) Resolve(path string) (string, error) {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	workingDir, _ := os.Getwd()
	if path == "" {
		return path, nil
	}

	if !filepath.IsAbs(path) {
		if strings.HasPrefix(path, "~/") {
			path = filepath.Join(homeDir, path[2:])
		} else {
			path = filepath.Join(workingDir, path)
		}
	}
	return path, nil
}

func ProvideResolveFileName(v *ResolveFileName) fx.Option {
	return fx.Provide(
		func() *ResolveFileName {
			return v
		})
}
