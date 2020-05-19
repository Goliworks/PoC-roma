package utils

import (
	"os/user"
	"path/filepath"
	"strings"
)

func AbsPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		usr, _ := user.Current()
		hDir := usr.HomeDir
		return filepath.Join(hDir, path[2:])
	}
	p, _ := filepath.Abs(path)
	return p
}
