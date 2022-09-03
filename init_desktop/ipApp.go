package main

import (
	"encoding/json"
	"io/ioutil"
    "net/http"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ipApp(){
	//New App
	// a:=app.New()
	//New Title and Window
	w:=myApp.NewWindow("IP check")
	//Resize
	w.Resize(fyne.NewSize(400,400))
	//UI 
	 //Title
	labelTitle:=widget.NewLabel("IP_lookup")
	labelIP:=widget.NewLabel("Your IP is:- ")
	label_Value:=widget.NewLabel("...")
	label_City:=widget.NewLabel("...")
	label_Country:=widget.NewLabel("...")
	btn:= widget.NewButton("Run",func() {
		label_Value.Text=myIP()
		label_Value.Refresh()

		label_City.Text=myCity()
		label_City.Refresh()

		label_Country.Text=myCountry()
		label_Country.Refresh()
	})
		
		data_con:=
			container.NewVBox(
				labelTitle,
				labelIP,
				label_Value,
				label_City,
				label_Country,
				btn,
			)
		
	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,data_con),
	)
    
	//ShowAndRun
	w.Show()
}

func myIP()string{
	req, err:=http.Get("http://ip-api.com/json/")
	if err!=nil{
		return err.Error()
	}
	defer req.Body.Close()
	body, err:=ioutil.ReadAll(req.Body)
	if err!=nil{
		return err.Error()
	}
    var ip IP
	json.Unmarshal(body,&ip)
	return ip.Query
}

func myCity()string{
	req, err:=http.Get("http://ip-api.com/json/")
	if err!=nil{
		return err.Error()
	}
	defer req.Body.Close()
	body, err:=ioutil.ReadAll(req.Body)
	if err!=nil{
		return err.Error()
	}
    var ip IP
	json.Unmarshal(body,&ip)
	return ip.City
}

func myCountry()string{
	req, err:=http.Get("http://ip-api.com/json/")
	if err!=nil{
		return err.Error()
	}
	defer req.Body.Close()
	body, err:=ioutil.ReadAll(req.Body)
	if err!=nil{
		return err.Error()
	}
    var ip IP
	json.Unmarshal(body,&ip)
	return ip.Country
}
type IP struct{
	Query string
	Country string
	City string
}