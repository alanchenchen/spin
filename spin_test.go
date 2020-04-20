package spin

import "time"

func ExampleNew() {
	spinner := spin.New(spin.Dot)
	spinner.Start()
	time.Sleep(20 * time.Second)
	spinner.Stop()
	spinner.Succeed("finish")
	spinner.Info("finish")
	spinner.Warn("finish")
	spinner.Fail("finish")
	// Output:
	// Start: spinner starts to spin
	// Stops: spinner stops and quit spinner goroutine
	// Succeed: spinner stops with successful icon and quit spinner goroutine
	// Info: spinner stops with info icon and quit spinner goroutine
	// Warn: spinner stops with warning icon and quit spinner goroutine
	// Fail: spinner stops with error icon and quit spinner goroutine
}

func ExampleNewCharacter() {
	customFrame := spin.NewCharacter([]string{
		"-",
		"--",
		"---",
		"----",
		"---",
		"--",
		"-",
	}, 100)
	// Output:
	// an new spin character frame
}
