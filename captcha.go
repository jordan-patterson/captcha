package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/pbnjay/pixfont"
)

type Captcha struct {
	//random text to be displayed as captcha
	Text string
}

func New() *Captcha {
	var text string
	rand.Seed(time.Now().Unix())
	for i := 0; i < 6; {
		switch rand.Intn(2) {
		case 0:
			text += string(rand.Intn(9) + 48)
		case 1:
			text += string(rand.Intn(25) + 65)
		case 2:
			text += string(rand.Intn(25) + 98)
		}
		i++
	}
	captcha := &Captcha{Text: text}
	return captcha
}
func (c *Captcha) WriteImage(w http.ResponseWriter) error {
	/*
		Writes captcha image to response writer.

		You will have to
		implement your http handler function
		to generate a new captcha on every page request
		and maintain the captcha text across a session
	*/
	buffer := new(bytes.Buffer)
	img := image.NewRGBA(image.Rect(0, 0, 150, 50))
	pixfont.DrawString(img, 12, 12, c.Text, color.Black)
	if err := png.Encode(buffer, img); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		return err
	}
	return nil
}
