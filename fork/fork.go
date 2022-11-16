package fork

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

/*
	cmd, err := fork.Fork()
	if err != nil {
		panic(err)
	}
	log.Printf("fork pid=%v name=%v\n",
		cmd.Process.Pid, filepath.Base(os.Args[0]))
*/

// reboot
func Fork() (*exec.Cmd, error) {
	path, err := os.Executable() // 与 os.Args[0] 类似
	if err != nil {
		return nil, err
	}
	cmd := &exec.Cmd{
		Path:        path,               // 文件路径 包括文件名
		Args:        []string{path},     // 执行的命令
		Dir:         filepath.Dir(path), // 文件路径 不包括文件名
		Env:         os.Environ(),       // 环境变量
		Stdin:       os.Stdin,           // 输入
		Stdout:      os.Stdout,          // 输出
		Stderr:      os.Stderr,          // 错误
		SysProcAttr: &syscall.SysProcAttr{},
		//	SysProcAttr: &syscall.SysProcAttr{
		//	Pdeathsig: syscall.SIGTERM,
		// },
	}
	err = cmd.Start() // cmd.Start() 不阻塞地执行, cmd.Run() 阻塞地执行
	if err != nil {
		return nil, err
	}
	//	if err := cmd.Wait(); err != nil { // 等待
	//	log.Panicf("failed to wait command: %s", err)
	// }
	return cmd, nil
}
