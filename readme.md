# Go Count Stuff!
This is a somewhat useless application. It runs a bunch of count queries (select count(1) ...) and compares the result with the expected one.

The purpose is to frequently run some queries and ensure that certain data in the database hasnt changed.

# How to use
## Configuraiton file
```json
{
  "connectionString": "conn string",
  "queries": [
    {
      "queryName": "Foo",
      "sql": "select count(1) from bla where x = y and z = 42",
      "expectedCount": 1
    }
  ]
}
```
- **connectionString**: Connection string to the database
- **queries**: List of queries that will be executed.
- **queries[x].queryName**: Name of the query that will be executed. Used as reference in log messages
- **queries[x].sql**: Sql command that will be used.
- **queries[x].expectedCount**: Count result that we expect for this query.

> All fields are required.


## Command line
```shell
gocount.exe -config=./path/to/config_file.json
```

