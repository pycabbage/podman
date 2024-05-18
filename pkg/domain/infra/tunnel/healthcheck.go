package tunnel

import (
	"context"

	"github.com/pycabbage/podman/v5/libpod/define"
	"github.com/pycabbage/podman/v5/pkg/bindings/containers"
	"github.com/pycabbage/podman/v5/pkg/domain/entities"
)

func (ic *ContainerEngine) HealthCheckRun(ctx context.Context, nameOrID string, options entities.HealthCheckOptions) (*define.HealthCheckResults, error) {
	return containers.RunHealthCheck(ic.ClientCtx, nameOrID, nil)
}
