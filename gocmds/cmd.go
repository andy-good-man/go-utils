// Package gocmds 执行终端命令,获取执行结果.
package gocmds

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"unsafe"
)

// Cmder 一个执行命令实例
type Cmder struct {
	terminal []string
}

func (c *Cmder) getCmd(arg ...string) *exec.Cmd {
	if c.terminal == nil || len(c.terminal) < 2 {
		c.terminal = make([]string, 2)
		if runtime.GOOS == "windows" {
			c.terminal[0] = "cmd"
			c.terminal[1] = "/c"
		} else {
			c.terminal[0] = "/bin/sh"
			c.terminal[1] = "-c"
		}
	}
	var cmdArg []string
	cmdArg = append(cmdArg, c.terminal...)
	cmdArg = append(cmdArg, arg...)
	return exec.Command(cmdArg[0], cmdArg[1:]...)
}

// Run 执行cmd并获取执行结果.设置等待完成执行的timeout,一旦timeout便杀死进程.
func (c *Cmder) Run(cmdLine string, timeout ...int) *Result {
	cmd := c.getCmd(cmdLine)
	ret := new(Result)

	cmd.Stdout = &ret.buf
	cmd.Stderr = &ret.buf
	cmd.Env = os.Environ()

	ret.err = cmd.Start()
	if ret.err != nil {
		return ret
	}

	if len(timeout) == 0 || timeout[0] <= 0 { // ... timeout可选参数，目前来看没太大必要
		ret.err = cmd.Wait()
		return ret
	}

	timer := time.NewTimer(time.Duration(timeout[0]) * time.Second)
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	select {
	case ret.err = <-done:
		timer.Stop()
	case <-timer.C:
		if err := cmd.Process.Kill(); err != nil {
			ret.err = fmt.Errorf("command timed out and killing process fail: %s", err.Error())
		} else {
			// 杀死进程后,等待cmd.Wait()的结果返回
			<-done
			ret.err = errors.New("command timed out")
		}
	}
	return ret
}

// Result 执行结果
type Result struct {
	buf bytes.Buffer
	err error
	str *string
}

// Err 返回错误日志
func (r *Result) Err() error {
	if r.err == nil {
		return nil
	}
	r.err = errors.New(r.String())
	return r.err
}

// String 返回执行日志
func (r *Result) String() string {
	if r.str == nil {
		b := bytes.TrimSpace(r.buf.Bytes())
		if r.err != nil {
			b = append(b, ' ', '(')
			b = append(b, r.err.Error()...)
			b = append(b, ')')
		}
		r.str = (*string)(unsafe.Pointer(&b))
	}
	return *r.str
}
