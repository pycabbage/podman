package vmconfigs

import (
	"os"

	"github.com/pycabbage/podman/v5/pkg/machine/applehv/vfkit"
)

type AppleHVConfig struct {
	// The VFKit endpoint where we can interact with the VM
	Vfkit vfkit.VfkitHelper
}

// Stubs
type HyperVConfig struct{}
type WSLConfig struct{}
type QEMUConfig struct{}

func getHostUID() int {
	return os.Getuid()
}
