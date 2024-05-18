package env

import (
	"github.com/pycabbage/podman/v5/pkg/rootless"
	"github.com/pycabbage/podman/v5/pkg/util"
)

func getRuntimeDir() (string, error) {
	if !rootless.IsRootless() {
		return "/run", nil
	}
	return util.GetRootlessRuntimeDir()
}
