package main

// #include <stdio.h>
// #include <stdlib.h>
// void myprint(char* s) {
//		printf("%s\n", s);
// }
import "C"
import "unsafe"
import (
	log "github.com/cihub/seelog"
)

func Print(s string) {
	cs := C.CString(s)
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func main() {
	log.Flush()
	log.Info("begining")
	Print("aaaa")
	log.Info("ending")
}
