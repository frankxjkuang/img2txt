package img2txt

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

// img2txt struct
type img2txt struct {
	// img path
	path string
	// step of scanning image points, default is 2
	step int
	// y scaling of x axis, default is y=3x
	scale int
	// printed character set, default is txt
	txt []string
}

// NewImg2txt get pointer of img2txt struct
func NewImg2txt() *img2txt {
	return &img2txt{"", 2, 3, txt}
}

// SetPath set the path of img2txt
func (i *img2txt) SetPath(path string) *img2txt {
	if path == "" {
		panic("path cannot be ''")
	}
	i.path = path
	return i
}

// SetStep set the step of img2txt
func (i *img2txt) SetStep(step int) *img2txt {
	if step <= 0 {
		panic("step cannot be less than or equal to 0")
	}
	i.step = step
	return i
}

// SetScale set scale of img2txt
func (i *img2txt) SetScale(scale int) *img2txt {
	if scale <= 0 {
		panic("step cannot be less than or equal to 0")
	}
	i.scale = scale
	return i
}

// SetTxt set txt of img2txt
func (i *img2txt) SetTxt(txt []string) *img2txt {
	i.txt = txt
	return i
}

// PrintTxt print the string after converting the picture with txt
func (i *img2txt) PrintTxt() (err error) {
	txt, err := i.GetTxt()
	if err != nil {
		return err
	}
	fmt.Print(txt)
	return nil
}

// GetTxt get the string after converting the picture with txt
func (i *img2txt) GetTxt() (string, error) {
	if i.path == "" {
		return "", errPathIsEmpty
	}

	if len(i.txt) == 0 {
		return "", errTxtIsEmpty
	}

	imgfile, err := os.Open(i.path)
	if err != nil {
		return "", err
	}
	defer imgfile.Close()

	var img image.Image
	if strings.Contains(i.path, "gif") {
		img, err = gif.Decode(imgfile)
	} else {
		img, _, err = image.Decode(imgfile)
	}
	if err != nil {
		return "", err
	}

	txtB := strings.Builder{}

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += i.step * i.scale {
		for x := bounds.Min.X; x < bounds.Max.X; x += i.step {
			txtB.WriteString(i.getImgGrayByPoint(img, x, y))
		}
		txtB.WriteString("\n")
	}

	return txtB.String(), nil
}

// getImgGrayByPoint get the txt string corresponding to the gray value of a point in the picture
func (i *img2txt) getImgGrayByPoint(img image.Image, x, y int) string {
	r, g, b, a := img.At(x, y).RGBA()

	// Formula Reference：https://www.kdnuggets.com/2019/12/convert-rgb-image-grayscale.html
	// idx / len(txt) = (0.2126*r + 0.7152*g + 0.0722*b) / a
	v := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
	if a == 0 {
		return " "
	}

	idx := len(i.txt) * int(v) / int(a)

	if idx >= len(i.txt) {
		idx = len(i.txt) - 1
	}
	return i.txt[idx]
}
