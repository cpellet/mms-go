package main

import "github.com/cpellet/mms-go"

func main() {
	mms := mms_go.NewMMS()
	mms.SetColor(0, 0, "G")
	mms.SetText(0, 0, "abc")
	for {
		if !mms.WallLeft() {
			mms.TurnLeft()
		}
		for mms.WallFront() {
			mms.TurnRight()
		}
		mms.MoveForward(1)
	}
}
