
console.log("\nstarting myPdfKitScript.js...");

const doc = new PDFDocument();
console.log( "PDFDocument version is: ",doc.version)
/*

// Try pdfkit example of node using fs.createWriteStream
doc.pipe(fs.createWriteStream('/path/to/file.pdf')); 
doc.addPage()
doc.text('Hello world!')

// finalize the PDF and end the stream
doc.end();
*/
console.log("enddding myPdfKitScript.js");
"OK";