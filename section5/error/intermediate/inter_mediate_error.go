package intermediate

import (
	"os/exec"

	e "github.com/yuizho/concurrent-programming-by-go/section5/error/error"
	"github.com/yuizho/concurrent-programming-by-go/section5/error/lowlevel"
)

type IntermediateErr struct {
	error
}

func RunJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := lowlevel.IsGloballyExec(jobBinPath)
	if err != nil {
		return err // ❶
	} else if isExecutable == false {
		return e.WrapError(nil, "job binary is not executable")
	}

	return exec.Command(jobBinPath, "--id="+id).Run() // ❶
}
