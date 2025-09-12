# pdfmakego

## Execute pdfmake from golang

Get pdfmake.js from: https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.2.12/pdfmake.js

Run you will get
```
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
- find: ```FileSaver```
- Comment this lines
```
54432 // var FileSaver = __webpack_require__(42616);
54433 // var saveAs = FileSaver.saveAs;
```