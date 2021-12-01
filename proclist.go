package realhttpworld

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func Proclist() {
	pid := strconv.Itoa(os.Getpid())
	out, _ := exec.Command("ls", "-al", fmt.Sprintf("/proc/%s/fd/", pid)).Output()
	fmt.Println(string(out))
}
