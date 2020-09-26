package main

import "go.uber.org/zap"

func main() {
	mobileNo := "0631846345"
	l := zap.NewExample()
	l = l.With(zap.Namespace("hometic"), zap.String("mobileNo", mobileNo))
	l.Info("Hello")
}
