const doc = new PDFDocument();
doc.pipe(fs.createWriteStream('/path/to/file.pdf')); 

doc.addPage()
doc.text('Hello world!')

// finalize the PDF and end the stream
doc.end();