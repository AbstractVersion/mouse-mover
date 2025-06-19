package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	ShaftLength   = 200
	HeadRadius    = 50
	BallRadius    = 30
	BallOffset    = 40
	MovementDelay = 20 * time.Millisecond
)

// MouseMover handles the mouse movement patterns
type MouseMover struct {
	screenWidth  int
	screenHeight int
	startX       int
	startY       int
}

// NewMouseMover creates a new MouseMover instance
func NewMouseMover() *MouseMover {
	width, height := robotgo.GetScreenSize()
	return &MouseMover{
		screenWidth:  width,
		screenHeight: height,
		startX:       width / 2,
		startY:       height / 2,
	}
}

// moveToPosition moves the mouse to a specific position with animation
func (m *MouseMover) moveToPosition(x, y int) {
	robotgo.MoveSmooth(x, y, 0.5, 0.5)
	time.Sleep(MovementDelay)
}

// drawShaft draws the vertical shaft of the pattern
func (m *MouseMover) drawShaft() {
	fmt.Println("Drawing shaft...")
	for i := 0; i < ShaftLength; i++ {
		m.moveToPosition(m.startX, m.startY+i)
	}
}

// drawHead draws the semicircular head
func (m *MouseMover) drawHead() {
	fmt.Println("Drawing head...")
	headCenterY := m.startY + ShaftLength

	for angle := 0; angle <= 180; angle += 5 {
		radians := float64(angle) * math.Pi / 180
		x := m.startX + int(float64(HeadRadius)*math.Cos(radians))
		y := headCenterY + int(float64(HeadRadius)*math.Sin(radians))
		m.moveToPosition(x, y)
	}
}

// drawBalls draws the two circular elements at the base
func (m *MouseMover) drawBalls() {
	fmt.Println("Drawing left ball...")
	// Left ball
	leftCenterX := m.startX - BallOffset
	leftCenterY := m.startY - 20

	for angle := 0; angle <= 360; angle += 10 {
		radians := float64(angle) * math.Pi / 180
		x := leftCenterX + int(float64(BallRadius)*math.Cos(radians))
		y := leftCenterY + int(float64(BallRadius)*math.Sin(radians))
		m.moveToPosition(x, y)
	}

	fmt.Println("Drawing right ball...")
	// Right ball
	rightCenterX := m.startX + BallOffset
	rightCenterY := m.startY - 20

	for angle := 0; angle <= 360; angle += 10 {
		radians := float64(angle) * math.Pi / 180
		x := rightCenterX + int(float64(BallRadius)*math.Cos(radians))
		y := rightCenterY + int(float64(BallRadius)*math.Sin(radians))
		m.moveToPosition(x, y)
	}
}

// ExecutePattern executes the complete movement pattern
func (m *MouseMover) ExecutePattern() {
	fmt.Printf("Starting pattern at center: (%d, %d)\n", m.startX, m.startY)

	// Move to starting position
	m.moveToPosition(m.startX, m.startY)
	time.Sleep(500 * time.Millisecond)

	// Execute the pattern
	m.drawShaft()
	m.drawHead()
	m.drawBalls()

	// Return to center
	fmt.Println("Returning to center...")
	m.moveToPosition(m.startX, m.startY)

	fmt.Println("Pattern completed!")
}

// KeepActive runs the pattern at regular intervals
func (m *MouseMover) KeepActive(intervalMinutes int) {
	fmt.Printf("Starting mouse activity simulation...\n")
	fmt.Printf("Mouse will move every %d minutes\n", intervalMinutes)
	fmt.Println("Press Ctrl+C to stop")

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
	defer ticker.Stop()

	// Execute pattern immediately on start
	fmt.Printf("Executing pattern at %s\n", time.Now().Format("15:04:05"))
	m.ExecutePattern()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Executing pattern at %s\n", time.Now().Format("15:04:05"))
			m.ExecutePattern()
		case sig := <-sigChan:
			fmt.Printf("\nReceived signal: %v\n", sig)
			fmt.Println("Stopping mouse movement script...")
			return
		}
	}
}

// printSystemInfo prints system information
func printSystemInfo() {
	fmt.Println("=== Mouse Movement Script ===")
	width, height := robotgo.GetScreenSize()
	fmt.Printf("Screen size: %dx%d\n", width, height)

	// Get current mouse position
	x, y := robotgo.GetMousePos()
	fmt.Printf("Current mouse position: (%d, %d)\n", x, y)
	fmt.Println("==============================")
}

func main() {
	printSystemInfo()

	// Create mouse mover instance
	mover := NewMouseMover()

	// Check command line arguments
	if len(os.Args) > 1 && os.Args[1] == "once" {
		// Run pattern once
		fmt.Println("Running pattern once...")
		mover.ExecutePattern()
	} else {
		// Run continuously (default: every 3 minutes)
		intervalMinutes := 3
		mover.KeepActive(intervalMinutes)
	}
}
