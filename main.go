package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

//go:embed source/HarmonyOS_Sans_SC_Regular.ttf
var customFont []byte

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(_ fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "newFont",
		StaticContent: customFont,
	}
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

//go:embed source/Genshin_Chinese_logo.svg
var genshinLogo []byte
var resourceGenshinChineselogoSvg = fyne.NewStaticResource("Genshin_Chinese_logo.svg", genshinLogo)

func main() {
	//fmt.Printf("%v", genshinLogo)
	myApp := app.New()
	myWindow := myApp.NewWindow("原神启动！")
	myWindow.SetPadded(false)
	myWindow.Resize(fyne.NewSize(800, 600))
	// 从文件读取 SVG 数据
	img := canvas.NewImageFromFile("source/Genshin_Chinese_logo.svg")
	img.FillMode = canvas.ImageFillOriginal
	fmt.Printf("%+v\n", img)
	myWindow.SetContent(img)
	myWindow.ShowAndRun()
}

/*
	img := canvas.NewImageFromResource(resourceGenshinChineselogoSvg)
	println("\n123\n")
	img.FillMode = canvas.ImageFillOriginal

	//// 创建一个黑色的矩形
	rect := canvas.NewRectangle(color.Black)

	// 创建一个白色容器
	bg := canvas.NewRectangle(color.White)
	content := container.NewMax(bg, img)
	println("\ncontent:\n")
	fmt.Printf("%+v\n", content)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	// 渐进效果
	go func() {
		for i := 255; i >= 0; i-- {
			rect.FillColor = color.NRGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
			rect.Refresh()
			time.Sleep(time.Millisecond * 10)
		}
	}()
	println("\n123\n")

}
*/
