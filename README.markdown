About This Fork
=======

This fork inhances original seelog with a wrapper to simulate the golang's default log behavior.

The original source tree can be found at here(https://github.com/cihub/seelog).


Quick-start
-----------

```go
package main

import log "github.com/cihub/seelog/seelogWrapper"

func main() {
    log.Panic("Hello, and, Panic!")
}
```
