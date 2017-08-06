package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"log"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.RGBA{255, 255, 255, 1}, color.RGBA{255, 0, 0, 1}, color.RGBA{255, 255, 0, 1}, color.RGBA{255, 255, 255, 1}, 
color.RGBA{0, 0, 255, 1}, color.RGBA{255, 0, 255, 1}, color.RGBA{0, 255, 255, 1}, color.RGBA{0, 255, 0, 1}}

var (
	cycles = 5
	res = 0.001
	size = 100
	nframes = 64
	delay = 8
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("11.11.1.10:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if r.Form["cycles"] != nil {
		cycles, _ = strconv.Atoi(r.Form["cycles"][0])
	}
	if r.Form["res"] != nil {
		res, _ = strconv.ParseFloat(r.Form["res"][0], 64)
	}
	if r.Form["size"] != nil {
		size, _ = strconv.Atoi(r.Form["size"][0])
	}
	if r.Form["nframes"] != nil {
		nframes, _ = strconv.Atoi(r.Form["nframes"][0])
	}
	if r.Form["delay"] != nil {
		delay, _ = strconv.Atoi(r.Form["delay"][0])
	}
	lissajous(w)
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		index := uint8(rand.Int() % len(palette))
		n := float64(cycles)*2*math.Pi
		for t := 0.0; t < n; t+= res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
