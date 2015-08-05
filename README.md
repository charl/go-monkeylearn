#### go-monkeylearn

Go client that interacts with http://www.monkeylearn.com.

#### Install

Simply install the package to your $GOPATH with the go tool from a shell:

```bash
$ go get github.com/charl/go-monkeylearn
```

#### Use

```go
import "charl/go-monkeylearn"

...

func main() {
    // Get your API Token from your MonkeyLearn account at https://app.monkeylearn.com/accounts/user/settings/#.
    token := "d3ad6eefd3ad6eefd3ad6eefd3ad6eef"
    client := gomonkeylearn.NewClient(token)

    // Classify some texts using the News Categorizer (https://app.monkeylearn.com/categorizer/projects/cl_hS9wMk9y/).
    category := "News Categorizer"
    texts := []string{"First text to classify", "Second text to classify"}
    classifications, err := client.Classify(category, texts)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(string(classifications))
    // Prints:
    // {"result": [[{"probability": 0.65, "label": "Arts & Culture"}, {"probability": 0.72, "label": "Books"}]]}
}
```

#### Test

Assuming you have go properly installed, $GOPATH set and performed the package install above:

```bash
$ cd $GOPATH/src/github.com/charl/go-monkeylearn
$ go test
```

Testing coverage:

```bash
$ cd $GOPATH/src/github.com/charl/go-monkeylearn
$ go test -cover
```

Testing with a coverage profile:

```bash
$ cd $GOPATH/src/github.com/charl/go-monkeylearn
$ go test -coverprofile=cover.out
```

The cover.out file will contain the profile.
