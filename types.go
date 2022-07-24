package main

import "image"

type ObjectDetected struct {
	Id   int
	Rect image.Rectangle
	Byte []byte
}
