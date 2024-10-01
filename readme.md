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
    FailNow()
    Error()
    Fatal()
}
```

## Assertions

disini kita menggunakan testify untuk handle if else atau condition

mereka ada 2 jenis assertion dan require sama seperti fail() dan failnow()

## Skip

skip ini hanya dalam keadaan tertentu

contoh perbedaan env atau os windods dan mac os

## before and after test

kita mengunkan testing.M yang dimana artinya main not before or after cuma kita mengunakan salah satu fitur nya saja

agak strict karna penuliosan func hbarus TestMain dan hanyua di eksekusi sekali ketika memulai dan tidak per function

## SubTest

membuat test didalam function testing
