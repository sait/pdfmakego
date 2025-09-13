package main

import (
	"fmt"
	"log"
	"os"

	"rogchap.com/v8go"
)

/*
Other packages to explore
	// "rogchap.com/v8go"
	// "go.kuoruan.net/v8go-polyfills/console"
	// "github.com/nzhenev/v8go-polyfills-extended/console"
	// "github.com/hhq163/v8go"
	// "github.com/hhq163/v8go-polyfills/console"
	// "github.com/hhq163/v8go-polyfills/timers"
	// "github.com/tommie/v8go"
	// "github.com/stroiman/v8go"

*/

// Read file content, returned as string
func readFile(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("readFile: Failed %s %v", filename, err)
	}
	return string(bytes)
}

// run a JS Script in a v8 context
func runScript(ctx *v8go.Context, filename string) {
	srcCode := readFile(filename)
	res, err := ctx.RunScript(srcCode, filename)
	if err != nil {
		e := err.(*v8go.JSError)  // JavaScript errors will be returned as the JSError struct
		fmt.Println(e.Message)    // the message of the exception thrown
		fmt.Println(e.Location)   // the filename, line number and the column where the error occured
		fmt.Println(e.StackTrace) // the full stack trace of the error, if available
		log.Fatalf("Failed to run JS module: %v\n", err)
	}
	fmt.Printf("RunScript(%s): %+v\n", filename, res)
}

// Test1 Run pdfmake0212.js and myScript.js
func test1() {

	pdfmakeCode := readFile("pdfmake.js")
	myCode := readFile("myScript.js")

	iso := v8go.NewIsolate()
	defer iso.Dispose()
	ctx := v8go.NewContext(iso)
	defer ctx.Close()

	// Run PdfMake code to create pdfmake functions
	res, err := ctx.RunScript(pdfmakeCode, "pdfmakecode.js")
	if err != nil {
		e := err.(*v8go.JSError)  // JavaScript errors will be returned as the JSError struct
		fmt.Println(e.Message)    // the message of the exception thrown
		fmt.Println(e.Location)   // the filename, line number and the column where the error occured
		fmt.Println(e.StackTrace) // the full stack trace of the error, if available
		log.Fatalf("Failed to run JS module: %v\n", err)
	}
	fmt.Printf("RunScript(pdfmake): %+v\n", res)

	// Run my Code
	res, err = ctx.RunScript(myCode, "mycode.js")
	if err != nil {
		e := err.(*v8go.JSError)  // JavaScript errors will be returned as the JSError struct
		fmt.Println(e.Message)    // the message of the exception thrown
		fmt.Println(e.Location)   // the filename, line number and the column where the error occured
		fmt.Println(e.StackTrace) // the full stack trace of the error, if available
		log.Fatalf("Failed to run JS module: %v\n", err)

	}
	fmt.Printf("RunScript(myCode): %+v\n", res)

	// Show Context Variable: myBase64
	myb64, err := ctx.Global().Get("myBase64")
	if err != nil {
		panic(err)
	}
	fmt.Printf("globalThis.myBase64 : %s\n", myb64.String())
}

// Test2 Run pdfkit
// wget https://cdn.jsdelivr.net/npm/pdfkit@latest/js/pdfkit.standalone.js
func test2() {
	iso := v8go.NewIsolate()
	defer iso.Dispose()
	ctx := v8go.NewContext(iso)
	defer ctx.Close()
	runScript(ctx, "TextEncoder.polyfill.js")
	runScript(ctx, "pdfkit.standalone.js")
	runScript(ctx, "myPdfKitScript.js")
}

func main() {
	fmt.Printf("Tests using v8 Version: %s\n", v8go.Version())
	// test1()
	test2()
	fmt.Printf("This is the end !\n")
}
