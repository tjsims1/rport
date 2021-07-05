//+build !windows

package chclient

import (
	"context"
	"os/exec"
)

func (e *CmdExecutorImpl) New(ctx context.Context, execCtx *CmdExecutorContext) *exec.Cmd {
	return e.newCmd(ctx, execCtx)
}

func (e *CmdExecutorImpl) DecodeOutput(output string) (string, error) {
	return output, nil
}
