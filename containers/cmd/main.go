package main 

import (
	"fmt"
	"os"
	"os/exec"
)



func main(){
  switch os.Args[1]{
  case "run":
	run()
  }
  default:
	panic("run is the only valid option for the moment")
}


fund run (){
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())
	
	cmd := exec.Command()
	cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func must (err error){
	if err != nil {
		panic(err)
	}
}