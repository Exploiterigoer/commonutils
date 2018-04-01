package pprofutil

import (
	"os"
	"runtime/pprof"

	"github.com/Exploiterigoer/commonutils/logutil"
)

// CPUAnalyze pprof for CPU
func CPUAnalyze(profileName string) {
	f, err := os.OpenFile(profileName+".prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logutil.LogInformation("logging in error for CPU")
	}
	pprof.StartCPUProfile(f)

	// something code here

	pprof.StopCPUProfile()
	f.Close()
}

// MemoryAnalyze pprof for memory
func MemoryAnalyze(profileName string) {
	f, err := os.OpenFile(profileName+".prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logutil.LogInformation("logging in error for memory")
	}
	pprof.WriteHeapProfile(f)

	// something code here

	f.Close()
}

// GthbAnalyze pprof for goroutine threadcreate heap block
func GthbAnalyze(profileType string, debug int) {
	profileName := profileType + ".prof"
	f, err := os.OpenFile(profileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logutil.LogInformation(err)
	}
	if err = pprof.Lookup(profileType).WriteTo(f, debug); err != nil {
		logutil.LogInformation("logging in error for goroutine")
	}
	f.Close()
}
