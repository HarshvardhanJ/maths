// Rose generates GIF animations of random rose petals figures.

package main

import (
    "io"
    "os"
    "log"
    "time"
    "image"
    "image/color"
    "image/gif"
    "math"
    "math/rand"
    "net/http"
)

var pallete = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // first color in pallete
    blackIndex = 1 // next color in pallete
)

func main() {
    // seeding randomness
    rand.Seed(time.Now().UTC().UnixNano())
    // rendering figure on web-page
    if len(os.Args) > 1 && os.Args[1] == "web" {
        // Usage: ./rose web
        http.HandleFunc("/", handler)
        log.Fatal(
          http.ListenAndServe("localhost:8000", nil),
        ) // => localhost:8000
        return
    }
    rose(os.Stdout)
}

func rose(out io.Writer) {
    const (
        // number of complete oscillator revolutions
        cycles  = 10
        res     = 0.002 // angular resolution (0.001)
        size    = 100 // image canvas covers
        // [-size...+size]
        nframes = 70 // no. of animation frames
        delay   = 8 // delay b/w frames in 10ms units
    )
    // relative frequency of y oscillator
    freq := rand.Float64() * 3 
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.5 // phase difference
    a := 0.5
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, pallete)
        a += 0.5
        for t := 0.0; t<cycles*2*math.Pi; t += res {
            // equations:
            x := a*math.Cos(t*freq) * math.Cos(t)
            y := a*math.Cos(t*freq) * math.Sin(t)

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

func handler(w http.ResponseWriter, r *http.Request){
    rose(w)
}
