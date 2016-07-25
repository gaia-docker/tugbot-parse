# tugbot-parse
*tugbot-parse* is a small utility for converting a JUnit XML into JSON.

##How to install:

```bash
go get github.com/gaia-docker/tugbot-parses
```

##How to use

###Parsing a JUnit XML

```go
import "github.com/gaia-docker/tugbot-parse"

testSet, err := ToTestSet([]byte(`
	<testsuite name="Mocha Tests" tests="1" failures="1" errors="1" skipped="0" timestamp="Mon, 25 Jul 2016 14:44:37 GMT" time="0.149">
		<testcase classname="Functional tests - voting" name="open ui and check title" time="0"><failure>expected -1 to be above -1
		AssertionError: expected -1 to be above -1
		   at Assertion.assertAbove (node_modules/chai/lib/chai/core/assertions.js:571:12)
		   at Assertion.ctx.(anonymous function) [as above] (node_modules/chai/lib/chai/utils/addMethod.js:41:25)
		   at Request._callback (specs/functional/vote-page-test.js:27:97)
		   at Request.self.callback (node_modules/request/request.js:187:22)
		   at Request.&lt;anonymous&gt; (node_modules/request/request.js:1044:10)
		   at IncomingMessage.&lt;anonymous&gt; (node_modules/request/request.js:965:12)
		   at endReadableNT (_stream_readable.js:913:12)
		   at _combinedTickCallback (internal/process/next_tick.js:74:11)
		   at process._tickCallback (internal/process/next_tick.js:98:9)</failure>
   		</testcase>
	</testsuite>`))
if err != nil {
		panic(err)
}
fmt.Print(testSet)
```
