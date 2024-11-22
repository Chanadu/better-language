Rewrite Scanner as an Interface and a Struct to get defualt value
Example:
```go
package candidate
// Exporting interface instead of struct
type Candidate interface {}
// Struct is not exported
type candidate struct {
Name string
Votes uint32 // Defaults to 0
}
// We are forced to call the constructor to get an instance of candidate
func NewCandidate(name string) Candidate {
return candidate{name, 0}  // enforce the default value here
}
```
