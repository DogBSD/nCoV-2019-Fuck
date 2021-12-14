package controller

import (
	"github.com/pterm/pterm"
)

func PrintLogo() {
	pterm.Info.Println("Everyone is responsible for fighting the epidemic.")
	pterm.Info.Println("Please reduce your activities to crowded places.")
	pterm.Print("\n\n")

	ss, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("FUCK")).Srender()
	pterm.DefaultCenter.Printfln(ss)

	// err := pterm.DefaultBigText.WithLetters(
	// 	pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgCyan)),
	// 	pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).Render()

	// if err != nil {
	// 	return
	// }
	pterm.DefaultCenter.WithCenterEachLineSeparately().Printfln("抗击疫情，人人有责\n请减少去往人员密集点")
}
