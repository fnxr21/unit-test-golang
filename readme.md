# Command

## basic comand

```cmd
go test
```

untuk runing di file root untuk check ada ngk unistesting

```cmd
go test -v
```

untuk runing di file root untuk check ada ngk unistesting dengan detail

```cmd
go test -v -run TestHelloIvan
```

untuk runing di file root untuk check ada ngk unistesting spesifik function

```cmd
go test -v ./...
```

untuk runing di file root untuk check ada ngk unistesting spesifik function

### note

panic disini tidak direkomendasikan dalam testing

_jangan gunakan panic_ dalam testing

============

Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test . namun diakhir ketika selesai,maka unit test tersebut dianggap gagal

FailNow() akan menggagalkan unit test saat ini juga ,tanpa menlanjutkan eksekusi unit test

```go
func Testng(t *testing.T) {
    Fail()
    FailNow() // akan menggagalkan unit test saat ini juga ,tanpa menlanjutkan eksekusi unit test
    Error() //
    Fatal() //
}
```
