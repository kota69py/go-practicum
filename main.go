package main

import (
	"github.com/kota69py/go-practicum/cmd"
	"github.com/kota69py/go-practicum/internal/exercdata"
)

func main() {
	cmd.SetExercFS(exercdata.FS())
	cmd.Execute()
}
