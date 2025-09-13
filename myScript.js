var docDefinition = {
    content: [
        'This is a Very Nice PDF !!!.',
        { text: 'It will be converted to Base64.', style: 'header' }
    ]
    };

const pdfDocGenerator = pdfMake.createPdf(docDefinition);


globalThis.myBase64 = "This is a variable named: myBase64 created in myScript.js";

pdfDocGenerator.getBase64(function(base64) {
    console.log("PDF en base64:", base64)
    globalThis.myBase64 = base64;
    return base64;
});


// V8.RunScript will return the last expression as result
myBase64;

pdfDocGenerator.version;
// JSON.stringify(pdfDocGenerator);

// pdfDocGenerator.getBase64();
// Failed to run JS module: getBase64 is an async method and needs a callback argument
