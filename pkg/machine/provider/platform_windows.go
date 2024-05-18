package provider

import (
	"fmt"
	"os"

	"github.com/pycabbage/podman/v5/pkg/machine/vmconfigs"
	"github.com/pycabbage/podman/v5/pkg/machine/wsl"

	"github.com/containers/common/pkg/config"
	"github.com/pycabbage/podman/v5/pkg/machine/define"
	"github.com/pycabbage/podman/v5/pkg/machine/hyperv"
	"github.com/sirupsen/logrus"
)

func Get() (vmconfigs.VMProvider, error) {
	cfg, err := config.Default()
	if err != nil {
		return nil, err
	}
	provider := cfg.Machine.Provider
	if providerOverride, found := os.LookupEnv("CONTAINERS_MACHINE_PROVIDER"); found {
		provider = providerOverride
	}
	resolvedVMType, err := define.ParseVMType(provider, define.WSLVirt)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("Using Podman machine with `%s` virtualization provider", resolvedVMType.String())
	switch resolvedVMType {
	case define.WSLVirt:
		return new(wsl.WSLStubber), nil
	case define.HyperVVirt:
		if !wsl.HasAdminRights() {
			return nil, fmt.Errorf("hyperv machines require admin authority")
		}
		return new(hyperv.HyperVStubber), nil
	default:
		return nil, fmt.Errorf("unsupported virtualization provider: `%s`", resolvedVMType.String())
	}
}
