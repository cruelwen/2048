package main

import (
  "github.com/nsf/termbox-go"
  "strconv"
  "os"
  "math/rand"
  "time"
  "fmt"
)

type typeCode [4][4]int

func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  termbox.SetInputMode(termbox.InputEsc)
  rand.Seed(time.Now().UnixNano())

  var code typeCode
  for {
    code.print()
    code.changeWithInput()
    code.addNew()
  }
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
  for _, c := range msg {
    termbox.SetCell(x, y, c, fg, bg)
    x++
  }
}

func (code *typeCode) print() {
  const coldef = termbox.ColorDefault
  termbox.Clear(coldef, coldef)
  x,y:=0,0
  for _,line := range code {
    for _,number := range line {
      str := strconv.Itoa(number)
      tbprint(x,y,coldef,coldef,str)
      x++
    }
    y++
    x=0
  }
  termbox.Flush()
}

func (code typeCode) max() (max int) {
  max = 0
  for _,line := range code {
    for _,number := range line {
      if number > max {
        max = number
      }
    }
  }
  return max
}

func (code typeCode) min() (min int) {
  min = 10
  for _,line := range code {
    for _,number := range line {
      if number < min {
        min = number
      }
    }
  }
  return min
}

func (code *typeCode) changeWithInput() {
  switch ev := termbox.PollEvent(); ev.Type {
  case termbox.EventKey:
    switch ev.Key {
    case termbox.KeyEsc:
      termbox.Close()
      fmt.Println("Bye")
      os.Exit(0)
    case termbox.KeyArrowLeft:
      code.changeLeft()
    case termbox.KeyArrowRight:
      code.changeRight()
    case termbox.KeyArrowUp:
      code.changeUp()
    case termbox.KeyArrowDown:
      code.changeDown()
    default:
    }
  case termbox.EventError:
    panic(ev.Err)
  }
}

func (code *typeCode) addNew() {
  if code.min() != 0 {
    termbox.Close()
    fmt.Println("Lose")
    os.Exit(0)
  }
  if code.max() == 10 {
    termbox.Close()
    fmt.Println("Win")
    os.Exit(0)
  }
  x,y:=rand.Intn(4),rand.Intn(4)
  for code[x][y] != 0 {
    x,y = rand.Intn(4),rand.Intn(4)
  }
  code[x][y] = rand.Intn(2) + 1
}

func (code *typeCode) changeLeft() {
  for i:=0;i<4;i++ {
    for j:=0;j<4;j++ {
      for k:=j+1;k<4;k++ {
        if code[i][k] != 0 {
          switch {
          case code[i][j] == 0:
            code[i][j] = code[i][k]
            code[i][k] = 0
          case code[i][j] == code[i][k]:
            code[i][j] ++ 
            code[i][k] = 0
          }
          break
        }
      }
    }
  }
}
func (code *typeCode) changeRight() {
  for i:=0;i<4;i++ {
    for j:=3;j>=0;j-- {
      for k:=j-1;k>=0;k-- {
        if code[i][k] != 0 {
          switch {
          case code[i][j] == 0:
            code[i][j] = code[i][k]
            code[i][k] = 0
          case code[i][j] == code[i][k]:
            code[i][j] ++ 
            code[i][k] = 0
          }
          break
        }
      }
    }
  }
}
func (code *typeCode) changeUp() {
  for j:=0;j<4;j++ {
    for i:=0;i<4;i++ {
      for k:=i+1;k<4;k++ {
        if code[k][j] != 0 {
          switch {
          case code[i][j] == 0:
            code[i][j] = code[k][j]
            code[k][j] = 0
          case code[i][j] == code[k][j]:
            code[i][j] ++ 
            code[k][j] = 0
          }
          break
        }
      }
    }
  }
}
func (code *typeCode) changeDown() {
  for j:=0;j<4;j++ {
    for i:=3;i>=0;i-- {
      for k:=i-1;k>=0;k-- {
        if code[k][j] != 0 {
          switch {
          case code[i][j] == 0:
            code[i][j] = code[k][j]
            code[k][j] = 0
          case code[i][j] == code[k][j]:
            code[i][j] ++ 
            code[k][j] = 0
          }
          break
        }
      }
    }
  }
}
