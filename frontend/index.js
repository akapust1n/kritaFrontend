const express = require('express');
const hostname = 'localhost';
const port = 80;
const app = express();
const request=require('request');
var router = express.Router();
app.set('view engine', 'ejs');

app.use(express.static(__dirname));

let http = require('http');



var temp = "there will be text"
router.get('/', (req, res) => {

// request('http://37.139.19.10/text', function (error, response, body) {
// temp = body;  
// //console.log('body:', body); // Print the HTML for the Google homepage. 
// });
    res.render('./index.ejs', {num: 1});
    // res.render('./index.ejs', {});
    //res.render('./views/index.ejs', {});


    
});

app.use(router);


app.listen(port, () => {
    console.log(`
        Server is running at http://${hostname}:${port}/ 
        Server hostname ${hostname} is listening on port ${port}!
    `);
});