package exercdata

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed exercdata
var exercFS embed.FS

func FS() fs.FS {
	sub, err := fs.Sub(exercFS, "exercdata")
	if err != nil {
		log.Fatalf("exercdata.FS: %v", err)
	}
	return sub
}
