package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"net/http"

	"github.com/nfnt/resize"
	"github.com/vova616/screenshot"
)

type RGBA struct {
	R, G, B, A uint8
}

func img2json(img image.Image) []byte {
	myImg := resize.Resize(128, 72, img, resize.Lanczos3)
	var imgjson [128][72]color.RGBA
	for x := 0; x < 128; x++ {
		var xa [72]color.RGBA
		for y := 0; y < 72; y++ {
			xa[y] = myImg.At(x, y).(color.RGBA)
		}
		imgjson[x] = xa
	}
	jsn, err := json.Marshal(imgjson)
	if err != nil {
		panic(err)
	}
	return jsn
}

func jsonimgsend(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/json" {
		http.Error(resp, "404 not found.", http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		http.Error(resp, "404 not found.", http.StatusNotFound)
		return
	}
	fmt.Println("JSON image requested")
	fmt.Fprintf(resp, string(img2json(scrn())))
}

func scrn() image.Image {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	return resize.Resize(128, 72, image.Image(img), resize.Lanczos3)
}

func main() {
	http.HandleFunc("/json", jsonimgsend)

	fmt.Printf("Opening port at 80")
	http.ListenAndServe(":80", nil)
}
