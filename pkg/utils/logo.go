package utils

import (
	"fmt"
	"goapp/pkg/consts"

	"github.com/mbndr/figlet4go"
)

// Logotype output
func Logo() {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	// options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
	}
	renderStr, _ := ascii.RenderOpts("GOAPP", options)
	fmt.Print(renderStr)
	fmt.Printf("\t%sCreate %sGOLANG%s application%s\n\n", consts.Green, consts.Blue, consts.Yellow, consts.Reset)
}
