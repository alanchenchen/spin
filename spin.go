package spin

import (
	"fmt"
	"os"
	"time"
)

const lockThreadsNum uint = 1

// spin character
type character struct {
	frame    []string
	interval uint
}

// spin content
type Spin struct {
	shouldWork     bool
	lockThreadsNum uint
	character      character
	ch             chan bool
}

func (s *Spin) run() {
	for range s.ch {
		for _, r := range s.character.frame {
			fmt.Printf("\r%v", r)
			time.Sleep(time.Duration(s.character.interval) * time.Millisecond)
		}
	}
}

func (s *Spin) waitForRun() {
	for {
		if s.shouldWork {
			s.ch <- true
		} else {
			break
		}
	}
	close(s.ch)
}

// start spinning
func (s *Spin) Start() {
	if s.lockThreadsNum == 1 {
		go s.waitForRun()
		go s.run()
	} else {
		fmt.Println("current spin locks only one goroutine, so don't call start method multiple times")
		os.Exit(1)
	}
	s.lockThreadsNum++
}

// stop spinning
func (s *Spin) Stop() {
	s.shouldWork = false
	fmt.Print("\r")
}

// stop spnning, echo succeessful tip
func (s *Spin) Succeed(text string) {
	s.Stop()
	fmt.Printf("%v %v\n", newSymbol(green, successSymbol), text)
}

// stop spnning, echo warning tip
func (s *Spin) Warn(text string) {
	s.Stop()
	fmt.Printf("%v %v\n", newSymbol(yellow, warningSymbol), text)
}

// stop spnning, echo info tip
func (s *Spin) Info(text string) {
	s.Stop()
	fmt.Printf("%v %v\n", newSymbol(blue, infoSymbol), text)
}

// stop spnning, echo error tip
func (s *Spin) Fail(text string) {
	s.Stop()
	fmt.Printf("%v %v\n", newSymbol(red, errorSymbol), text)
}

// struct a spin instance
func New(spinType character) *Spin {
	s := new(Spin)
	s.shouldWork = true
	s.lockThreadsNum = lockThreadsNum
	s.character = spinType
	s.ch = make(chan bool)
	return s
}

// custom spin character
func NewCharacter(frame []string, interval uint) character {
	return character{frame, interval}
}
