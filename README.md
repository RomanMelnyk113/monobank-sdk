# MonoBank SDK
Simple MonoBank wrapper over public API


## Example

```go
package main

import (
    "fmt"
    "os"
    "time"
    "github.com/RomanMelnyk113/monobank-sdk"

    client := monobank.NewClient("token")
    user, err := client.GetUserInfo()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    fmt.Println("User details:", user)
    from := time.Now().Add(-24 * time.Hour)
    to := time.Now()

    transactions, err := client.GetTransactions(user.Accounts[0].AccountID, from, to)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    fmt.Println("Transactions:", transactions)
)
```

