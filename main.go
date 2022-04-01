package main

func main() {
    // Ubah "world" menjadi "Go" dibawah ini
    println("Hello, Go")
}

package main

func main() {
	println("Hello, world")
}
func sub() {
	println("Hello, Go")
}

package main

func main() {
    // Cetak "Hello, world"
    println("Hello, world")
    
    //Baris ini seharusnya merupakan sebuah komentar
    
}

package main

func main() {
    // Cetak angka 7
    println(7)
    
    // Cetak penjumlahan dari 9 dan 3
    println(9+3)
    
    // Cetak "9 + 3" sebagai sebuah string
    println("9+3")
    
}

package main

func main() {
    // Cetak hasil dari 12 dibagi dengan 3
    println(12/3)
    
    // Cetak hasil dari 3 dikali 6
    println(3*6)
    
    // Cetak sisa hasil pembagian dari 8 dibagi dengan 3
    println(8%3)
    
}

package main

func main() {
    // Deklarasikan variable `message` dan tetapkan nilai "Hello, world"
    var message string = "Hello, world"

    // Cetak nilai dari dari variable `message`
    println(message)
    
}

package main

func main() {
	var message string = "Hello, world"

	// Perbarui variable `message` dengan "Hello, Go"
	message = "Hello, Go"
    
	println(message)
}

package main

func main() {
    // Deklarasikan variable `message` dan tetapkan "Hello, world" pada-nya
    message := "Hello, world"
    
    // Deklarasikan variable `number` dan tetapkan `100` pada-nya 
    number := 100
    
    // Cetak nilai dari `message` dan `number`
    println(message,number)
  
}

package main

func main() {
    // Hapus baris `println` di bawah ini karena `s1` belum terdeklarasikan
    
    
    s1 := "Hi, "
    s2 := "world"
    
    // Ubah nilai yang diberikan pada variable saat pendefinisian
    s1 = "Hello, "

    // Karena variable `s1` bertipe string, hapus baris di bawah
    
    
    // Cetak nilai-nilai dari variable s1 dan s2
    println(s1, s2)
    
}

package main

func main() {
    n := 100
    // Kurangkan 10 dari variable `n` lalu tetapkan hasilnya kembali pada `n`
    n -= 10
    
    println(n)
    
    
    m := 10
    // Tambahkan 1 pada variable `m` lalu tetapkan hasilnya kembali pada `m`
    m++
    
    println(m)
}

package main

func main() {
    x := 123
    y := 5 * 6
    
    // Ketika x lebih besar atau sama dengan 100, cetak "x lebih besar atau sama dengan 100"
    if x >= 100 {
        println("x lebih besar atau sama dengan 100")
    }

    // Ketika y kurang dari 40, cetak "y kurang dari 40"
    if y < 40 {
        println("y kurang dari 40")
    }

}

package main

func main() {
    x := 7 * 10
    y := 5 * 6
    
    // Buatlah sebuah pernyataan `if` untuk ketika nilai x sama dengan 70
    if x == 70 {
        println("x sama dengan 70")
    }
    
    // Buatlah sebuah pernyataan `if` untuk ketika nilai y tidak sama dengan 40
    if y != 40 {
        println("y tidak sama dengan 40")
    }
    
}

package main

func main() {
    money := 100
    price := 200
    
    // Tambahkan sebuah pernyataan `else`
    if money >= price {
        println("Anda bisa membeli produk ini")
    } else{
        println("Anda tidak memiliki cukup uang")
    }
    
}

package main

func main() {
    money := 200
    price := 200
    
    // Tambahkan pernyataan `else if` dan  `else` berikut ini
    if money > price {
        println("Anda bisa membeli item ini")
    } else if money == price {
        println("Anda bisa membeli item ini, namun uang tidak akan tersisa")
    } else {
        println("Anda tidak memiliki cukup uang")
    }
}

package main

func main() {
    x := 20
    // Ketika x lebih besar dari atau sama dengan 10 dan kurang dari atau sama dengan 30,
    // cetak pesan "x paling sedikit 10 dan paling banyak 30"
    if x >= 10 && x <= 30 {
        println("x paling sedikit 10 dan paling banyak 30")
    }
    
    y := 60
    // Ketika y kurang dari 10 atau lebih besar dari 30,
    // cetak pesan "y kurang dari 10 atau lebih besar dari 30"
    if y < 10 || y > 30 {
        println("y kurang dari 10 atau lebih besar dari 30")
    }
    
    z := 55
    // Ketika z tidak sama dengan 77,
    // cetak pesan "z tidak sama dengan 77"
    if !(z == 77) {
        println("z tidak sama dengan 77")
    }
 
}

package main

func main() {
    n := 3
    switch n {
        // Tambahkan sebuah `case` untuk mencetak "Belum beruntung" ketika `n` bernilai 0
        case 0:
            println("Belum beruntung")
        // Tambahkan sebuah `case` untuk mencetak "Sedikit beruntung" ketika `n` bernilai 1 atau 2
        case 1, 2:
            println("Sedikit beruntung")
        // Tambahkan sebuah `case` untuk mencetak "Beruntung" ketika `n` bernilai 3 atau 4
        case 3, 4:
            println("Beruntung")
        // Tambahkan sebuah `case` untuk mencetak "Sangat beruntung" ketika `n` bernilai 5
        case 5:
            println("Sangat beruntung")
    }
}
