package gitx

import (
	"os/exec"

	"github.com/samber/lo"
)

func GoFmt() {
	lo.Must0(exec.Command("gofmt", "-w", ".").Run())
}
