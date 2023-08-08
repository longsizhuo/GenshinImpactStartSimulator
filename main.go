package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"time"
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

//go:embed source/1.svg
var _ []byte

func main() {
	//fmt.Printf("%v", genshinLogo)
	myApp := app.New()
	myWindow := myApp.NewWindow("原神，启动！")
	myWindow.Resize(fyne.NewSize(365, 280))
	img := canvas.NewImageFromFile("source/Genshin_Chinese_logo.svg")
	// TODO:bug: 无法正常显示
	//img.FillMode = canvas.ImageFillOriginal
	rect := canvas.NewRectangle(color.Black)
	// 创建一个白色容器
	bg := canvas.NewRectangle(color.White)
	content := container.NewMax(bg, img)
	go func() {
		for i := 255; i >= 0; i-- {
			rect.FillColor = color.NRGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
			rect.Refresh()
			time.Sleep(time.Millisecond * 10)
		}
	}()
	myWindow.SetContent(content)
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
