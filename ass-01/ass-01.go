package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Biodata struct {
	nama        string
	kodePeserta int
	alamat      string
	pekerjaan   string
	alasan      []string
}

func main() {
	var x Biodata

	x.nama = "Dita Lastri"
	x.kodePeserta = 477
	x.alamat = "Bekasi"
	x.pekerjaan = "Belum bekerja"
	x.alasan = []string{"1. Pernah liat tutorial dasar di youtube", "2. Penasaran bagaimana pengimplementasian nya"}

	var y = os.Args //os.Args bertipe []string
	find(y, x)
}

func find(arg []string, lala Biodata) {

	arg = arg[1:]
	kod := strconv.FormatInt(int64(lala.kodePeserta), 10)
	//strconv.FormatInt(i int64, base int)
	som := strings.ToLower(lala.nama)

	for _, v := range arg {
		if strings.ToLower(v) == som || v == kod {
			fmt.Printf("%+v\n", lala)
			break
		}
		if strings.ToLower(v) != som || v != kod {
			fmt.Println("cobalagi")
		}
	}
}
