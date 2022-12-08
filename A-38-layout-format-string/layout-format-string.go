package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	// numerik menjadi string numerik berbasis 2(biner)
	fmt.Printf("%b\n", data.age)
	fmt.Println()

	// numerik unicode menjadi string karakter unicodenya
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)
	fmt.Println()

	// numerik menjadi string numerik berbasis 10(basis yang kita gunakan)
	fmt.Printf("%d\n", data.age)
	fmt.Println()

	// numerik desimal menjadi notasi numerik standar
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)
	fmt.Println()

	// numerik desimal dengan lebar desimal bisa disesuaikan(default 6 digit)
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)
	fmt.Println()

	// numerik desimal dengan lebar desimal bisa disesuaikan
	fmt.Printf("%e\n", 0.123123123123)
	fmt.Printf("%f\n", 0.123123123123)
	fmt.Printf("%g\n", 0.123123123123)
	// perbedaan lain yaitu digit desimal sesuai dengan data. tidak bisa dicustom
	fmt.Printf("%g\n", 0.12)
	fmt.Printf("%.5g\n", 0.12)
	fmt.Println()

	// numerik menjadi string numerik berbasis 8(oktal)
	fmt.Printf("%o\n", data.age)
	fmt.Println()

	// mengembalikan alamat pointer referensi variabelnya
	fmt.Printf("%p\n", &data.name)
	fmt.Println()

	// escape string
	fmt.Printf("%q\n", `" name \ height "`)
	fmt.Println()

	// data string
	fmt.Printf("%s\n", data.name)
	fmt.Println()

	// data boolean
	fmt.Printf("%t\n", data.isGraduated)
	fmt.Println()

	// mengambil tipe variabel yang digunakan
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.age)
	fmt.Printf("%T\n", data.isGraduated)
	fmt.Printf("%T\n", data.hobbies)
	fmt.Println()

	// data apa saja(termasuk interface{})
	fmt.Printf("%v\n", data)
	fmt.Println()

	// data struct, mengembalikan nama properti dan nilainya
	fmt.Printf("%+v\n", data)
	fmt.Println()

	// data struct, mengembalikan nama dan nilai serta bagaimana objek dideklarasikan
	fmt.Printf("%#v\n", data)
	fmt.Println()

	// numerik menjadi string numerik berbasis 16
	fmt.Printf("%x\n", data.age)
	// jika digunakan pada string mengembalikan kode heksadesimal tiap karakter
	var d = data.name
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	fmt.Printf("%x\n", d)
	fmt.Println()

	// menulis karakter %
	fmt.Printf("%%\n")
}
