package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

func main() {
	//fmt.Printf("%v", genshinLogo)
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{})
	myWindow := myApp.NewWindow("原神，启动！")
	myWindow.Resize(fyne.NewSize(365*4, 280*4))
	svgRes, err := fyne.LoadResourceFromPath("source/Genshin_Chinese_logo.svg")
	if err != nil {
		fmt.Print(err)
	}
	img := canvas.NewImageFromResource(svgRes)

	// panic: img.FillMode = canvas.ImageFillOriginal // 保持原始大小
	//imgSize := fyne.NewSize(365, 280)
	//img.FillMode = canvas.ImageFillContain
	//img.FillMode = canvas.ImageFillOriginal
	// 创建一个黑色的矩形
	rect := canvas.NewRectangle(color.Black)
	// 创建一个白色容器
	//bg := canvas.NewRectangle(color.White)
	//println(bg)
	//img.Resize(fyne.NewSize(365, 280))

	// 创建居中的图像，并添加边框
	label := container.NewCenter(widget.NewLabel("mihuyo澳大利亚科技有限公司"))
	//box1 := layout.NewHBoxLayout()
	//var spacer = layout.NewSpacer()
	//box2 := layout.NewVBoxLayout()

	content2 := container.NewMax(img)
	layout3 := layout.NewGridWrapLayout(fyne.NewSize(365, 280))
	content3 := container.New(layout3, content2)
	content4 := container.NewCenter(content3)
	content5 := container.NewBorder(nil, label, nil, nil, content4)
	content6 := container.NewMax(rect, content5)
	go func() {
		for i := 0; i <= 255; i++ {
			if i == 0 {
				time.Sleep(time.Second * 1)
			}
			rect.FillColor = color.NRGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
			rect.Refresh()
			time.Sleep(time.Millisecond * 20)
		}
	}()
	//
	myWindow.SetContent(content6)
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
