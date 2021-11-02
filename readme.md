# Go Count Stuff!
This is a somewhat useless application. It runs a bunch of count queries (select count(1) ...) and compares the result with the expected one.

The purpose is to frequently run some queries and ensure that certain data in the database hasnt changed.


# How to use

## Configuration file
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

## Example usage

### Configuration file
Filename: ```config_file.json```
```json
{
  "connectionString": "Server=tcp:localhost,1433;Initial Catalog=mydb;Persist Security Info=False;User ID=root;Password=super#secret@42;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;MultipleActiveResultSets=True;",
  "queries": [
    {
      "queryName": "Count general Audit Summary entries",
      "sql": "select count(1) from AUDIT_SUMMARY where id <= 7",
      "expectedCount": 10
    },
    {
      "queryName": "Count Audit detail rows",
      "sql": "select count(1) from AUDIT_DETAILS where id <= 12",
      "expectedCount": 12
    }
  ]
}
```

### Command line
```shell
cd c:\go-count-db-stuff
gocount.exe -config=config_file.json
```

### Output
```shell
** GO Count Stuff! (1.0.0) **
File 'dev.config.json' was relative and was converted to absolute path: 'c:\go-count-db-stuff\dev.config.json'
Found '2' queries to run.
Connecting...
Connected!
Count general Audit Summary entries: FAILED
Count Audit detail rows: SUCCESS
All done! (elapsed time: 0.4723352s)
```