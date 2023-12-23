package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()
    var a,b,c int
    fmt.Fscan(in,&a,&b,&c)
    if (a==7&&b==5&&c==5)||(a==5&&b==7&&c==5)||(a==5&&b==5&&c==7) {
        fmt.Fprintln(out,"YES")
    } else {
        fmt.Fprintln(out,"NO")
    }
}
