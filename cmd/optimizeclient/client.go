package main

import (
	"fmt"
	"liangrun/internal/optimizeclient"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

func main() {
	if os.Getenv("ENABLEPROFILE") == "1" {
		profFileName := fmt.Sprintf("/root/profile_%v", os.Getenv("TESTCASE"))
		traceFileName := fmt.Sprintf("/root/trace_%v", os.Getenv("TESTCASE"))
		profF, _ := os.Create(profFileName)
		traceF, _ := os.Create(traceFileName)
		pprof.StartCPUProfile(profF)
		trace.Start(traceF)
		defer pprof.StopCPUProfile()
		defer trace.Stop()
	}
	optimizeclient.StartClient()
}
