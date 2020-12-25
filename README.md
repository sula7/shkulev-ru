### An HTTP server to validate parentheses

#### Run
`go build -o ./shru . && ./shru`

#### Endpoints
POST `/api/v1/validate`  
A JSON array of strings like:  
`[ "[]", "(}"]`  
The response would be:  
`{"success":true,"message":"","data":[{"value":"[]","is_valid":true},{"value":"(}","is_valid":false}]`
