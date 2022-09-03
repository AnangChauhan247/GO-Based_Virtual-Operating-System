package main

import (
	"fyne.io/fyne/v2"
	"strconv"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
    w.Resize(fyne.NewSize(600, 400))
	output:=""
	input:= widget.NewLabel(output)

	historyStr:=""
	history:=widget.NewLabel(historyStr)
    var historyArr[]string;
	isHistory:=false
	//history button
	historyBtn:=widget.NewButton("History", func(){
		if isHistory{
             historyStr=""
		}else{
			for i:=len(historyArr)-1; i>=0; i--{
			historyStr=historyStr+historyArr[i];
			historyStr+="\n";
		}
	   }
	    isHistory=!isHistory
		history.SetText(historyStr);
	})
    
	//back button
	backBtn:=widget.NewButton("Back", func(){
		if len(output)>0{
		output=output[:len(output)-1];
		input.SetText(output)
		}
	})

	//clear button
	clearBtn:=widget.NewButton("Clear", func(){
		output="";
		input.SetText(output);
	})

	// open bracket button (
	openBracketBtn:=widget.NewButton("(",func(){
		output=output+"(";
		input.SetText(output)
	})	

	//close bracket button )
	closeBracketBtn:=widget.NewButton(")",func(){
		output=output+")";
		input.SetText(output);
	})

	//divide button /
	divideBtn:=widget.NewButton("/",func(){
		output=output+"/";
		input.SetText(output)
	})

	// 7 button 
	sevenBtn:=widget.NewButton("7",func(){
		output=output+"7";
		input.SetText(output)
	})

	// 8 button 
	eightBtn:=widget.NewButton("8",func(){
		output=output+"8";
		input.SetText(output)
	})

	// 9 button
	nineBtn:=widget.NewButton("9",func(){
		output=output+"9";
		input.SetText(output)
	})

	// multiply button *
	multiplyBtn:=widget.NewButton("*",func(){
		output=output+"*";
		input.SetText(output)
	})

	// 4 button
	fourBtn:=widget.NewButton("4",func(){
		output=output+"4";
		input.SetText(output)
	})

	// 5 button
	fiveBtn:=widget.NewButton("5",func(){
		output=output+"5";
		input.SetText(output)
	})

	// 6 button
	sixBtn:=widget.NewButton("6",func(){
		output=output+"6";
		input.SetText(output)
	})

	// minus button -
	minusBtn:=widget.NewButton("-",func(){
		output=output+"-";
		input.SetText(output)
	})

	// 1 button
	oneBtn:=widget.NewButton("1",func(){
		output=output+"1";
		input.SetText(output)
	})

	// 2 button
	twoBtn:=widget.NewButton("2",func(){
		output=output+"2";
		input.SetText(output)
	})

	// 3 button
	threeBtn:=widget.NewButton("3",func(){
		output=output+"3";
		input.SetText(output)
	})

	// plus button +
	plusBtn:=widget.NewButton("+",func(){
		output=output+"+";
		input.SetText(output)
	})

	// 0 button
	zeroBtn:=widget.NewButton("0",func(){
		output=output+"0";
		input.SetText(output)
	})

	// dot button .
	dotBtn:=widget.NewButton(".",func(){
		output=output+".";
		input.SetText(output)
	})

	// equal button
	equalBtn:=widget.NewButton("=",func(){
		expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil{
			result, err := expression.Evaluate(nil);
			if err==nil{
				ans:=strconv.FormatFloat(result.(float64), 'f', -1, 64);
				strToAppend:=output+" = "+ans;
				historyArr=append(historyArr,strToAppend);
				output=ans;
			}else{
				output="Error";
			}
		}else{
			output="Error";
		}
		input.SetText(output);
	})

	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			//row-1
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
            //row-2
			container.NewGridWithColumns(4,
				clearBtn,
				openBracketBtn,
				closeBracketBtn,
				divideBtn,
			),
			//row-3 
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiplyBtn,
			),
			//row-4
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				minusBtn,
			),
			//row-5
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),
			//row-6
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
				   zeroBtn,
				   dotBtn,
				),
				equalBtn,
			),
		),
	))

	w.ShowAndRun()
}