package sudoku


type Sudoku struct {
  n [9][9]rune
}

func NewSudoku() *Sudoku {
  s := new(Sudoku)
  for i, line := range s.n {
    for j := range line {
      s.n[i][j] = '_'
    }
  }

  return s
}

func (s Sudoku) String() string {
  var sr string
  for i, line := range s.n {
    switch i {
    case 0, 3, 6:
      sr = sr + "+-----+-----+-----+\n"
    }

    for j := range line {
      v := string(s.n[i][j])
      switch j {
      case 0, 3, 6:
        sr = sr + "|" + v
      case 1, 2, 4, 5, 7:
        sr = sr + " " + v
      case 8:
        sr = sr + " " + v + "|"
      default:
        sr = sr + v
      }
    }

    if i < 8 {
      sr = sr + "\n"
    } else {
      sr = sr + "\n+-----+-----+-----+"
    }
  }

  return sr
}
