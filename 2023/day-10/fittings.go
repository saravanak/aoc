package main

import (
	"fmt"
	"log"
	// "runtime/debug"
	"slices"

	"github.com/alecthomas/participle/v2/lexer"
)

type FittingBehaviour interface {
	opensRight() bool
	opensLeft() bool
	opensTop() bool
	opensBottom() bool
	// Gives the next fitting and the entry diretion for the next fitting (left|bottom|right|top)
	getNext(cellEntry string, pipeMap *PipeMap) (*Fitting, string)
	getAdjacentFittings(s *PipeMap) AdjCells
	row() int
	column() int
}

func getDirections(behviour FittingBehaviour) []string {
	var directions = make([]string, 0)

	if behviour.opensRight() {
		directions = append(directions, "right")
	}
	if behviour.opensLeft() {
		directions = append(directions, "left")
	}
	if behviour.opensTop() {
		directions = append(directions, "top")
	}
	if behviour.opensBottom() {
		directions = append(directions, "bottom")
	}
	return directions
}

type NonePipeFitting struct {
	BaseFitting
}

func (s NonePipeFitting) opensBottom() bool                           { return false }
func (s NonePipeFitting) opensRight() bool                            { return false }
func (s NonePipeFitting) opensLeft() bool                             { return false }
func (s NonePipeFitting) opensTop() bool                              { return false }
func (s NonePipeFitting) getNext(string, *PipeMap) (*Fitting, string) { return nil, "" }

type VertPipeFitting struct {
	BaseFitting
}

func (s VertPipeFitting) opensBottom() bool { return true }
func (s VertPipeFitting) opensRight() bool  { return false }
func (s VertPipeFitting) opensLeft() bool   { return false }
func (s VertPipeFitting) opensTop() bool    { return true }
func (s VertPipeFitting) getNext(cellEntry string, pipeMap *PipeMap) (*Fitting, string) {
	if !slices.Contains([]string{"top", "bottom"}, cellEntry) {
		panic(fmt.Sprintf("Invalid entry direction %s", cellEntry))
	}
	if cellEntry == "top" {
		return s.getBottom(pipeMap), "top"
	}
	return s.getTop(pipeMap), "bottom"
}

type HorizontalPipeFitting struct {
	BaseFitting
}

func (s HorizontalPipeFitting) opensBottom() bool { return false }
func (s HorizontalPipeFitting) opensRight() bool  { return true }
func (s HorizontalPipeFitting) opensLeft() bool   { return true }
func (s HorizontalPipeFitting) opensTop() bool    { return false }
func (s HorizontalPipeFitting) getNext(cellEntry string, pipeMap *PipeMap) (*Fitting, string) {
	if !slices.Contains([]string{"left", "right"}, cellEntry) {
		panic(fmt.Sprintf("Invalid entry direction %s", cellEntry))
	}
	if cellEntry == "left" {
		return s.getRight(pipeMap), "left"
	}
	return s.getLeft(pipeMap), "right"
}

type LPipeFitting struct {
	BaseFitting
}

func (s LPipeFitting) opensBottom() bool { return false }
func (s LPipeFitting) opensRight() bool  { return true }
func (s LPipeFitting) opensLeft() bool   { return false }
func (s LPipeFitting) opensTop() bool    { return true }

func (s LPipeFitting) getNext(cellEntry string, pipeMap *PipeMap) (*Fitting, string) {
	if !slices.Contains([]string{"right", "top"}, cellEntry) {
		panic(fmt.Sprintf("Invalid entry direction %s", cellEntry))
	}
	if cellEntry == "right" {
		return s.getTop(pipeMap), "bottom"
	}
	return s.getRight(pipeMap), "left"
}

type JPipeFitting struct {
	BaseFitting
}

func (s JPipeFitting) opensBottom() bool { return false }
func (s JPipeFitting) opensRight() bool  { return false }
func (s JPipeFitting) opensLeft() bool   { return true }
func (s JPipeFitting) opensTop() bool    { return true }
func (s JPipeFitting) getNext(cellEntry string, pipeMap *PipeMap) (*Fitting, string) {
	if !slices.Contains([]string{"left", "top"}, cellEntry) {
		panic(fmt.Sprintf("Invalid entry direction %s", cellEntry))
	}
	if cellEntry == "top" {
		return s.getLeft(pipeMap), "right"
	}
	return s.getTop(pipeMap), "bottom"
}

type FPipeFitting struct {
	BaseFitting
}

func (s FPipeFitting) opensBottom() bool { return true }
func (s FPipeFitting) opensRight() bool  { return true }
func (s FPipeFitting) opensLeft() bool   { return false }
func (s FPipeFitting) opensTop() bool    { return false }
func (s FPipeFitting) getNext(cellEntry string, pipeMap *PipeMap) (*Fitting, string) {
	if !slices.Contains([]string{"bottom", "right"}, cellEntry) {
		panic(fmt.Sprintf("Invalid entry direction %s", cellEntry))
	}
	if cellEntry == "bottom" {
		return s.getRight(pipeMap), "left"
	}
	return s.getBottom(pipeMap), "top"
}

type SevenPipeFitting struct {
	BaseFitting
}

func (s SevenPipeFitting) opensBottom() bool { return true }
func (s SevenPipeFitting) opensRight() bool  { return false }
func (s SevenPipeFitting) opensLeft() bool   { return true }
func (s SevenPipeFitting) opensTop() bool    { return false }
func (s SevenPipeFitting) getNext(cellEntry string, pipeMap *PipeMap) (*Fitting, string) {
	if !slices.Contains([]string{"left", "bottom"}, cellEntry) {
		panic(fmt.Sprintf("Invalid entry direction %s", cellEntry))
	}
	if cellEntry == "left" {
		return s.getBottom(pipeMap), "top"
	}
	return s.getLeft(pipeMap), "right"
}

type BaseFitting struct {
	Pos lexer.Position
}
type AdjCells struct {
	left   *Fitting
	right  *Fitting
	top    *Fitting
	bottom *Fitting
}

func (b BaseFitting) getAdjacentFittings(s *PipeMap) AdjCells {
	var (
		left   *Fitting = nil
		right  *Fitting = nil
		top    *Fitting = nil
		bottom *Fitting = nil
	)

	left = b.getLeft(s)
	right = b.getRight(s)
	top = b.getTop(s)
	bottom = b.getBottom(s)
	return AdjCells{
		left,
		right,
		top,
		bottom,
	}
}

func (b BaseFitting) row() int {
	return b.Pos.Line - 1
}

func (b BaseFitting) column() int {
	return b.Pos.Column - 1
}
func (b BaseFitting) getLeft(s *PipeMap) *Fitting {
	log.Println("Moving left")
	var row = b.row()
	var column = b.column()
	var line = s.Line[row]
	if column > 0 {
		return &line.Fitting[column-1]
	}
	return nil
}
func (b BaseFitting) getRight(s *PipeMap) *Fitting {
	log.Println("Moving right")
	var row = b.row()
	var column = b.column()
	var totalColumns = len(s.Line[0].Fitting)
	if column < totalColumns {
		return &s.Line[row].Fitting[column+1]
	}
	return nil
}
func (b BaseFitting) getBottom(s *PipeMap) *Fitting {
	log.Println("Moving bottom")
	var row = b.row()
	var column = b.column()
	var totalLines = len(s.Line)
	if row < totalLines {
		return &s.Line[row+1].Fitting[column]
	}
	return nil
}
func (b BaseFitting) getTop(s *PipeMap) *Fitting {

	log.Println("Moving top")
	var row = b.row()
	var column = b.column()
	if row > 0 {
		return &s.Line[row-1].Fitting[column]
	}
	return nil
}
