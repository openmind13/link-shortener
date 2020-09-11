// let error = true

// let res = [
//   db.container.drop(),
//   db.container.createIndex({ myfield: 1 }, { unique: true }),
//   db.container.createIndex({ thatfield: 1 }),
//   db.container.createIndex({ thatfield: 1 }),
//   db.container.insert({ myfield: 'hello', thatfield: 'testing' }),
//   db.container.insert({ myfield: 'hello2', thatfield: 'testing' }),
//   db.container.insert({ myfield: 'hello3', thatfield: 'testing' }),
//   db.container.insert({ myfield: 'hello3', thatfield: 'testing' }),
//   db.other.
// ]

// printjson(res)

// if (error) {
//   print('Error, exiting')
//   quit(1)
// }


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