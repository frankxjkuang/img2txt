package img2txt

import (
	"testing"
)

func TestImg2txt_PrintTxt_png(t *testing.T) {
	img2txt := NewImg2txt()
	err := img2txt.
		SetPath("./look.png").
		SetStep(1).
		SetScale(3).
		PrintTxt()
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestImg2txt_PrintTxt_jpeg(t *testing.T) {
	img2txt := NewImg2txt()
	err := img2txt.
		SetPath("./timg.jpeg").
		SetStep(4).
		SetScale(3).
		PrintTxt()
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestImg2txt_PrintTxt_jpg(t *testing.T) {
	img2txt := NewImg2txt()
	err := img2txt.
		SetPath("./modi.jpg").
		SetStep(6).
		SetScale(3).
		SetTxt([]string{" ", ".", ":", "o", "*", "&", "8", "#", "@", "W"}).
		PrintTxt()
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestImg2txt_PrintTxt_gif(t *testing.T) {
	img2txt := NewImg2txt()
	err := img2txt.
		SetPath("./lcsm.gif").
		SetStep(1).
		SetScale(3).
		PrintTxt()
	if err != nil {
		t.Fatal("err:", err)
	}
}

func TestImg2txt_GetTxt(t *testing.T) {
	img2txt := NewImg2txt()
	txt, err := img2txt.
		SetPath("./look.png").
		SetStep(2).
		SetScale(3).
		GetTxt()
	if err != nil {
		t.Logf("err: %v", err)
	} else {
		t.Log(txt)
	}
}
