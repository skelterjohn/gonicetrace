//target:gonicetrace.googlecode.com/hg/nicetrace

//Pretty print a recovered stack trace.
package nicetrace

import (
	"runtime"
	"os"
	"fmt"
)

func Print() {
	e := recover()
	if e != nil {
		fmt.Fprintf(os.Stderr, "recover: %v\n", e)
		for skip:=1; ; skip++ {
			 pc, file, line, ok := runtime.Caller(skip)
			 if !ok {
				break
			 }
			 if file[len(file)-1] == 'c' {
				continue
			 }
			 f := runtime.FuncForPC(pc)
			 fmt.Fprintf(os.Stderr, "%s:%d %s()\n", file, line, f.Name())
		}
	}
}
