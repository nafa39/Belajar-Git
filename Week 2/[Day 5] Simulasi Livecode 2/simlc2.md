## Rules
- Untuk kampus remote: WAJIB melakukan share screen (DESKTOP/ENTIRE SCREEN), open cam dan unmute microphone ketika Live Code berjalan (tidak melakukan share screen/salah screen atau yang lainnya akan diingatkan).
- Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
- Kerjakan di openclass pada task selain RULES.
- Waktu pengerjaan: 90 menit untuk 3 soal.
- Boleh mencoba code di text editor lain seperti vscode, sublime, dsb.
- Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
- Dilarang membuka referensi yang berasal dari Hacktiv8 seperti repository di organisasi tugas, challenge openclass yang sudah pernah diberikan, slide show lecture, video recording, lecture, dsb.
- Lakukan submit apabila telah yakin selesai atau ketika waktu telah habis dan diperintahkan instruktur untuk submit.
- Apabila ketika submit masih ada yang salah, selama waktu belum habis maka student boleh merubah code dan melakukan re-submit terakhir.
- Berhati-hati ketika melakukan re-submit, pastikan semua sudah sesuai, karena setelah re-submit student tidak dapat melakukan submit lagi.
- Penilaian berbasis logika dan hasil akhir. Pastikan keduanya sudah benar.

### Notes
Built-in functions yang tidak dilarang oleh RULES:
- Basic primitive data types: `Number()`, `String()`, `Boolean()`
- Informatives: `typeof`, `.length`, `isInteger()`, `isArray()`, `isNaN()`
- Maths: `Math.floor()`, `Math.round()`, `Math.ceil()`, `Math.abs()`, `Math.sqrt()`, `Math.random()`
- Strings: `toString()`, `toUpperCase()`, `toLowerCase()`, `parseInt()`

## Score 
- Soal 1 => 21 point
- Soal 2 => 33 point
- Soal 3 => 46 point

### Reminder:
- Live Code 1 = 10%
- Live Code 2 = 15%
- Live Code 3 = 25%
- Daily challenges = 30%
- Milestones = 20%
- Passing score phase 0 = 61%

---

## Soal 1 - DISNEY ISLAND

### Restrictions
- Built-in function yang dibolehkan mengikuti aturan umum saat briefing.

### Objectives
- Mampu menyelesaikan masalah yang diberikan.
- Mengerti konsep dan cara penggunaan conditional dan data primitive.

### Directions
Buatlah program mesin non-tunai sederhana untuk mensimulasikan transaksi tiket masuk ke wahana bermain di DISNEY ISLAND. Berikut adalah daftar harga tiket masuk ke wahana bermain.

**Wahana Utara**:
- Anak-anak: Rp 85.000
- Dewasa: Rp 125.000
- Lansia: Rp 125.000

**Wahana Selatan**:
- Anak-anak: Rp 143.000
- Dewasa: Rp 165.000
- Lansia: Rp 165.000

Dimana kategori umur dapat dikategorikan sebagai berikut:
- Anak-anak: Usia 2 s/d 12 tahun
- Dewasa: Usia 13 s/d 49 tahun
- Lansia: 50 tahun ke atas

**Validasi**:
- Sebagai contoh, jika pengunjung berumur 28 tahun dan masuk ke 'Wahana Utara' maka tarifnya Rp 125.000.
- Semua yang tidak ada di dalam tabel dianggap tidak sesuai.
- Apabila pengunjung berumur 1 tahun ke bawah maka tampilkan pesan "Dilarang masuk".

**Output yang diinginkan** terlihat di console adalah sebagai berikut:
- Jika wahana dan kategori usia sesuai dan saldo cukup: 
  "Sisa saldo anda adalah RP `<sisa_saldo>`,00. Selamat bermain."
- Jika wahana dan kategori usia sesuai dan saldo tidak cukup: 
  "Saldo anda kurang RP `<selisih_saldo_dan_tarif>`,00. Tidak cukup untuk membeli tiket."
- Jika wahana atau kategori usia tidak sesuai:
  "Tiket tidak ditemukan!"
- Jika usia 1 tahun ke bawah:
  "Dilarang Masuk"

### Notes
- Jangan mengubah penamaan variabel yang telah diberikan (`usia`, `wahana`, `saldo`, `tarif`).
- Untuk variabel `tarif` nilainya jangan diisi (assign) langsung. Gunakan program untuk mengisi ini!

### Examples

#### Contoh 1
**Input**:
```javascript
let wahana = "Wahana Utara"
let usia = 28
let saldo = 180000
let tarif
```
**Output**:
```cli
"Sisa saldo anda adalah Rp 55000,00. Selamat bermain."
```

#### Contoh 2
**Input**:
```javascript
let wahana = "Wahana Selatan"
let usia = 8
let saldo = 140000
let tarif
```
**Output**:
```cli
"Saldo anda kurang Rp 3000,00. Tidak cukup untuk membeli tiket."
```

#### Contoh 3
**Input**:
```javascript
let wahana = "Wahana Barat"
let usia = 28
let saldo = 180000
let tarif
```
**Output**:
```cli
"Tarif tidak ditemukan!"
```

#### Contoh 4
**Input**:
```javascript
let wahana = "Wahana Utara"
let usia = 1
let saldo = 180000
let tarif
```
**Output**:
```cli
"Dilarang Masuk"
```

## Soal 2 - AYO JOGET

### Details

Terdapat permainan ayo joget, dimana user memasukkan input berupa arah sesuai dengan exercise yang ada. Untuk setiap arah pada userInput yang sesuai dengan exercise, maka score akan bertambah 10. Program juga akan menampilkan kategori berdasarkan persentase:

- 100% - Perfect
- 80% - 99% - Great
- 60% - 79% - Good
- 0% - 59% - Bad

## Notes
Persentase didapat dari poin dibagi total poin yang bisa didapat dikali 100, pembulatan ke bawah.
Jika panjang variabel exercise dan userInput berbeda maka tampilkan "Input yang anda masukkan tidak lengkap!"


### Examples

#### Contoh 1
**Input**:
```javascript
var exercise = '<>^v>>>'
var userInput = '<>^^>><'
```
**Output**:
```cli
Anda mendapatkan score 50 / 70. Persentase: 71%, Kategori : Good
```

#### Proses

- Kita akan melakukan compare terhadap setiap karakter
- < adalah karakter pertama di variable exercise akan di compare dengan < karakter pertama di variable userInput.
Karena keduanya adalah arah yang sama maka user mendapatkan 10 point
- > adalah karakter kedua di variable exercise akan di compare dengan > karakter kedua di variable userInput.
Karena keduanya adalah arah yang sama maka user mendapatkan 10 point
- ^ adalah karakter ketiga di variable exercise akan di compare dengan ^ karakter ketiga di variable userInput.
Karena keduanya adalah arah yang sama maka user mendapatkan 10 point
- V adalah karakter keempat di variable exercise akan di compare dengan ^ karakter kempat di variable userInput.
Karena keduanya adalah arah yang tidak sama maka user tidak mendapatkan point.
- Setelah di check keseluruhan inputnya, didapatkan 2 output yang tidak sesuai, dan 5 output yang sesuai.
- Oleh karena itu output yang diharapkan adalah Anda mendapatkan score 50 / 70. Persentase: 71%, Kategori : Good
- 50/70 didapatkan karena kita mendapatkan 50 point dari 70 point yang tersedia.
- 71% adalah persentase kesamaan dari setiap karakter pada variable userInput dan exercise ( 50 / 70 * 100 %)
- Good didapatkan sesuai kategori persentase yang didapatkan.

##### RULES:
- Hanya boleh menggunakan built in function yang diperbolehkan pada intruksi diatas.

**Test Case**:
```javascript
let exercise = '<>^v>>>'
let userInput = '<>^^>><'
// Anda mendapatkan score 50 / 70. Persentase: 71%, Kategori : Good

let exercise = '<>^v'
let userInput = '<>^v'
// Anda mendapatkan score 40 / 40. Persentase: 100%, Kategori : Perfect

var exercise = '<>^v'
var userInput = '<'
// Input yang anda masukkan tidak lengkap!

```


## Soal 3 - Asterix in the Box

### Details
Buatlah sebuah program dengan satu buah variable row yang akan mewakilkan jumlah baris yang akan kita buat. Setiap baris yang dibuat akan memiliki 5 buah col.
Buatlah sebuah output '#' dengan jumlah maksimum yang bisa di capai dari informasi row dan col yang diberikan.

Tambah kan 1 buak asterix ('*') sesuai dengan koordinat yang berada di variabel coordinate dimana:

- karakter pertama menunjukkan baris (row)
- karakter kedua menunjukkann kolom (col)

**Contoh**:
```javascript
let row = 3
let coordinate = "23"
// Output : 
// # # # # #
// # # * # #  (baris ke 2 , kolom ke 3)
// # # # # #

let row = 5
let coordinate = "45"
// Output : 
// # # # # #
// # # # # #  
// # # # # #
// # # # # *  (baris ke 4, kolom ke 5)
// # # # # #

let row = 1
let coordinate = "11"
// output
// * # # # #
```

##### RULES:

- Tidak diperbolehkan menggunakan built-in function:
.map .filter .reduce .split .join .indexOf .findIndex .substrin

##### NOTES

- Diasumsikan input koordinat selalu benar


Selamat Mengerjakan.. Good Luck

