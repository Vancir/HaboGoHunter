package osutil

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

const (
	DefaultDirPermission  = 0755
	DefaultFilePermission = 0644
	DefaultExecPermission = 0755
)

func setPdeathsig(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = new(syscall.SysProcAttr)
	}
	cmd.SysProcAttr.Pdeathsig = syscall.SIGKILL
	cmd.SysProcAttr.Setpgid = true
}

func killPgroup(cmd *exec.Cmd) {
	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
}

// Command is similar to os/exec.Command, but also sets PDEATHSIG on linux.
func Command(bin string, args ...string) *exec.Cmd {
	cmd := exec.Command(bin, args...)
	setPdeathsig(cmd)
	return cmd
}

func RunCmd(timeout time.Duration, dir, bin string, args ...string) ([]byte, error) {
	cmd := Command(bin, args...)
	cmd.Dir = dir
	return run(timeout, cmd)
}

func run(timeout time.Duration, cmd *exec.Cmd) ([]byte, error) {
	output := new(bytes.Buffer)
	if cmd.Stdout == nil {
		cmd.Stdout = output
	}
	if cmd.Stderr == nil {
		cmd.Stderr = output
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start %v %+v: %v", cmd.Path, cmd.Args, err)
	}

	done := make(chan bool)
	timedout := make(chan bool, 1)
	timer := time.NewTimer(timeout)
	go func() {
		select {
		case <-timer.C:
			timedout <- true
			killPgroup(cmd)
			cmd.Process.Kill()
		case <-done:
			timedout <- false
			timer.Stop()
		}
	}()
	err := cmd.Wait()
	close(done)
	if err != nil {
		text := fmt.Sprintf("failed to run %q: %v", cmd.Args, err)
		if <-timedout {
			text = fmt.Sprintf("timedout %q", cmd.Args)
		}
		return output.Bytes(), errors.New(text)
	}
	return output.Bytes(), nil
}
