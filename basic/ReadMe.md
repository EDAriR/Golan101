檔案名 xxx.go
測試檔案命名 xxx_test.go

執行 go run xxx.go 
測試 go test
运行列表中指定的测试用例： go test -run TestArea/Rectangle


table driven test
https://github.com/golang/go/wiki/TableDrivenTests

当你创建一个值，例如一个 wallet，它就会被存储在内存的某处。你可以用 &myval 找到那块内存的地址。

    fmt.Println("address of balance in test is", &wallet.balance)    
    fmt.Println("address of balance in Deposit is", &w.balance)


