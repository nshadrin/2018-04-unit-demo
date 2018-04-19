// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
	"math"
	"math/rand"
    "net/http"
    "nginx/unit"
)
func pi(n int) float64 {
    ch := make(chan float64)
    for k := 0; k <= n; k++ {
        go term(ch, float64(k))
    }
    f := 0.0
    for k := 0; k <= n; k++ {
        f += <-ch
    }
    return f
}

func term(ch chan float64, k float64) {
    ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("X-Content-Length", "256");

    var b [10]byte
    var res int = 1

    for res > 0 {
        res, _ = r.Body.Read(b[:])

        if res > 0 {
            fmt.Fprintf(w, "body portion: '%s'\n", b[:res]);
        }
    }
	var n = rand.Intn(400)
    fmt.Fprintf(w, "<head><meta http-equiv=\"refresh\" content=1><body>Go Go Go. %v digits of Pi are %v\n", n, pi(n))
}

func main() {
    http.HandleFunc("/", handler)
    //http.ListenAndServe(":8080", nil)
    unit.ListenAndServe(":8080", nil)
}
