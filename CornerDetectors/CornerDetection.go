package CornerDetectors

import (
	"fmt"
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"image"

)

func GetSIFTKeyPoints(img gocv.Mat) {
	si := contrib.NewSIFT()
	kp := si.Detect(img)
	for i := 0 ; i < len(kp) ; i++ {
		fmt.Println(kp[i].Y)
	}
}

func GetFastKeyPoints(img gocv.Mat) []gocv.KeyPoint {
	ffd := gocv.NewFastFeatureDetector()
	defer ffd.Close()
	kp := ffd.Detect(img)
	return kp
}


func GetPatchAreasFromImage(img gocv.Mat , kp []gocv.KeyPoint){
	for i:=0 ; i < len(kp) ; i++ {
		// crop out a rectangular regio
		roiArea := image.Rect((int)(kp[0].X - 7  ), (int)(kp[0].Y - 7  ), (int)(kp[0].X + 7  ), (int)(kp[0].Y + 7  ))
		croppedArea := img.Region(roiArea)
		window := gocv.NewWindow("nabers")
		window.IMShow(croppedArea)
		window.WaitKey(0)


	}

}