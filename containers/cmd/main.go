package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"strconv"
	"path/filepath"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("run is the only valid option for the moment")
	}
}

func run() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())
	cg()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/home/etcheve/code-examples/containers/fake-fs"))
	must(syscall.Chdir("/")) 
	must(syscall.Mount("proc","proc","proc",0,""))
	must(cmd.Run())
	must(syscall.Unmount("proc",0))
}


func cg() {
    base := "/sys/fs/cgroup"
    path := filepath.Join(base, "test_container_go")

    // create the cgroup directory
    if err := os.Mkdir(path, 0755); err != nil && !os.IsExist(err) {
    	panic(err)
    }
    // enable pids controller in parent
    // write "+pids" into cgroup.subtree_control
    must(os.WriteFile(filepath.Join(base, "cgroup.subtree_control"), []byte("+pids"), 0o644))

    // now you can set pids.max
    must(os.WriteFile(filepath.Join(path, "pids.max"), []byte("10"), 0o644))

    // add current process to this cgroup
    pid := strconv.Itoa(os.Getpid())
    must(os.WriteFile(filepath.Join(path, "cgroup.procs"), []byte(pid), 0o644))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
