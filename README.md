XML-to-JSON ETL
===============

Extracts information form XML feed and transforms to JSON

The program uses sqlite3 as a persistent datastore to keep parsed data.
All XML data file(challenge.xml) is located in testdata directory.
The default port is :8080 and sqlite3 database file name can be changed with flags
```
age of ./xmlroom:
  -db string
        database file path (default "./db")
  -port string
        server port (default ":8080")

```
