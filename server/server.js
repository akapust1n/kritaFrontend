const express = require('express');
 const fs = require('fs');
const hostname = '188.166.160.252';
 const port = 80;
 const app = express();
 
 let cache = [];// Array is OK!
 cache[0] = fs.readFileSync( __dirname + '/index.html');
 
 
 app.get('/', (req, res) => {
     res.setHeader('Content-Type', 'text/html');
     res.send( cache[0] );
 });
 
 
 app.listen(port, () => {
     console.log(`
         Server is running at http://${hostname}:${port}/ 
         Server hostname ${hostname} is listening on port ${port}!
     `);
 });
