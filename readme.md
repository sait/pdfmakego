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