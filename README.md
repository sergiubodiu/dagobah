Step 1: Working with feeds
--------------------------

go get -u github.com/jteeuwen/go-pkg-rss

go run main.go fetch --rsstimeout=1

Step 2: Storing in the database
-------------------------------

go get -u gopkg.in/mgo.v2
$ wmic os get caption
$ wmic os get osarchitecture
https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/

### Starting MongoDB Server

   c:\Program Files\MongoDB\Server\3.4\bin\mongod.exe


### Looking at the data

   db.items.findOne()

Step 3: Building a web server
-----------------------------

### Getting Gin
    go get -u github.com/gin-gonic/gin

### Running it
    go run main.go server

### Check it out
    http://localhost:1138

Step 4: Building a dynamic server
---------------------------------

### Getting Go.Rice
    go get -u github.com/GeertJohan/go.rice

### Get your static & templates on
    c.spf13.com/OSCON/step6-static.zip