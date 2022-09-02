// Lissajous generates GIF animations of random Lissajous figures.

package main

import (
    "io"
    "os"
    "image"
    "image/color"
    "image/gif"
    "math"
    "math/rand"
)

var pallete = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // first color in pallete
    blackIndex = 1 // next color in pallete
)

func main() {
    // TODO: Exercise 1.5
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        // number of complete oscillator revolutions
        cycles  = 10 // 5
        res     = 0.001 // angular resolution
        size    = 150 // image canvas covers (100)
        // [-size...+size]
        nframes = 64 // no. of animation frames
        delay   = 8 // delay b/w frames in 10ms units
    )
    // relative frequency of y oscillator
    freq := rand.Float64() * 3 
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.5 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, pallete)
        for t := 0.0; t<cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(
                size + int(x*size + 0.5),
                size + int(y*size + 0.5),
                blackIndex,   
            )
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
    // NOTE: ignoring encoding errors
}
