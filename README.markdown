About This Fork
=======

Seelog is a powerful and easy-to-learn logging framework that provides functionality for flexible dispatching, filtering, and formatting log messages.
It is natively written in the [Go](http://golang.org/) programming language.
The original source tree can be found at (https://github.com/cihub/seelog).

This fork inhances original seelog with a wrapper to simulate the golang's default log behavior.

Quick-start
-----------

```go
package main

import log "github.com/cihub/seelog/seelogWrapper"

func main() {
    log.Panic("Hello, and, Panic!")
}
```
