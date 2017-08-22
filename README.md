# GoCarRegistrationAPI
A wrapper for the Car Registration API in Go

###Usage
```go
package main
import "github.com/infiniteloopltd/GoCarRegistrationAPI"
import "fmt"

func main() {  
    data := car_registration.Australia_lookup("YHC14Y","NSW","***Your Username***","***Your Password***")
    m := data.(map[string]interface{})
    //fmt.Printf("%+v", m)
    fmt.Println(m["Description"])
}
```
