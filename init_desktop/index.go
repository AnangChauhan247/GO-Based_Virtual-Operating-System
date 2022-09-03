 package main

 import(
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	
 )

 var myApp fyne.App=app.New()
 var myWindow fyne.Window=myApp.NewWindow("Sharp OS")
 var btn1 fyne.Widget
 var btn2 fyne.Widget
 var btn3 fyne.Widget
 var btn4 fyne.Widget
 var btn5 fyne.Widget
 var img fyne.CanvasObject
 var DeskBtn fyne.Widget
 var panelContent *fyne.Container
 func main(){
	 myApp.Settings().SetTheme(theme.DarkTheme())
	img=canvas.NewImageFromFile("D:\\GO_lang\\Virtual_os\\init_desktop\\wall.jpg")

	btn1=widget.NewButtonWithIcon("Weather App",theme.InfoIcon(), func(){
		showWeatherApp(myWindow)
	})

	btn2=widget.NewButtonWithIcon("Calculator",theme.ContentAddIcon(), func(){
		showcalC()
	})

	btn3=widget.NewButtonWithIcon("Gallery",theme.StorageIcon(), func(){
		showGal(myWindow)
	})

	btn4=widget.NewButtonWithIcon("Text Editor",theme.DocumentIcon(), func(){
		showTE()
	})

	btn5=widget.NewButtonWithIcon("IP_info",theme.DocumentIcon(), func(){
		ipApp()
	})

	DeskBtn=widget.NewButtonWithIcon("Home",theme.HomeIcon(), func(){
		myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	})

	panelContent=container.NewVBox((container.NewGridWithColumns(5, DeskBtn,btn1,btn2,btn3,btn4,btn5)))

	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen();
	myWindow.SetContent(
		container.NewBorder(panelContent, nil ,nil, nil, img),
	)
	myWindow.ShowAndRun();
	
 }	