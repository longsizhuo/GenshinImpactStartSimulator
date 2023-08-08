1. 在macOS上运行失败，fyne还不支持在macOS上运行SVG，要么去做开发版本，要么去Windows上解决
2. 在Windows上运行失败，原因是缺少OpenGL
3. // BUG: 无法正常显示
    //img.FillMode = canvas.ImageFillOriginal
4. 为什么这个代码没有任何效果
    ```Go
   bg := canvas.NewRectangle(color.White)
   content := container.NewMax(rect, bg, img)
   go func() {
   for i := 255; i >= 0; i-- {
   rect.FillColor = color.NRGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
   rect.Refresh()
   time.Sleep(time.Millisecond * 10)
   }
   }()
   ```
   以下代码才有效果呢？
   ```go
   
   bg := canvas.NewRectangle(color.White)
   content := container.NewMax(bg, rect, img)
   go func() {
   for i := 0; i <= 255; i++ {
   rect.FillColor = color.NRGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
   rect.Refresh()
   time.Sleep(time.Millisecond * 10)
   }
   }()
    ```

container.NewMax 函数的参数顺序决定了它们在容器中的层叠顺序。在第一个代码段中，rect 是第一个参数，所以它会被放在最底层。随后的 bg 和 img 参数会被放在 rect 的上层，所以当你改变 rect 的颜色时，它的效果会被上层的对象遮挡。

在第二个代码段中，bg 是第一个参数，所以它会被放在最底层。rect 是第二个参数，所以它会被放在 bg 的上层，但在 img 的下层。当你改变 rect 的颜色时，它的效果会显示在 bg 上方，但在 img 下方，所以你可以看到效果。

另外，第二个代码段中的循环是从 0 到 255，所以 rect 会从白色渐变到黑色。在第一个代码段中的循环是从 255 到 0，所以 rect 会从黑色渐变到白色。但由于 rect 被其他对象遮挡，所以你无法看到这个效果。

总的来说，这两个代码段的不同之处在于对象的层叠顺序和渐变的方向。