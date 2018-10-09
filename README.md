# LEARNING GO BY WRITING TESTS

## Running the tests
`go test ./...` runs all tests in the folder
`go test ./... -v` runs all tests in the folder with verbose output
`go test ./arrays` runs all test in a package/directory (in this case `./arrays`)
`go test ./... -cover` runs all tests and outputs a coverage report
`go test ./... -bench=.` runs all tests and includes benchmark reports as well


## Docs
Add a test called ExampleTestName() to your _test files
The test should have an `fmt.Println(output)` 
You can also add `//` comments before the func under tests, these will be included in the docs

Run `godoc -http=:6060` to start a local page in `http://localhost:6060/pkg/`