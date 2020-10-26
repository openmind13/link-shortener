// for initializing mongodb into docker container

var MongoClient = require('mongodb').MongoClient;
var url = "mongodb://localhost:27017/linkshortener";
MongoClient.connect(url, function (err, db) {
    if (err) throw err;
    db.createCollection("links", function (err, res) {
        if (err) throw err;
        console.log("Collection is created!");
        db.close();
    });
});  