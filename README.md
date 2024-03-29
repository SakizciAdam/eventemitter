# Event Emitter

Event Emitter is  lightweight Go package that implements a simple Event Bus.

[![Go Reference](https://pkg.go.dev/badge/github.com/SakizciAdam/eventemitter.svg)](https://pkg.go.dev/github.com/SakizciAdam/eventemitter)

## Installation

```bash
go get https://github.com/SakizciAdam/eventemitter
```
## Usage

```go
package main

import (
	"fmt"

	"github.com/SakizciAdam/eventemitter"
)

func main() {
	bus := eventemitter.New()

	testListener := func(args ...interface{}) {
		fmt.Println(args[0])
	}

	bus.AddListener("test", testListener)

	bus.EmitSync("test", "Hellooo world")
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Found an Issue?
Create a bug report!

## License
[MIT](https://choosealicense.com/licenses/mit/)
