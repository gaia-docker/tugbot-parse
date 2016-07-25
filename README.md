# tugbot-parse
*tugbot-parse* is a small utility for converting a JUnit XML into JSON.

##How to install:

```bash
go github.com/gaia-docker/tugbot-parses
```

##How to use

###Parsing a JUnit XML

```go
import "github.com/gaia-docker/tugbot-parse"

testSet, err := ToTestSet(getMyJUnitXML())
if err != nil {
		panic(err)
}
fmt.Print(testSet)
```
