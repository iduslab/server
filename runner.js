const child = require('child_process')
const process = child.exec("./main");
process.stdout.on("data", (content) => console.log(content));
process.stderr.on("data", (content) => console.log(content));