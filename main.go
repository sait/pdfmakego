package main

import (
	"fmt"
	"log"
	"os"

	"go.kuoruan.net/v8go-polyfills/console"
	"go.kuoruan.net/v8go-polyfills/timers"
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

// Test1 Run pdfmake.js and myScript.js
// wget https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.2.12/pdfmake.js
func test1(ctx *v8go.Context) {
	runScript(ctx, "pdfmake.js")
	runScript(ctx, "myScript.js")

	// Show Context Variable: myBase64
	myb64, err := ctx.Global().Get("myBase64")
	if err != nil {
		panic(err)
	}
	fmt.Printf("globalThis.myBase64 : %s\n", myb64.String())

}

// Test2 Run pdfkit
// wget https://cdn.jsdelivr.net/npm/pdfkit@latest/js/pdfkit.standalone.js
func test2(ctx *v8go.Context) {
	runScript(ctx, "TextEncoder.polyfill.js")
	runScript(ctx, "pdfkit.standalone.js")
	runScript(ctx, "myPdfKitScript.js")
}

func main() {
	fmt.Printf("Tests using v8 Version: %s\n", v8go.Version())

	// Prepare v8 context, inject timers and console for setTimeout() and console.Log()
	iso := v8go.NewIsolate()
	global := v8go.NewObjectTemplate(iso)
	if err := timers.InjectTo(iso, global); err != nil {
		panic(err)
	}
	defer iso.Dispose()
	ctx := v8go.NewContext(iso, global)
	defer ctx.Close()
	console.InjectTo(ctx)

	//test1(ctx)
	test2(ctx)
	fmt.Printf("ending main.go  !\n")
}
