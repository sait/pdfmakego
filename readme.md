# pdfmakego

## Execute pdfmake from golang

### Build and Run

Build program:
```
# Before test this program, get pdfmake.js, using:
wget https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.2.12/pdfmake.js
go mod tidy
go build
./pdfmakego

```

### Error: TypeError: Cannot read properties of undefined (reading 'navigator')
```
./pdfmakego

TypeError: Cannot read properties of undefined (reading 'navigator')
pdfmakecode.js:61176:1184
TypeError: Cannot read properties of undefined (reading 'navigator')
    at Object.<anonymous> (pdfmakecode.js:61176:1184)
    at pdfmakecode.js:61175:35
    at 42616 (pdfmakecode.js:61176:109)
    at __webpack_require__ (pdfmakecode.js:74137:42)
    at 45314 (pdfmakecode.js:54432:17)
    at __webpack_require__ (pdfmakecode.js:74137:42)
    at 36164 (pdfmakecode.js:21432:34)
    at __webpack_require__ (pdfmakecode.js:74137:42)
    at pdfmakecode.js:74173:37
    at pdfmakecode.js:74176:12
2025/09/12 16:26:28 Failed to run JS module: TypeError: Cannot read properties of undefined (reading 'navigator')

```

Comment some request in pdfmake.js
- Open pdfmake.js
- Search for: ```FileSaver```
- Comment this 2 lines
- Save
```
54432 // var FileSaver = __webpack_require__(42616);
54433 // var saveAs = FileSaver.saveAs;
```

### Final Run
```
ignacio@igt25:~/projects/pdfmakego$ ./pdfmakego

Tests using v8 Version: 10.9.194.9-v8go
RunScript(pdfmake): undefined
RunScript(myCode): This is a variable named: myBase64 created in myScript.js
globalThis.myBase64 : This is a variable named: myBase64 created in myScript.js
This is the end !

```

### Testing PdfKit

In test2() we are trying to run a script using PdfKit because [PdfMake](http://pdfmake.org/#/) is based on [PdfKit](https://pdfkit.org/)

Download pdfkit.standalone.js
```

wget https://cdn.jsdelivr.net/npm/pdfkit@latest/js/pdfkit.standalone.js

./pdfmakego

Tests using v8 Version: 10.9.194.9-v8go
RunScript(TextEncoder.js): function(octets){var string="",i=0;while(i<octets.length){var octet=octets[i],bytesNeeded=0,codePoint=0;octet<=...<omitted>...g}
RunScript(pdfkit.standalone.js): undefined
ReferenceError: setTimeout is not defined
pdfkit.standalone.js:43457:30
ReferenceError: setTimeout is not defined
    at runTimeout (pdfkit.standalone.js:43457:30)
    at process.nextTick (pdfkit.standalone.js:43561:9)
    at maybeReadMore (pdfkit.standalone.js:45703:13)
    at addChunk (pdfkit.standalone.js:45469:3)
    at readableAddChunk (pdfkit.standalone.js:45445:11)
    at Readable.push (pdfkit.standalone.js:45411:10)
    at PDFDocument._write (pdfkit.standalone.js:5899:10)
    at new PDFDocument (pdfkit.standalone.js:5805:10)
    at myPdfKitScript.js:1:13
2025/09/12 17:21:48 Failed to run JS module: ReferenceError: setTimeout is not defined
```

#### setTimeout is not defined

We need to add setTimeout to our context, please check https://github.com/sait/polyfills-kuoruan