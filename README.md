# sdk-go

## Message

It's easy to send message to any number using the SDK, all you have to do is the following.

```
import (
    "fmt"
    "github.com/zendlk/go-sdk"
)

Zend := zend.NewClient("{_token_}", "{_masking_}")
message, err := Zend.Message("+94777123456", "Message text here")
if err != nil { fmt.Println(err) }
fmt.Println( message )
```
