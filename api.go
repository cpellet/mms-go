package mms_go

import "fmt"

// --- API Interface ---

type MMS struct{}

func NewMMS() *MMS {
	return &MMS{}
}

// --- Command Helpers ---

func (m *MMS) getIntCmd(cmd string) int {
	fmt.Println(cmd)
	var i int
	_, err := fmt.Scanln(&i)
	if err != nil {
		panic(err)
	}
	return i
}

func (m *MMS) getBoolCmd(cmd string) bool {
	fmt.Println(cmd)
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
	}
	return s == "true"
}

func (m *MMS) getFloatCmd(cmd string) float64 {
	fmt.Println(cmd)
	var f float64
	_, err := fmt.Scanln(&f)
	if err != nil {
		panic(err)
	}
	return f
}

func (m *MMS) ackCmd(cmd string) bool {
	fmt.Println(cmd)
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
	}
	return s == "ack"
}

// --- API Commands ---

func (m *MMS) MazeWidth() int {
	return m.getIntCmd("mazeWidth")
}

func (m *MMS) MazeHeight() int {
	return m.getIntCmd("mazeHeight")
}

func (m *MMS) WallFront() bool {
	return m.getBoolCmd("wallFront")
}

func (m *MMS) WallRight() bool {
	return m.getBoolCmd("wallRight")
}

func (m *MMS) WallLeft() bool {
	return m.getBoolCmd("wallLeft")
}

func (m *MMS) MoveForward(steps int) bool {
	return m.ackCmd(fmt.Sprintf("moveForward"))
}

func (m *MMS) MoveForwardHalf(steps int) bool {
	return m.ackCmd(fmt.Sprintf("moveForwardHalf %d", steps))
}

func (m *MMS) TurnRight() bool {
	return m.ackCmd("turnRight")
}

func (m *MMS) TurnLeft() bool {
	return m.ackCmd("turnLeft")
}

func (m *MMS) TurnRight45() bool {
	return m.ackCmd("turnRight45")
}

func (m *MMS) TurnLeft45() bool {
	return m.ackCmd("turnLeft45")
}

func (m *MMS) SetWall(x, y int, dir string) {
	fmt.Println(fmt.Sprintf("setWall %d %d %s", x, y, dir))
}

func (m *MMS) ClearWall(x, y int, dir string) {
	fmt.Println(fmt.Sprintf("clearWall %d %d %s", x, y, dir))
}

func (m *MMS) SetColor(x, y int, color string) {
	fmt.Println(fmt.Sprintf("setColor %d %d %s", x, y, color))
}

func (m *MMS) ClearColor(x, y int) {
	fmt.Println(fmt.Sprintf("clearColor %d %d", x, y))
}

func (m *MMS) ClearAllColor() {
	fmt.Println("clearAllColor")
}

func (m *MMS) SetText(x, y int, text string) {
	fmt.Println(fmt.Sprintf("setText %d %d %s", x, y, text))
}

func (m *MMS) ClearText(x, y int) {
	fmt.Println(fmt.Sprintf("clearText %d %d", x, y))
}

func (m *MMS) ClearAllText() {
	fmt.Println("clearAllText")
}

func (m *MMS) WasReset() bool {
	return m.getBoolCmd("wasReset")
}

func (m *MMS) AckReset() bool {
	return m.ackCmd("ackReset")
}

func (m *MMS) GetStat(stat string) float64 {
	return m.getFloatCmd(fmt.Sprintf("getStat %s", stat))
}
