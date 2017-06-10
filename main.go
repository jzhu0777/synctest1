package main

import (
	"fmt"
	"os"
	"os/exec"

	"strconv"

	"github.com/juju/errors"
)

func main() {
	n := 10000
	var cmd = ""
	fmt.Println("start")
	for n < 20000 {

		cmd = "echo hello > ./" + strconv.Itoa(n) + ".txt"
		_ = RunCMD(cmd)
		fmt.Println(n)
		n++
	}
	fmt.Println("done")
}

func RunCMD(cmd string) (err error) {
	o := exec.Command("bash", "-c", cmd)
	o.Stdout = os.Stdout
	o.Stderr = os.Stderr
	return errors.Trace(o.Run())
}
