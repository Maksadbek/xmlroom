Interview assignment by Housinganywhere(that is failed)

XML-to-JSON ETL
===============

Extracts information form XML feed and transforms to JSON

The application is separated into 2 services.
1. The XML Parser that parses data and puts into database
2. The WEB API that listens on :8080 by default port and extracts data from datastore and sends it transforming to JSON

The program uses sqlite3 as a persistent datastore to keep parsed data.
All XML data file(challenge.xml) is located in testdata directory.

The API service receives the sqltie3 database file and web server port from flags
```
Usage of ./webapi:
  -db string
        database file path (default "./db")
  -port string
        server host and port (default ":8080")

```

The XML Parser(xmldumber) also receives database file path from flags
```
Usage of ./xmldumper:
  -db string
        database file path (default "./db")

```

The ```Makefile``` included commands to run tests, to build and to lint.
```make build``` - builds both services
```make test``` - runs test
```make lint``` - runs code linter 
