// Code generated by "stringer -type Code"; DO NOT EDIT.

package er

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UncaughtException-0]
}

const _Code_name = "UncaughtException"

var _Code_index = [...]uint16{0}

func (i Code) String() string {
	if i < 0 || i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}
