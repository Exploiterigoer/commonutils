package pprofutil

import (
	"commonutils/logutil"
	"os"
	"runtime/pprof"
)

// CPU pprof 分析
func CPUAnalyze(profileName string) {
	f, err := os.OpenFile(profileName+".prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logutil.LogInformation("记录CPU信息出错")
	}
	pprof.StartCPUProfile(f)

	// something code here

	pprof.StopCPUProfile()
	f.Close()
}

// memory pprof 分析
func MemoryAnalyze(profileName string) {
	f, err := os.OpenFile(profileName+".prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logutil.LogInformation("记录内存信息出错")
	}
	pprof.WriteHeapProfile(f)

	// something code here

	f.Close()
}

// goroutine threadcreate heap block 分析
func GthbAnalyze(profileType string, debug int) {
	profileName := profileType + ".prof"
	f, err := os.OpenFile(profileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logutil.LogInformation(err)
	}
	if err = pprof.Lookup(profileType).WriteTo(f, debug); err != nil {
		logutil.LogInformation("记录pprof.Lookup信息出错")
	}
	f.Close()
}
