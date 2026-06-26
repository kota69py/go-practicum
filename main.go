package main

import (
	"github.com/kota69py/go-practicum/cmd"
	"github.com/kota69py/go-practicum/internal/exercdata"
)

func main() {
	r := cmd.NewRunner(exercdata.FS())
	r.Execute()
}
