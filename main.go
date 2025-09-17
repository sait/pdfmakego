package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

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

func saveB64(base64string string, filename string) {
	// Decode the base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(base64string)
	if err != nil {
		log.Fatalf("Error decoding base64 string: %v", err)
	}
	// Write the decoded bytes to a file
	err = os.WriteFile(filename, decodedBytes, 0644) // 0644 grants read/write for owner, read-only for others
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
	fmt.Printf("File '%s' saved successfully from base64 string.\n", filename)
}

func main() {
	fmt.Printf("Tests using v8 Version: %s\n", v8go.Version())

	// Prepare v8 context, inject timers and console
	// to have setTimeout() and console.Log() in global context
	iso := v8go.NewIsolate()
	global := v8go.NewObjectTemplate(iso)
	if err := timers.InjectTo(iso, global); err != nil {
		panic(err)
	}
	defer iso.Dispose()
	ctx := v8go.NewContext(iso, global)
	defer ctx.Close()
	console.InjectTo(ctx)

	// Run Script
	// Please be sure to:
	// 1) Have pdfmake.js download it
	// wget https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.2.12/pdfmake.js
	// 2) Comment the 2 FileSaver references
	// 54432: // var FileSaver = __webpack_require__(42616);
	// 54433: // var saveAs = FileSaver.saveAs;
	runScript(ctx, "pdfmake.js")
	runScript(ctx, "myScript.js")

	// Wait the script to finish
	time.Sleep(1 * time.Second)

	//  Get result variable:myBase64 and save the file
	myb64, err := ctx.Global().Get("myBase64")
	if err != nil {
		panic(err)
	}
	saveB64(myb64.String(), "myDocument.pdf")

	fmt.Printf("ending main.go  !\n")
}
