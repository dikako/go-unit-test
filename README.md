# Unit Test

* Unit test akan fakus menguji bagian terkecil dari kode program, biasanya menguji sebuah method
* Unit test biasanya dibuat kecil dan cepat, oleh karena itu biasanya kode unit test lebih banyak dari kode program aslinya, karena semua skenario pengujian akan di coba di unit test
* Unit test bisa digunakan sebagai cara untuk meningkatkan kualitas kode program kita

## Testing package
* Dibahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framwork khusus untuk unit test
* Berbeda dengan Go-Lang, di Go-Lang untuk unit test sudah disediakan sebuah package khusus bernama testing
* Selain untuk menjalankan unit test, di Go-Lang juga sudah disediakan perintahnya
* Hal ini membuat implementasi unit testing di Go-Lang sangat mudah dibanding dengan bahasa  pemrograman lain
* Dokumetasi Go-Lang [package testing](https://golang.org/pkg/testing)

### testing.T 
* Go-Lang menyediakan sebuah struct yang bernama `testing.T`
* Struct ini digunakan untuk unit test di Go-Lang

### testing.M
* `testing.M` adalah struct yang disediakan go-Lang untuk mengatur file cycle testing

### testing.B
* `testing.B` adalah struct yang disediakan Go-Lang untuk melakukan benchmarking
* Di Go-Lang untuk melakukan benchmark (mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library atau framework tambahan

## Aturan File Test
* Go-Lang memiliki aturan cara membuat file khusus untuk unit test
* Untuk membuat file unit test, kita harus menggunakan akhiran `_test`
* Jadi misal kita membuat file `hello_world.go`, artinya untuk membuat unit testnya kita harus membuat file `hello_world_test.go`

## Aturan Function Unit Test
* Selain aturan nama file, di Go-Lang juga sudah diatur untuk nama function unit test
* Nama function untuk unit test harus diawali dengan nama Test
* Misal jika  kita ingin mengetes function HelloWorld, maka kita akan membuat function unit test dengan nama `TestHelloWorld`
* Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penitng adalah harus diawalin dengan kata Test
* Selanjutnya harus memiliki parameter `(t *testing.T)` dan tidak mengembalikan return value

## Cara menjalankan Unit Test
* Untuk menjalankan unit test kita bisa menggunakan perintah: <br>`go test`
* Jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah: <br>`go test -v`
* Jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah: <br>`go test -v -run=TestNamaFunction`<br>Ini akan menjalankan semua prefix yang berawalan `TestNamaFunction`, jika kita memiliki function `TestNamaFunctionDua` ini juga akan ikut di running
* Jika kita ingin menjalankan semua unit test (semua package) dari top folder module nya, kita bisa menggunakan perintah:<br>`go test ./...`

## Cara mengagalkan Unit Test
* Menggagalkan unit test menggunakan `panic` bukanlah hal yang bagus
* Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan `testing.T`
* Terdapat function `Fail()`, `FailNow()`, `Error()` dan `Fatal()` jika kita ingin menggagalkan unit test

### t.Fail() dan t.FailNow()
* Terdapat dua function untuk menggagalkan unit test, yaitu `Fail()` dan `FailNow()`
* `Fail()` akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test selanjutnya. Namun diakhir ketika unit test selesai maka unit test tersebut akan dianggap gagal
* `FailNow()` akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test selanjutnya

### t.Error(args..) dan t.Fatal(args...)
* `Error()` function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function `Fail()`, sehingga mengakibatkan unit test dianggap gagal
* Namun karena hanya memanggil `Fail()`, artinya eksekusi unit test akan tetap berjalan sampai selesai
* `Fatal()` mirip dengan `Error()`, hanya saja setelah melakukan log error, dia akan memanggil `FailNow()`, sehingga mengakibatkan eksekusi unit test berhenti
* Contoh penggunaan: [hello_world_test.go](helper/hello_world_test.go)

## Assertion
* Melakukan pengecekan di unit test secara manual menggunakan `if` sangatlah menyebalkan
* Apalagi jika result data yang harus di cek banyak
* Oleh karena itu, sangat disarankan untuk menggunakan `Assertion` untuk melakukan pengecekan
* Sayangnya di Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini

### Testify
* Salah satu library assertion yang paling populer di Go-Lang adalah `Testify`
* Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
* [Testify Github Repository](https://github.com/stretchr/testify)
* Kita bisa menambahkan `Testify` di Go module kita dengan perintah:<br>`go get github.com/stretchr/testify`
* Contoh penggunaan: [hello_world_testify_test.go](helper/hello_world_testify_test.go)

#### assert vs require
* `Testify` menyediakan dua package untuk assertion, yaitu `assert` dan `require`
* Saat kita menggunakan `assert`, jika pengecekan gagal, maka assert akan memanggil `Fail()`, artinya eksekusi unit test akan tetap dijalankan
* Sedangkan jika kit amenggunakan `require`, jika pengecekan gagal, maka `require` akan memamnggil `FailNow()`, artinya eksekusi unit test tidak akan dilanjutkan

## Skip Test
* Kadang dalam keadaan tertentu kita ingin membatalkan eksekusi unit test
* Di Go-Lang jug akita bisa membatalkan eksekusi unit test jika kita mau
* Untuk membatalkan unit test kit abis amenggunakan function `Skip()`
* Contoh penggunaan: [hello_world_skip_test.go](helper/hello_world_skip_test.go)

## Before dan After Test
* Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
* Jikalau kode yang kita lakukan sebelumnya dan setelah selalu sama antar unit test function, maka membuat manual di unit test function nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
* Untungnya di Go-Lang terdapat feature bernama `testing.M`
* Fitur ini bernama `Main`, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test
* Untuk mengatur eksekusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter `testing.M`
* Jika terdapat function TestMain tersebut, maka secara otomatis Go-Lang akan mngeksekusi function ini tiap kali akan menjalankan unit test di sebuah package
* Dengan ini kita bis amengatur `Before` dan `After` unit test sesuai dengan yang kita mau
* Ingat, function `TestMain` itu dieksekusi hanya sekali per Go-Lang package, buka per tiap function unit test
* Contoh penggunaan: [before_after_test.go](helper/before_after_test.go)

## Sub Test
* Go-lang mendukung fitur pembuatan function unit test di dalam function unit test
* Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman lain
* Untuk membuat sub test, kita bisa menggunakan function `Run()`
* Contoh penggunaan: [sub_test.go](helper/sub_test.go)

### Menjalankan hanya Sub Test atau salah satu Sub Test
* Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perinatah:<br>`go test -v -run=TestNamaFunction`
* Jika kita ingin menjalankan hanya salah satu sub test, kita bisa menggunakan perintah:<br>`go test -v -run=TestNamaFunction/NamaSubTest`

## Table Test
* Jika diperhatikan sebenarnya dengan `Sub Test` kita bisa membuat test secara dinamis
* Dan fitur `Sub Test` ini bisa digunakan oleh programmer Go-Lang untuk membuat test dengan konsep `Table Test`
* `Table Test` yaitu dimana kita menyediakan data berupa `slice` yang berisi parameter dan ekspektasi hasil dari unit test
* Lalu `slice` tersebut kita iterasi menggunakan `Sub Test`
* Contoh penggunaan: [table_test.go](helper/table_test.go)

## Mock
* `Mock` adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program dari awal
* `Mock` adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing
* Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke `third party service` dan belum tentu response nya sesuai dengan apa yang kita mau
* Pada kasus seperti ini, cocok sekali untuk menggunakan mock object

### Testify Mock
* Untuk membuat mock object, tidak ada feature bawaan Go-Lang, namun jita bis amenggunakan library `Testify`
* `Testify` mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
* Namun, perlu diperhatikan jika desain kode program kita jelek akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain program kita dengan baik
* Contoh penguunaan `Mock`: [category_repository_mock.go](repository/category_repository_mock.go)
* Contoh penguunaan `Test Mock`: [category_service_test.go](service/category_service_test.go)

## Benchmark
* Selain unit test, Go-Lang testing package jua mendukung untuk melakukan `benchmark`
* `Benchmark' adalah mekanisme menghitung kecepatan performa kode apliaksi kita
* `Benchmark` di Go-Lang dilakukan secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
* Kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh `testing.B` bawaan dari testing package

### testing.B
* `testing.B` adalah struct yang digunakan untuk melakukan `benchmark`
* `testing.B` mirip dengan `testing.T` terdapat function `Fail()`, `FailNow()`, `Error()`, `Fatal()` dan lain lain
* Yang membedakan adalah ada beberapa attribute dan function tambahan yang digunakan untuk melakukan `benchmark`
* Salah satunya adalah attribute `N` yang digunakan untuk melakukan total iterasi sebuah `benchmark`

### Cara kerja Benchmark
* Cara kerja `benchmark` di Go-Lang sangat sederhana
* Gimana kita hanya perlu membuat perulangan sejumlah `N` attribute
* Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu medeteksi berapa lama proses tersebut berjalan dan disimpulkan performa benchmark nya dalam waktu

### Benchmark Function
* Mirip seperti unit test, untuk `benchmark` pun, di Go-Lang sudah ditentukan nama function nya, harus diawali dengan kata `Benchmark`<br>Contoh: `BenchmarkHelloWorld`, `BenchmarkLogin`
* Selain itu harus memiliki parameter `(b *testing.B)`
* Dan tidak boleh mengembalikan return value
* Untuk nama file `benchmark` sama seperti unit test, diakhiri dengan `_test`<br>Contoh: `hello_world_test.go`, `login_test.go`
* Contoh penggunaan `Benchmark`: [benchmark_test.go](helper/benchmark_test.go)
* Contoh penggunaan `Benchmark Sub Test`: [benchmark_sub_test.go](helper/benchmark_sub_test.go)
* Contoh penggunaan `Benchmark Table Test`: [benchmark_table_test.go](helper/benchmark_table_test.go)

### Cara menjalankan Benchmark
* Untuk mejalankan seluruh `benchmark` di module, kita bis amenggunakan perintah sama seperti test, namun ditambahkan parameter `bench`:<br>`go test -v -bench=.`
* Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah:<br>`go test -v -run=NotMathUnitTest -bench=.` atau<br>`go test -v -run=NotMathUnitTest -bench=BenchmarkFunction` untuk menjalan specific `benchmark`
* Jika kita ingin menjalankan sepcific `Benchmark sub test` bisa gunakan perintah:<br>`go test -v -run=NotMathUnitTest -bench=BenchmarkFunction/BenchmarkSubTestName`
* Jika kita ingin menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah:<br>`go test -v -bench=../...`
