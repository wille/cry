package common

// Extensions to encrypt
var Extensions = [...]string{
	"png",
	"jpg",
	"jpeg",
	"tif",
}

// IgnoreNames will skip files and directories that contains the string
var IgnoreNames = [...]string{
	"AppData",
}

const (
	// LockedExtension to append to file name when encrypted
	LockedExtension string = ".locked"

	// ProcessMax X files, then stop
	ProcessMax int = 10
)
