package main

import (
  "fmt"
  "os"
  "strconv"

  "github.com/elias-boemeke/sudoku-wizard/sudoku"
)

func showHelp() {
  help :=
`Welcome to Sudoku-Wizard!

USAGE:
sw help     display help
sw read     read sudoku and output to stdout or file
sw solve    solve sudoku`

  fmt.Println(help)
}

// MAIN
func main() {

  if len(os.Args) < 2 {
    // no argument
    showHelp()
  } else {
    option := os.Args[1]
    switch option {
    case "solve":
      fmt.Println("solve")
    case "read":
      fmt.Println("read")
      s := sudoku.NewSudoku()
      fmt.Println(s)
    case "help":
      showHelp()
    default:
      panic("unknown option '" + option + "'. run sw help for more info")
    }
  }
}
