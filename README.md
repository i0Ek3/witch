<h1 align="center">witch</h1>
<p>
  <a href="https://www.npmjs.com/package/witch" target="_blank">
    <img alt="Version" src="https://img.shields.io/npm/v/witch.svg">
  </a>
</p>

> Operate variable by Go reflect.

You can change the value of given variable easily.

BUT DO NOT USE THIS PACKAGE IN YOUR PRODUCTION ENVIRONMENT, YOUR OWN RISK!


## Getting Started

### Install 

`$ go get https://github.com/i0Ek3/witch`

### Usage

```Go
package main

import (
    "github.com/i0Ek3/witch"
)

func main() {
    // set variable x
    witch.Piu(x, v, "ptr"[, "set"])
    // cancel changes
    witch.Xiu()
}
```

## How

![](https://github.com/i0Ek3/witch/blob/master/witch.jpg)


## TODO

- [ ] improve Piu() and implement Xiu()
- [ ] refactor witch_test.go
- [ ] bugfix(later...)


## Show your support

Give a ⭐️ if this project helped you!
