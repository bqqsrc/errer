package errer

import "fmt"

type callerErr struct {
	caller string
	format string
	args   []interface{}
}

func (e *callerErr) Error() string {
	fmtStr := fmt.Sprintf("call by: %s,\nerror is : %s\n", e.caller, e.format)
	return fmt.Sprintf(fmtStr, e.args...)
}

//test
func CallerErr(caller, format string, args ...interface{}) error {
	return &callerErr{caller, format, args}
}
