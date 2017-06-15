const express = require('express');
const hostname = 'localhost';
const port = 8092;
const app = express();
const request = require('request');
var router = express.Router();
app.set('view engine', 'ejs');

app.use(express.static(__dirname));

let http = require('http');

function jsonToMap(jsonStr) {
    return new Map(JSON.parse(jsonStr));
}


var temp = "there will be text"
router.get('/', (req, res) => {
    var temp
    request('http://localhost:8080/agregatedData', function (error, response, body) {
        temp = body;
        console.log(temp)
        //console.log('body:', body); // Print the HTML for the Google homepage. 
        //console.log(jsonToMap(body))
        console.log(temp)
        var decodeHtmlEntity = function (x) {
            return x.replace(/&#(\d+);/g, function (match, dec) {
                return String.fromCharCode(dec);
            });
        };
        temp = decodeHtmlEntity(temp)
        res.render('./index.ejs', {
            num: temp
        });
    });

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