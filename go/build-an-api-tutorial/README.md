# Go API tutorial
We will use [gin]() to create our simple API.


## Getting started
First create a `go.mod` file to track all of the dependencies.
```
go mod init example/go-api-tutorial
```

Install the gin framework
```
go get github.com/gin-gonic/gin
```
- note that `go.mod` has been modified


## main.go
We will create a library API that allow the following functions:
- check in a book
- check out a book
- add a book
- view all books
- get book by id

### Details

We will store our books in a [struct]().

Field names should start with a **capital letter** so that Go recognizes them as an [exported field](). This will make the fields **public** to all Go modules, allowing us to use these fields with other Go files that will need to access these fields). The JSON field should be lowercase because JSON needs lowercase values. Go will serialize the fields into lowercase values.

There's a data structure that is a **slice of books** (lines 16-20)

The **main** function will represent our API. The **router** is responsible for routing requests to endpoint for the API. The router will send requests to other functions in our code.

The **getBooks** function accepts a **context** argument

### Testing the GET API
Run the file
```
go run main.go
```

Test the API
```
curl www.localhost:8080/books
```

```
❯ curl www.localhost:8080/books
[
    {
        "id": "1",
        "title": "In Search of Lost Time",
        "author": "Marcel Proust",
        "quantitiy": 2
    },
    {
        "id": "2",
        "title": "The Great Gatsby",
        "author": "F. Scott Fitzgerald",
        "quantitiy": 5
    },
    {
        "id": "3",
        "title": "War and Peace",
        "author": "Leo Tolstoy",
        "quantitiy": 6
    }
]
```

### Testing the POST API
Send a POST request using the `./body.json` file. 
```
cd ~/learn/go/api-tutorial/
curl www.localhost:8080/books --header "Content-Type: application/json" -d @body.json --request "POST"
```

If you run the GET request, you will see that new book added to the data structure.
```
curl www.localhost:8080/books
```
