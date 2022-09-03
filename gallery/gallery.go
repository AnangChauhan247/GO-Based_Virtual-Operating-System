package main

import (
	"fmt"
    "io/ioutil"
    "log"
	"strconv"
	"strings"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("sharp Image Viewer")
		w.Resize(fyne.NewSize(800,600));

	root_src:="D:\\Jerry\\images\\wallpaper";

	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }
    var picsArr []string;
    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		

		if !file.IsDir(){
			extension:=strings.Split(file.Name(), ".")[1];
			if extension=="png"||extension=="jpeg" ||extension=="jpg"{
				picsArr=append(picsArr,root_src+"\\"+file.Name());

			}
		}
    }
	tabs := container.NewAppTabs(
		
		container.NewTabItem("Image 0",canvas.NewImageFromFile(picsArr[0]) ),
	)
	for i:=1; i<len(picsArr); i++{
		image:=canvas.NewImageFromFile(picsArr[i])
		tabs.Append(container.NewTabItem("Image "+strconv.Itoa(i),image))
		}
			 
	//to make listing vertical	
    // tabs.SetTabLocation(container.TabLocationLeading);
	w.SetContent(tabs);
	w.ShowAndRun()
}