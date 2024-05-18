package abi

import (
	"context"

	"github.com/pycabbage/podman/v5/pkg/autoupdate"
	"github.com/pycabbage/podman/v5/pkg/domain/entities"
)

func (ic *ContainerEngine) AutoUpdate(ctx context.Context, options entities.AutoUpdateOptions) ([]*entities.AutoUpdateReport, []error) {
	return autoupdate.AutoUpdate(ctx, ic.Libpod, options)
}
