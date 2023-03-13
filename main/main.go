// gotest/cmd/libfoo/main.go
package main

import "C"

// other imports should be seperate from the special Cgo import
import (
	"fmt"
	"gotest/ucloud"
)

//export reverse
// func reverse(in *C.char) *C.char {
// 	return C.CString(foo.Reverse(C.GoString(in)))
// }

//export listVM
func listVM(limit int, offset int) *C.char {
	return C.CString(ucloud.ListVM(limit, offset))
}

//export getipnum
// func getipnum(in *C.char) *C.char {
// 	return C.CString(ucloud.Getipnum())
// }

//export getipnum
func getipnum() int {
	return ucloud.Getipnum()
}

func main() { fmt.Println("ggggg") }

//the "export getipnum" comment must be there, can you believe it?

// https://rogchap.com/2020/09/14/running-go-code-on-ios-and-android/

// CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CC="/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang -arch x86_64 -isysroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator15.5.sdk -mios-version-min=10.0" go build -buildmode=c-archive -tags ios -o foo.a ./
