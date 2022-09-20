package img2txt

import "errors"

var (
	errPathIsEmpty = errors.New("Img2txt.Path can not be ''")
	errTxtIsEmpty  = errors.New("Img2txt.Txt can not be ''")
	// Character set printed by default
	txt = []string{"W", "@", "#", "8", "&", "*", "o", ":", ".", " "}
)
