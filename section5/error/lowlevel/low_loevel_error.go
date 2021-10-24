package lowlevel

import (
	"os"

	e "github.com/yuizho/concurrent-programming-by-go/section5/error/error"
)

type LowLevelErr struct {
	error
}

func IsGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{(e.WrapError(err, err.Error()))}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}
