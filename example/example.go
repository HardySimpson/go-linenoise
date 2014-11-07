package main

import (
  "fmt"
  "linenoise"
  "unsafe"
  "strings"
  "os"
)

// this func is called every time when people hit the tab button
func cb(buf string, lc unsafe.Pointer) {
  if !strings.HasPrefix(buf, "cd") { return }
  words := strings.Fields(buf)
  if len(words) == 1 { return }
  baseDir, _ := os.Open(".")
  names, _ := baseDir.Readdirnames(0)
  for _, n := range names {
    if strings.HasPrefix(n, words[1]) {
      linenoise.AddCompletion(lc, "cd " + n)
    }
  }
  return
}

func main() {
  linenoise.SetMultiLine(true)            //with multi line set,
                                          //the input will display in multi line when you type more than a line long
  linenoise.SetCompletionCallback(cb)
  linenoise.HistoryLoad("history.txt")    //load from disk
  linenoise.HistorySetMaxLen(10)          //max line keep in memory and disk
  for {
    line, end := linenoise.Scan("hello> ")
    if end {
      return
    }
    linenoise.HistoryAdd(line)            //add to memory
    linenoise.HistorySave("history.txt")  //add to disk
    fmt.Println(line)
    if line == "exit" || line == "quit" {
      break
    }
    if line == "clear" {
      linenoise.ClearScreen()
    }
  }
}
