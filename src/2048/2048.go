package main

import (
  "github.com/nsf/termbox-go"
  "strconv"
  "os"
  "math/rand"
  "time"
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

func genCode() (ret [16]int) {
  return
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

func (code *typeCode) changeWithInput() {
  switch ev := termbox.PollEvent(); ev.Type {
  case termbox.EventKey:
    switch ev.Key {
    case termbox.KeyEsc:
      termbox.Close()
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
  x := rand.Intn(4)
  y := rand.Intn(4)
  code[x][y] = 1
}

func (code *typeCode) changeLeft() {
}
func (code *typeCode) changeRight() {
}
func (code *typeCode) changeUp() {
}
func (code *typeCode) changeDown() {
}
