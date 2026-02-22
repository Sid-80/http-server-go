## Why * is used in http.Request

The * in *http.Request means it's a pointer to http.Request, not a value. The standard library uses a pointer here for these reasons:

1. Performance & Efficiency
http.Request is a large struct with many fields (URL, Method, Header, Body, etc.). Passing by value would copy all fields, which is expensive. Passing a pointer avoids the copy.

2. Memory Efficiency
Large structs are better passed as pointers to avoid stack copying.

3. Standard Library Convention
The net/http package consistently uses *http.Request in handler signatures. For example:
http.HandlerFunc expects: func(http.ResponseWriter, *http.Request)
http.Handler interface uses: ServeHTTP(http.ResponseWriter, *http.Request)

4. Flexibility
While handlers usually shouldn't modify the request, using a pointer allows middleware or other code to modify it when needed.

Example Comparison:
// Using pointer (standard, efficient)
func handler(w http.ResponseWriter, r *http.Request) {
    // r is a pointer - no copying of the large struct
}

// Using value (would work but inefficient)
func handler(w http.ResponseWriter, r http.Request) {
    // r is a copy - all fields are duplicated in memory
}

In short, *http.Request is the standard, efficient way to pass the request to handlers in Go.


## Why json used here??

The json:"..." tags are struct tags that control how Go encodes/decodes the struct to/from JSON.

Purpose of JSON Tags
1. Field name mapping: When encoding to JSON, the tag sets the JSON key name. 2. Without tags, Go uses the struct field name as-is.
    JSON encoding/decoding: Used by the encoding/json package when you call:
    - json.Marshal() - converts Go struct → JSON string
    - json.Unmarshal() - converts JSON string → Go struct

Example:
type Article struct {
    Title       string `json:"Title"`       // Maps to "Title" in JSON
    Description string `json:"Description"` // Maps to "Description" in JSON
    Content     string `json:"Content"`     // Maps to "Content" in JSON
}

When you encode this struct to JSON, it produces:
{
  "Title": "Some Title",
  "Description": "Some Description",
  "Content": "Some Content"
}
