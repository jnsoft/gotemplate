# Init
```
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

go mod init github.com/jnsoft/<my-project>

go get -u github.com/google/uuid
go get -u github.com/jnsoft/jngo

# debug:
dlv debug src/main.go
```

# Build and Test
```
go test ./... -v
go test -v -race ./src/pqueue/ # identify race conditions

go run  ./src/main.go

#linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .bin/app ./src/main.go

#win
GOOS=windows GOARCH=amd64 go build -o .bin/app-windows-amd64.exe ./src/main.go

#run
.bin/app
```

```
type Point struct {
    X, Y float64
}
```

```
hashmap := make(map[string]int)
hashmap["A"] = 25
value, exists := hashmap["A"]
isEmpty := len(hashmap) == 0
for key, value := range hashmap {
        fmt.Printf("%s -> %d\n", key, value)
}
toSlice := make([]int, 0, len(s.data))
    for key := range s.data {
        result = append(result, key)
}
delete(hashmap, "A")
```

RWMutex (instead of sync.Mutex) allows multiple readers to access the resource concurrently, while ensuring that only one writer can access it at any time. Use in read-heavy scenarios.
m.mu.RLock() allows multiple readers
m.mu.Lock() allows only one writer


## Code Review
### Input
* new input?
* changes in handling the input
* changes in storing
* changes in how the input is used

### AAA
* Authentication - is the actor who they say they are?
* Authorization - is the actor allowed to do this?
* Auditing - have we made a note of what happened?

### New dependencies
* Trusted?
* Popular?
* Well maintained?
* Do we need it?

[Check with Snyk Advisor](https://snyk.io/advisor)


## Functional programming principles
Helps produce readable, predictable and testable code
* Pure functions (given the same input always returns the same result and no side effects, or simply the function can be substituted by a lookup table like a mathematical function)
* Referential transparency (any bound variable or subexpression can be replaced by it's value without changing the meaning of the program, regardless of the order in which they are evaluated)
* Immutability (once a value is declared, it is unchangeable)
* Functions as first class entities (functions can be passed as arguments, returned from other functions, stored in data structures and assigned to variables)
* Higher order funtions (functions that takes functions as values or returns functions, like map, filter and reduce)

