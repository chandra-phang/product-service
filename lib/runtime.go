package lib

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// GetFunctionNameInSource gets the name of the given function from the source code.
// Example:
//
//	func alice() {}
//	func main()  { fmt.Println(GetFunctionNameInSource(alice)) }
//
// Output: "alice"
func GetFunctionNameInSource(theFunction interface{}) string {
	ptr := reflect.ValueOf(theFunction).Pointer()
	fn := runtime.FuncForPC(ptr)
	if fn != nil {
		return stripMangler(fn.Name())
	}
	return "(function was nil)"
}

// WhoCalledMe returns the file+lineNo of the file that called the caller.
// Example:
//
//	/* a.go */ func alice() string { return WhoCalledMe() }
//	/* b.go */ func bob()   string { return alice() }
//	/* main */ func main()         { fmt.Println(bob()) }
//
// Output: "b.go:1"
func WhoCalledMe() string {
	_, file, lineNo, ok := runtime.Caller(2)
	if !ok {
		return "(failed to get caller)"
	}
	file = trimpath(file)
	callerSourceFile := fmt.Sprintf("%s:%d", file, lineNo)
	return callerSourceFile
}

// stripMangler strips a suffix added by the compiler to mangle function names.
// See https://stackoverflow.com/q/32925344
func stripMangler(functionName string) string {
	// mangler is suffixed to the names for receiver methods by the compiler.
	const mangler = "-fm"
	return strings.TrimSuffix(functionName, mangler)
}

func trimpath(path string) string {
	// Trim the directory out of source file paths that are baked into in the app binary.
	// The outcome equivalent to -trimpath compiler flag. For example:
	//    /var/lib/jenkins/workspace/INTERIM_INFRA/BUILD/INVENTORY/apps/inventory/golang/src/.../wrapper/helper.go
	// Becomes:
	//    .../wrapper/helper.go
	const src = "/src/"
	i := strings.LastIndex(path, src)
	if i >= 0 {
		path = path[i+len(src):]
	}
	return path
}
