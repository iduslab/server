package utils

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

func DownloadImageViaURL(url, filePath string) error {
	response, e := http.Get(url)
	if e != nil {
		return e
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func Hex2RGB(hex string) (*color.RGBA, error) {
	values, err := strconv.ParseUint(hex, 16, 32)

	if err != nil {
		return nil, err
	}

	return &color.RGBA{uint8(values >> 16), uint8((values >> 8) & 0xFF), uint8(values & 0xFF), 0}, nil
}

// WaterMark for adding a watermark on the image
func WaterMark(imgReader io.Reader, markText string, color *color.RGBA) (image.Image, error) {
	img, _, err := image.Decode(imgReader)
	if err != nil {
		return nil, err
	}

	// image's length to canvas's length
	bounds := img.Bounds()
	w := vg.Length(bounds.Max.X) * vg.Inch / vgimg.DefaultDPI
	h := vg.Length(bounds.Max.Y) * vg.Inch / vgimg.DefaultDPI
	diagonal := vg.Length(math.Sqrt(float64(w*w + h*h)))

	// create a canvas, which width and height are diagonal
	c := vgimg.New(diagonal, diagonal)

	// draw image on the center of canvas
	rect := vg.Rectangle{}
	rect.Min.X = diagonal/2 - w/2
	rect.Min.Y = diagonal/2 - h/2
	rect.Max.X = diagonal/2 + w/2
	rect.Max.Y = diagonal/2 + h/2
	c.DrawImage(rect, img)

	// make a fontStyle, which width is vg.Inch * 0.7
	fontStyle, _ := vg.MakeFont("Courier", vg.Inch*0.7)

	// repeat the markText
	markTextWidth := fontStyle.Width(markText)
	unitText := markText
	for markTextWidth <= diagonal {
		markText += " " + unitText + " "
		markTextWidth = fontStyle.Width(markText)
	}

	// set the color of markText
	c.SetColor(color)

	// set a random angle between 0 and π/2
	c.Rotate(0.3)

	// set the lineHeight and add the markText
	lineHeight := fontStyle.Extents().Height * 2.7
	for offset := -2 * diagonal; offset < 2*diagonal; offset += lineHeight {
		c.FillString(fontStyle, vg.Point{X: 0, Y: offset}, markText)
	}

	// canvas writeto jpeg
	// canvas.img is private
	// so use a buffer to transfer
	jc := vgimg.PngCanvas{Canvas: c}
	buff := new(bytes.Buffer)
	jc.WriteTo(buff)
	img2, _, err := image.Decode(buff)
	if err != nil {
		return nil, err
	}

	// get the center point of the image
	ctp := int(diagonal * vgimg.DefaultDPI / vg.Inch / 2)

	// cutout the marked image
	size := bounds.Size()
	bounds = image.Rect(ctp-size.X/2, ctp-size.Y/2, ctp+size.X/2, ctp+size.Y/2)
	rv := image.NewRGBA(bounds)
	draw.Draw(rv, bounds, img2, bounds.Min, draw.Src)
	return rv, nil
}

func ImageEncode(img image.Image, ext string) (rv *bytes.Buffer, err error) {
	ext = strings.ToLower(ext)
	rv = new(bytes.Buffer)
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(rv, img, &jpeg.Options{Quality: 100})
	case ".png":
		err = png.Encode(rv, img)
	}
	return rv, err
}
