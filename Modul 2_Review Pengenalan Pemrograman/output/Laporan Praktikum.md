# <h1 align="center">Laporan Praktikum Modul 2 <br> Review Pengenalan Pemrograman </h1>
___
<p align="center">NAFILA SETYANI - 103112430019</p>
___
## Dasar Teori
___
Pemrograman Go (Golang) adalah bahasa pemrograman open-source yang dikembangkan oleh Google dengan fokus pada kesederhanaan, efisiensi, dan kecepatan eksekusi. Go memiliki sintaks yang ringkas dan mudah dibaca, mendukung konkurensi (goroutines) untuk pemrosesan paralel, serta memiliki sistem manajemen memori yang otomatis (garbage collection). Go juga dikenal dengan kecepatan kompilasi yang tinggi, tipe data yang kuat, serta pustaka standar yang kaya untuk membangun aplikasi backend, layanan web, hingga sistem terdistribusi dengan performa tinggi.
## Unguided
___
### Soal Latihan 2A

##### Soal 1
>Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

```go
package main
import "fmt"
func main() {
    var (
        satu, dua, tiga string
        temp            string
    )
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&satu)
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&dua)
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&tiga)
    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp
    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```

> Output
> ![Screenshot bagian x](output/2a_soal1.png)
> 

Program di atas bertujuan untuk menerima tiga input string dari pengguna, kemudian menampilkan urutan awal dan hasil pertukaran nilai antar variabel. Pertama, program meminta tiga input string (`satu`, `dua`, `tiga`) dari pengguna menggunakan `fmt.Scanln()`. Setelah itu, program mencetak urutan awal dari ketiga string tersebut. Selanjutnya, proses pertukaran dilakukan dengan menggunakan variabel sementara (`temp`) untuk menyimpan nilai `satu`, lalu `satu` diisi dengan nilai `dua`, `dua` diisi dengan nilai `tiga`, dan `tiga` diisi dengan nilai awal `satu` yang disimpan di `temp`. Terakhir, program mencetak hasil pertukaran string dalam format yang baru.
___
##### Soal 2
 >Tahun kabisat adalah tahun yang habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100. Buatlah sebuah program yang menerima input sebuah bilangan bulat dan memeriksa apakah bilangan tersebut merupakan tahun kabisat (true) atau bukan (false).
 
```go
package main
import "fmt"
func isKabisat(year int) bool {

    return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}
func main() {
    var tahun int
    fmt.Print("Tahun: ")
    fmt.Scanln(&tahun)
    fmt.Println("Kabisat:", isKabisat(tahun))

}
```

>Output![Screenshot bagian x](output/2a_soal2.png)
> 

Program di atas digunakan untuk menentukan apakah suatu tahun merupakan tahun kabisat atau tidak. Fungsi `isKabisat(year int) bool` akan mengembalikan `true` jika tahun memenuhi salah satu dari dua kondisi berikut: habis dibagi 400, atau habis dibagi 4 tetapi tidak habis dibagi 100. Dalam fungsi `main()`, program meminta input tahun dari pengguna, lalu memanggil fungsi `isKabisat()` untuk mengevaluasi apakah tahun tersebut kabisat. Hasil evaluasi kemudian ditampilkan dengan mencetak `true` jika kabisat dan `false` jika bukan.
___
##### Soal 3
>Buat program Bola yang menerima input jari-jari suatu bola (bilangan bulat). Tampilkan Volume dan Luas kulit bola. 𝑣𝑜𝑙𝑢𝑚𝑒𝑏𝑜𝑙𝑎 = 4 3 𝜋𝑟 3 dan 𝑙𝑢𝑎𝑠𝑏𝑜𝑙𝑎 = 4𝜋𝑟 2 (π ≈ 3.1415926535).

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    var r float64
    fmt.Print("Jejari = ")
    fmt.Scanln(&r)
    volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
    luas := 4 * math.Pi * math.Pow(r, 2)
    fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
```

>Output
>![Screenshot bagian x](output/2a_soal3.png)

Program di atas menghitung volume dan luas permukaan bola berdasarkan jejari (radius) yang dimasukkan oleh pengguna. Setelah menerima input jejari `r`, program menghitung volume bola menggunakan rumus V=43πr3V = \frac{4}{3} \pi r^3 dan luas permukaannya menggunakan rumus A=4πr2A = 4 \pi r^2, dengan bantuan fungsi `math.Pi` dan `math.Pow()` dari paket `math`. Hasil perhitungan kemudian ditampilkan dalam format desimal dengan empat angka di belakang koma menggunakan `fmt.Printf()`.
___
##### Soal 4
>Dibaca nilai temperatur dalam derajat Celsius. Nyatakan temperatur tersebut dalam Fahrenheit 𝐶𝑒𝑙𝑠𝑖𝑢𝑠 = (𝐹𝑎ℎ𝑟𝑒𝑛ℎ𝑒𝑖𝑡 − 32) × 5/9 𝑅𝑒𝑎𝑚𝑢𝑟 = 𝐶𝑒𝑙𝑐𝑖𝑢𝑠 × 4/5 𝐾𝑒𝑙𝑣𝑖𝑛 = (𝐹𝑎ℎ𝑟𝑒𝑛ℎ𝑒𝑖𝑡 + 459.67) × 5/9

```go
package main
import "fmt"
func main() {
    var celsius float64
    fmt.Print("Temperatur Celsius: ")
    fmt.Scanln(&celsius)
    fahrenheit := (celsius * 9 / 5) + 32
    reamur := celsius * 4 / 5
    kelvin := celsius + 273.15
    fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
    fmt.Printf("Derajat Reamur: %.0f\n", reamur)
    fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
```

>Output
>![Screenshot bagian x](output/2a_soal4.png)

Program di atas mengonversi suhu dari derajat Celsius ke tiga skala suhu lainnya: Fahrenheit, Reamur, dan Kelvin. Pengguna diminta memasukkan suhu dalam Celsius, kemudian program menghitung konversinya menggunakan rumus: Fahrenheit = 95×\frac{9}{5} \times Celsius + 32, Reamur = 45×\frac{4}{5} \times Celsius, dan Kelvin = Celsius + 273.15. Hasil konversi ditampilkan dengan format tanpa desimal menggunakan `fmt.Printf()`.
___
##### Soal 5
>Tipe karakter sebenarnya hanya apa yang tampak dalam tampilan. Di dalamnya tersimpan dalam bentuk biner 8 bit (byte) atau 32 bit (rune) saja. Buat program ASCII yang akan membaca 5 buat data integer dan mencetaknya dalam format karakter. Kemudian membaca 3 buah data karakter dan mencetak 3 buah karakter setelah karakter tersebut (menurut tabel ASCII) Masukan terdiri dari dua baris. Baris pertama berisi 5 buah data integer. Data integer mempunyai nilai antara 32 s.d. 127. Baris kedua berisi 3 buah karakter yang berdampingan satu dengan yang lain (tanpa dipisahkan spasi). Keluaran juga terdiri dari dua baris. Baris pertama berisi 5 buah representasi karakter dari data yang diberikan, yang berdampingan satu dengan lain, tanpa dipisahkan spasi. Baris kedua berisi 3 buah karakter (juga tidak dipisahkan oleh spasi).

```go
package main
import "fmt"
func main() {
    var nums [5]int
    var input string

    // Membaca 5 angka integer

    fmt.Print("Masukkan 5 angka ASCII: ")
    for i := 0; i < 5; i++ {
        fmt.Scan(&nums[i])
    }
    // Membaca 3 karakter langsung sebagai string
    fmt.Print("Masukkan 3 karakter tanpa spasi: ")
    fmt.Scan(&input)
    // Konversi dan cetak 5 angka ke karakter
    for i := 0; i < 5; i++ {
        fmt.Printf("%c", nums[i])
    }
    fmt.Println()
    // Cetak 3 karakter berikutnya
    fmt.Println(input)

}
```

>Output
>![Screenshot bagian x](output/2a_soal5.png)

Program di atas membaca lima angka ASCII dan tiga karakter langsung dari input pengguna, kemudian mencetak hasil konversi angka ASCII menjadi karakter serta mencetak tiga karakter yang dimasukkan. Pertama, program meminta pengguna memasukkan lima angka yang mewakili kode ASCII, lalu menyimpannya dalam array `nums`. Setelah itu, program meminta tiga karakter tanpa spasi dan menyimpannya dalam variabel `input`. Dalam tahap output, program mengonversi angka dalam array `nums` menjadi karakter menggunakan `printf("%c")` dan mencetaknya, diikuti dengan pencetakan langsung tiga karakter yang telah dimasukkan.
___



### Soal Latihan 2B

##### Soal 1
>Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.

```go
package main
import (
    "fmt"
)
func main() {

    // Urutan warna yang diharapkan

    expected := [4]string{"merah", "kuning", "hijau", "ungu"}
    success := true
    // Loop untuk 5 kali percobaan

    for i := 1; i <= 5; i++ {
        var input [4]string
        fmt.Printf("Percobaan %d: ", i)
        for j := 0; j < 4; j++ {

            fmt.Scan(&input[j])
        }
        // Periksa apakah input sesuai dengan urutan warna yang diharapkan
        if input != expected {
            success = false
        }
    }

    // Cetak hasil akhir
    fmt.Println("BERHASIL:", success)
}
```

>Output
>![Screenshot bagian x](output/2b_soal1.png)

Program di atas memeriksa apakah pengguna dapat memasukkan urutan warna yang benar dalam lima kali percobaan. Program memiliki array `expected` yang berisi urutan warna yang diharapkan: _merah, kuning, hijau, ungu_. Pada setiap percobaan, pengguna diminta memasukkan empat warna, yang disimpan dalam array `input`. Setelah menerima input, program membandingkan array `input` dengan `expected`. Jika ada satu percobaan yang tidak sesuai, variabel `success` diubah menjadi `false`. Setelah lima percobaan selesai, program mencetak hasil akhir dengan menampilkan _BERHASIL: true_ jika semua percobaan benar, atau _BERHASIL: false_ jika ada kesalahan dalam urutan warna.
___

##### Soal 2
>Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan ‘– ‘, contoh pita diilustrasikan seperti berikut ini. Pita: mawar – melati – tulip – teratai – kamboja – anggrek Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita. (Petunjuk: gunakan operasi penggabungan string dengan operator “+” ). Tampilkan isi pita setelah proses input selesai.

```go
package main
import (
    "fmt"
    "strings"
)
func main() {
    var n int
    var flower, ribbon string
    var count int
    fmt.Print("N: ")
    fmt.Scan(&n)
    // Mode pertama: Input berdasarkan jumlah N
    if n > 0 {
        for i := 1; i <= n; i++ {
            fmt.Printf("Bunga %d: ", i)
            fmt.Scan(&flower)
            ribbon += flower + " - "
            count++
        }
        // Hapus tanda " - " di akhir string
        ribbon = strings.TrimSuffix(ribbon, " - ")
    }
    // Mode kedua: Input hingga "SELESAI"
    if n == 0 {
        fmt.Println("Gunakan 'SELESAI' untuk menghentikan input.")
        for {
            fmt.Printf("Bunga %d: ", count+1)
            fmt.Scan(&flower)
            if flower == "SELESAI" {
                break
            }
            ribbon += flower + " - "
            count++
        }
        // Hapus tanda " - " di akhir string jika ada input bunga
        ribbon = strings.TrimSuffix(ribbon, " - ")
    }

    // Output hasil pita
    fmt.Println("Pita:", ribbon)
    fmt.Println("Bunga:", count)

}

```

>Output
>![Screenshot bagian x](output/2b_soal2.png)

Program di atas berfungsi untuk membuat pita berisi nama-nama bunga yang dimasukkan oleh pengguna, dengan dua mode input. Pada mode pertama, jika pengguna memasukkan angka `N` lebih dari 0, program akan meminta `N` nama bunga dan menyusunnya dalam satu string yang dipisahkan dengan tanda " - ". Pada mode kedua, jika pengguna memasukkan `N = 0`, program akan terus meminta input bunga sampai pengguna mengetik "SELESAI". Setelah semua bunga dikumpulkan, tanda " - " di akhir string akan dihapus menggunakan `strings.TrimSuffix()`, lalu program mencetak pita hasil rangkaian bunga serta jumlah bunga yang dimasukkan.
___

##### Soal 3
>Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg. Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    var left, right float64
    var totalWeight float64
    for {
        fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
        fmt.Scan(&left, &right)
        // Jika salah satu input negatif, hentikan program
        if left < 0 || right < 0 {
            fmt.Println("Proses selesai.")
            break
        }
        // Hitung total berat dan selisih
        totalWeight = left + right
        diff := math.Abs(left - right)
        // Cek jika total berat melebihi 150 kg
        if totalWeight > 150 {
            fmt.Println("Proses selesai.")
            break
        }
        // Cek apakah sepeda oleng
        oleng := diff > 9
        // Tampilkan hasil
        fmt.Println("Sepeda motor Pak Andi akan oleng:", oleng)

    }

}
```

>Output
>![](output/2b_soal3.png)

Program ini berfungsi untuk mengevaluasi keseimbangan sepeda motor Pak Andi berdasarkan berat belanjaan di dua kantong. Pengguna memasukkan dua angka yang mewakili berat belanjaan di masing-masing kantong. Jika salah satu angka negatif atau total berat melebihi 150 kg, program akan berhenti. Selisih berat antara kedua kantong dihitung menggunakan `math.Abs()`, dan jika selisihnya lebih dari 9 kg, sepeda dianggap oleng. Hasil akhir akan menampilkan apakah sepeda oleng atau tidak berdasarkan perbedaan berat kantong.
___

##### Soal 4
>Diberikan sebuah persamaan sebagai berikut ini. 𝑓(𝑘) = (4𝑘 + 2) 2 (4𝑘 + 1)(4𝑘 + 3) Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas.

```go
package main
import (
    "fmt"
    "math"
)
// Fungsi untuk menghitung nilai f(K) sesuai dengan rumus
func f(k int) float64 {
    return (math.Pow(float64(4*k+2), 2)) / (float64(4*k+1) * float64(4*k+3)
}
// Fungsi untuk menghitung aproksimasi akar 2 berdasarkan rumus produk tak hingga
func sqrtApproximation(K int) float64 {
    result := 1.0
    for k := 0; k < K; k++ {
        result *= f(k)
    }
    return result
}
func main() {
    var K int
    // Meminta input dari pengguna
    fmt.Print("Masukkan nilai K: ")
    fmt.Scan(&K)
    // Menghitung f(K)
    fk := f(K)
    fmt.Printf("Nilai f(K) = %.10f\n", fk)
    // Menghitung aproksimasi akar 2 berdasarkan K
    approxSqrt2 := sqrtApproximation(K)
    fmt.Printf("Nilai akar 2 = %.10f\n", approxSqrt2)
}
```

>Output
>![](output/2b_soal4.png)

Program ini menghitung nilai fungsi f(K)f(K) dan aproksimasi akar 2 menggunakan rumus produk tak hingga. Fungsi `f(k)` menghitung nilai sesuai rumus f(K)=(4K+2)2(4K+1)(4K+3)f(K) = \frac{(4K+2)^2}{(4K+1)(4K+3)}. Fungsi `sqrtApproximation(K)` menggunakan perhitungan produk berulang dari fungsi f(k)f(k) untuk mendekati nilai akar 2. Pengguna memasukkan nilai KK, lalu program menampilkan hasil perhitungan f(K)f(K) dengan presisi 10 desimal, serta aproksimasi akar 2 yang semakin mendekati nilai sebenarnya seiring bertambahnya KK.
___

### Soal Latihan 2C

##### Soal 1
>PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

```go
package main
import "fmt"
func main() {
    var berat, kg, gram, harga int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scanln(&berat)
    kg = berat / 1000
    gram = berat % 1000
    fmt.Println("Detail berat:", kg, "kg", "+", gram, "gram")
    if gram >= 500 {
        gram = gram * 5
    } else {
        gram = gram * 15
    }
    if kg < 10 {
        kg = kg * 10000
        harga = kg + gram
    } else {
        kg = kg * 10000
        harga = kg
    }
    fmt.Println("Detail biaya: Rp", kg, "+ Rp", gram)
    fmt.Println("Total biaya: Rp", harga)
}
```

>Output
>![](output/2c_soal1.png)

Program ini menghitung biaya pengiriman parsel berdasarkan berat dalam gram. Pengguna memasukkan berat total, lalu program menghitung jumlah kilogram dan sisa gram menggunakan operasi pembagian dan modulus. Jika sisa gram ≥ 500, biaya gram dihitung dengan tarif Rp 5 per gram, sedangkan jika < 500, tarifnya Rp 15 per gram. Untuk berat kurang dari 10 kg, biaya dihitung dengan tarif Rp 10.000 per kg ditambah biaya gram, sedangkan untuk 10 kg atau lebih, hanya dihitung berdasarkan tarif kilogram tanpa tambahan biaya gram. Akhirnya, program menampilkan rincian berat serta total biaya pengiriman.
___

##### Soal 2
>Jawablah pertanyaan-pertanyaan berikut: a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal? b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya! c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.

```go
package main
import "fmt"
func main() {
    var nam float64
    var nmk string
    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scanln(&nam)
    if nam > 80 {
        nmk = "A"
    } else if nam > 72.5 && nam <= 80 {
        nmk = "AB"
    } else if nam > 65 && nam <= 72.5 {
        nmk = "B"
    } else if nam > 57.5 && nam <= 65 {
        nmk = "BC"
    } else if nam > 50 && nam <= 57.5 {
        nmk = "C"
    } else if nam > 40 && nam <= 50 {
        nmk = "D"
    } else if nam <= 40 {
        nmk = "E"
    }
    fmt.Println("Nilai mata kuliah:", nmk)
}
```

>Output
>![](output/2c_soal2.png)

Program ini mengonversi nilai akhir mata kuliah ke dalam bentuk huruf berdasarkan rentang tertentu. Pengguna memasukkan nilai akhir, lalu program mengevaluasi rentang nilai tersebut menggunakan struktur **if-else**. Jika nilai lebih dari 80, diberikan huruf **A**, sedangkan nilai di bawahnya dikategorikan menjadi **AB, B, BC, C, D, atau E** sesuai batas yang ditentukan. Hasil akhirnya adalah nilai dalam bentuk huruf yang ditampilkan ke layar.
___

##### Soal 3
>Sebuah bilangan bulat b memiliki faktor bilangan f > 0 jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2. Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut! Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri. Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.

```go
package main
import (
    "fmt"
)
func main() {
    var b int
    // Input bilangan
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&b)
    // Pastikan bilangan lebih dari 1
    if b <= 1 {
        fmt.Println("Bilangan harus lebih dari 1.")
        return
    }
    // Mencari faktor bilangan

    fmt.Print("Faktor: ")
    var faktor []int
    for i := 1; i <= b; i++ {
        if b%i == 0 {
            faktor = append(faktor, i)
        }
    }
    // Menampilkan faktor
    for _, f := range faktor {
        fmt.Print(f, " ")
    }
    fmt.Println()

    // Menentukan apakah bilangan prima
    isPrime := len(faktor) == 2
    fmt.Println("Prima:", isPrime)
}
```

>Output
>![](output/2c_soal3.png)

Program ini menentukan faktor dari suatu bilangan dan mengecek apakah bilangan tersebut adalah bilangan prima. Pengguna diminta memasukkan bilangan bulat lebih dari 1. Program kemudian mencari faktor-faktornya dengan membagi bilangan tersebut dengan angka dari 1 hingga bilangan itu sendiri. Faktor yang ditemukan disimpan dalam slice dan ditampilkan. Jika bilangan memiliki tepat dua faktor (1 dan bilangan itu sendiri), maka bilangan tersebut dikategorikan sebagai bilangan prima. Hasil akhir menampilkan daftar faktor dan status apakah bilangan tersebut prima atau tidak.
___
