package process

import (
	"fmt"
	"os"
	"runtime"
)

func Run() {
	println("=====================================")
	println("================procs================")
	println("=====================================")

	fmt.Println("CPU Num:", runtime.NumCPU())

	nprocs := os.Getenv("GOMAXPROCS")
	fmt.Println("Env GOMAXPROCS:", nprocs)

	nprocs_runtime := runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	fmt.Println("Runtime GOMAXPROCS:", nprocs_runtime)
}
