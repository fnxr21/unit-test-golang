# Unit Testing

Credit

Folder API

Richard han [github](https://github.com/RichardHanitio)

_others_
eko khanedy kurniawan [linkedin](https://www.linkedin.com/in/khannedy/)

<!--
[link doc](https://docs.google.com/presentation/d/1XxMEaA-JsPHr9BUw2oIOPlEL_psI3EaUFUpuvdlDB_Q/edit#slide=id.gb233370586_0_300) -->

## Agenda

- Pengenalan Software Testing
- testing Package
- Unit Test
- Assertion
- Mock, dan
- Benchmark

## Pengenalan Software Testing

- Software testing adalah salah satu disiplin ilmu dalam software engineering
- Tujuan utama dari software testing adalah memastikan kualitas kode dan aplikasi kita baik
- Ilmu untuk software testing sendiri sangatlah luas, pada materi ini kita hanya akan fokus ke unit testing

## test piramid

![image](https://martinfowler.com/bliki/images/testPyramid/test-pyramid.png)

## Contoh High Level Architecture Aplikasi

### basic comand

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

## tableTest

banyak dev golang makai ini untuk melakukan test

ini semua untuk menghidari testing secara manual atau doubel

## MOCK

golang tidak memiliki mock bawaan jadi kita mengunakan mock testify

## Benchmark

menghitung kecepatan perfoma code aplikasi
