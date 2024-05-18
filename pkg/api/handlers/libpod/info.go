package libpod

import (
	"net/http"

	"github.com/pycabbage/podman/v5/libpod"
	"github.com/pycabbage/podman/v5/pkg/api/handlers/utils"
	api "github.com/pycabbage/podman/v5/pkg/api/types"
	"github.com/pycabbage/podman/v5/pkg/domain/infra/abi"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	runtime := r.Context().Value(api.RuntimeKey).(*libpod.Runtime)
	containerEngine := abi.ContainerEngine{Libpod: runtime}
	info, err := containerEngine.Info(r.Context())
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, info)
}
