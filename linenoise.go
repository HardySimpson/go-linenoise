package linenoise

/*
#include <stdlib.h>
#include "linenoise.h"

extern void goCompletionEntry(char * buf, void * lc);

*/
import "C"
import "unsafe"

var cc CompletionCallback

type CompletionCallback func(buf string, lc unsafe.Pointer)

//export goCompletionEntry
func goCompletionEntry(buf * C.char, lc unsafe.Pointer) {
  if cc != nil {
    cc(C.GoString(buf), lc)
    return
  }
  return
}

func SetCompletionCallback(fn CompletionCallback) {
  cc = fn
  return
}

func AddCompletion(lc unsafe.Pointer, completion string) {
  c := C.CString(completion)
  defer C.free(unsafe.Pointer(c))
  C.linenoiseAddCompletion(lc, c)
}

func Scan(prompt string) (line string, end bool) {
  pt := C.CString(prompt)
  defer C.free(unsafe.Pointer(pt))
  l := C.linenoise(pt)
  defer C.free(unsafe.Pointer(l))
  if l == nil {
    end = true
  } else {
    end = false
  }
  return C.GoString(l), end
}

func HistoryAdd(line string) (done bool) {
  l := C.CString(line)
  defer C.free(unsafe.Pointer(l))
  i := C.linenoiseHistoryAdd(l)
  if i == 1 {
    done = true
  } else {
    done = false
  }
  return done
}

func HistorySetMaxLen(length int) {
  C.linenoiseHistorySetMaxLen(C.int(length))
  return
}

func HistorySave(filename string) {
  f := C.CString(filename)
  defer C.free(unsafe.Pointer(f))
  C.linenoiseHistorySave(f)
  return
}

func HistoryLoad(filename string) {
  f := C.CString(filename)
  defer C.free(unsafe.Pointer(f))
  C.linenoiseHistoryLoad(f)
  return
}

func ClearScreen() {
  C.linenoiseClearScreen()
  return
}

func SetMultiLine(yes bool) {
  if yes {
    C.linenoiseSetMultiLine(C.int(1))
  } else {
    C.linenoiseSetMultiLine(C.int(0))
  }
  return
}

func PrintKeyCodes() {
  C.linenoisePrintKeyCodes()
  return
}
