package resolve_test

import (
    "fmt"
    "syscall"
    "github.com/arcpop/resolve/resolve"
)


func ExampleGetAddrInfo(){
    as, e := resolve.GetAddrInfo("localhost", 80, syscall.AF_INET)
    fmt.Println(as)
    fmt.Println(e)
    as, e = resolve.GetAddrInfo("localhost", 22, syscall.AF_INET6)
    fmt.Println(as)
    fmt.Println(e)
    // Output:
    // &{80 [127 0 0 1] {0 0 [0 0 0 0] [0 0 0 0 0 0 0 0]}}
    // <nil>
    // &{22 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] {0 0 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 0}}
    // <nil>
}