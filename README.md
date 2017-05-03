# jsonandgo

### Encoding
This is easy task because only one way to do that exists. We have object and want to convert it to JSON. 

We can do it using **json.Marshal** method:
```go
objects := []Object{}
b, _ := json.Marshal(objects)
fmt.Println(string(b))
```
Output JSON will be:
```
[{"person":{"name":{"first":"Leonid","last":"Bugaev","fullName":"Leonid Bugaev"},
"github":{"handle":"buger","followers":109},
"avatars":[{"url":"https://avatars1.githubusercontent.com/u/14009?v=3\u0026s=460","type":"thumbnail"},
{"url":"https://avatars1.githubusercontent.com/u/14009?v=3\u0026s=460","type":"thumbnail"}]},
"company":{"name":"Acme"}}]
```
### Decoding
This is pretty difficult task therefore many libraries exist to simplify that job.

I used 3 methods for decoding JSON data to go structs:
1. Using standard [encoding/json](https://godoc.org/encoding/json) library - [code](https://github.com/hesidoryn/jsonandgo/blob/master/standard.go)
2. Using [github.com/buger/jsonparser](https://github.com/buger/jsonparser) - [code](https://github.com/hesidoryn/jsonandgo/blob/master/jsonparser.go)
3. Using [github.com/Jeffail/gabs](https://github.com/Jeffail/gabs) - [code](https://github.com/hesidoryn/jsonandgo/blob/master/gabs.go)

I can say that every method has some benefits: 
* more understandable program with known types when using `encoding/json` library;
* `buger/jsonparser` allows to work with JSON-data in functional style;
* `Jeffail/gabs` allows to work in imperative form and gives many syntactic sugar: `Exists`, `Search` and `SearchP` methods etc.
