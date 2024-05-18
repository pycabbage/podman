package main

import (
	"fmt"

	"github.com/pycabbage/podman/v5/version"
)

func main() {
	fmt.Print(version.Version.String())
}
