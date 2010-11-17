package nicetrace

import (
	"runtime"
	"os"
	"fmt"
)

func Print() {
	e := recover()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		for skip:=1; ; skip++ {
			 _, file, line, ok := runtime.Caller(skip)
			 if !ok {
				break
			 }
			 if file[len(file)-1] == 'c' {
				continue
			 }
			 fmt.Fprintf(os.Stderr, "%s:%d\n", file, line)
		}
	}
}
