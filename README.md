# GoCarRegistrationAPI
A wrapper for the Car Registration API in Go. To use this package you will need a username and password from
http://www.vehicleregistrationapi.com 

## Usage
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

The methods of car_registration are Australia_lookup, USA_lookup and European_lookup, all methods reuturn an interface{},
which can be cast to a string map as shown above.

### Australia_lookup

Australia lookup accepts four parameters, the registration number, the state, and your username and password, in that order.

### USA_lookup

USA lookup accepts four parameters, the registration number, the state, and your username and password, in that order.

### European_lookup

European lookup accepts four parameters, the endpoint, the registration number, and your username and password, in that order,
where the endpoint can be one of:
* Check (UK)
* CheckBelgium 
* CheckCroatia 
* CheckCzechRepublic 
* CheckDenmark 
* CheckEstonia 
* CheckFinland 
* CheckFrance 
* CheckHungary 
* CheckIndia 
* CheckIreland 
* CheckItaly 
* CheckNetherlands 
* CheckNewZealand 
* CheckNigeria 
* CheckNorway 
* CheckPortugal 
* CheckRussia 
* CheckSlovakia 
* CheckSouthAfrica 
* CheckSpain 
* CheckSriLanka 
* CheckSweden 
* CheckUAE 

