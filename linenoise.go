package linenoise

/*
#include <stdlib.h>
#include "linenoise.h"

extern void goCompletionEntry(char * buf, linenoiseCompletions * linenoiseCompletions);

*/
import "C"
import "unsafe"

var cc CompletionCallback

type CompletionCallback func(buf string, linenoiseCompletions *C.struct_linenoiseCompletions)

//export goCompletionEntry
func goCompletionEntry(buf * C.char, linenoiseCompletions *C.struct_linenoiseCompletions) {
  if cc != nil {
    cc(C.GoString(buf), linenoiseCompletions)
    return
  }
  return
}

func linenoiseAddCompletion(linenoiseCompletions *C.struct_linenoiseCompletions, completion string) {
  c := C.CString(completion)
  defer C.free(unsafe.Pointer(c))
  C.linenoiseAddCompletion(linenoiseCompletions, c)
}

func linenoise(prompt string) (line string, end bool) {
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

func linenoiseHistoryAdd(line string) (done bool) {
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
