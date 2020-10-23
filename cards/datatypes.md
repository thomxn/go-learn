## Basic datatypes

1. `bool`
2. `string`
3. `int`
4. `float64`

`:=` is used for variable declarations.
Thus `var card = "Hai"` and `card:="Hai"` are same

There are two datatypes for storing lists
1. `array` - Fixed length
2. `slice` - Grow or shrink

Both array and slice should contain similar datatype

Specify the dependency files in the CLI along with main.go `go run main.go deck.go`

## Functions

**Receivers** sets up methods on variables that we create. Receiver methods specify the type just after the `func` declaration
Eg: `func (d deck) print() { }`

Use `_` as placeholder for variables
Eg: `for _, value := range cardValues`

### Range
Use array indexes inside `[]` to get a subset from an existing array
`array[startIndexIncluding: endIndexExcluding]`