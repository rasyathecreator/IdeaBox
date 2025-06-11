package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ide struct {
	ID          int
	Judul       string
	Deskripsi   string
	Kategori    string
	Status      string
	VotePositif int
	VoteNegatif int
	Komentar    []string
}

var daftarIde []Ide
var idBerikutnya int = 1

func input(prompt string) string {
	fmt.Print(prompt)
	pembaca := bufio.NewReader(os.Stdin)
	teks, _ := pembaca.ReadString('\n')
	return strings.TrimSpace(teks)
}

func tambahIde() {
	judul := input("Judul ide: ")
	deskripsi := input("Deskripsi singkat: ")
	kategori := input("Kategori (Produk/Marketing/Fitur): ")
	ideBaru := Ide{
		ID:        idBerikutnya,
		Judul:     judul,
		Deskripsi: deskripsi,
		Kategori:  kategori,
		Status:    "Baru",
	}
	idBerikutnya++
	daftarIde = append(daftarIde, ideBaru)
	fmt.Println("✅ Ide berhasil ditambahkan!\n")
}

func lihatIde() {
	if len(daftarIde) == 0 {
		fmt.Println("Belum ada ide yang tercatat.\n")
		return
	}
	for _, ide := range daftarIde {
		fmt.Printf("[ID: %d] %s\nKategori: %s | Status: %s\nDeskripsi: %s\n👍 %d 👎 %d | Komentar: %d\n\n",
			ide.ID, ide.Judul, ide.Kategori, ide.Status, ide.Deskripsi, ide.VotePositif, ide.VoteNegatif, len(ide.Komentar))
	}
}

func ubahStatus() {
	idStr := input("Masukkan ID ide yang ingin diubah statusnya: ")
	id, _ := strconv.Atoi(idStr)
	for i := range daftarIde {
		if daftarIde[i].ID == id {
			status := input("Status baru (Baru/Ditinjau/Dikembangkan/Ditolak): ")
			daftarIde[i].Status = status
			fmt.Println("✅ Status berhasil diubah!\n")
			return
		}
	}
	fmt.Println("❌ ID tidak ditemukan.\n")
}

func tambahKomentar() {
	idStr := input("Masukkan ID ide yang ingin dikomentari: ")
	id, _ := strconv.Atoi(idStr)
	for i := range daftarIde {
		if daftarIde[i].ID == id {
			komentar := input("Masukkan komentar Anda: ")
			daftarIde[i].Komentar = append(daftarIde[i].Komentar, komentar)
			fmt.Println("✅ Komentar ditambahkan!\n")
			return
		}
	}
	fmt.Println("❌ ID tidak ditemukan.\n")
}

func voteIde() {
	idStr := input("Masukkan ID ide yang ingin diberi vote: ")
	id, _ := strconv.Atoi(idStr)
	for i := range daftarIde {
		if daftarIde[i].ID == id {
			vote := input("Vote (1 untuk 👍, 0 untuk 👎): ")
			if vote == "1" {
				daftarIde[i].VotePositif++
			} else {
				daftarIde[i].VoteNegatif++
			}
			fmt.Println("✅ Vote dicatat!\n")
			return
		}
	}
	fmt.Println("❌ ID tidak ditemukan.\n")
}

func hapusIde() {
	idStr := input("Masukkan ID ide yang ingin dihapus: ")
	id, _ := strconv.Atoi(idStr)
	for i, ide := range daftarIde {
		if ide.ID == id {
			daftarIde = append(daftarIde[:i], daftarIde[i+1:]...)
			fmt.Println("✅ Ide berhasil dihapus!\n")
			return
		}
	}
	fmt.Println("❌ ID tidak ditemukan.\n")
}

func filterIde() {
	kategori := input("Masukkan kategori (Produk/Marketing/Fitur): ")
	adaIde := false
	for _, ide := range daftarIde {
		if strings.EqualFold(ide.Kategori, kategori) {
			fmt.Printf("[ID: %d] %s\nKategori: %s | Status: %s\nDeskripsi: %s\n👍 %d 👎 %d | Komentar: %d\n\n",
				ide.ID, ide.Judul, ide.Kategori, ide.Status, ide.Deskripsi, ide.VotePositif, ide.VoteNegatif, len(ide.Komentar))
			adaIde = true
		}
	}
	if !adaIde {
		fmt.Println("❌ Tidak ada ide dalam kategori tersebut.\n")
	}
}

func lihatKomentar() {
	idStr := input("Masukkan ID ide untuk melihat komentarnya: ")
	id, _ := strconv.Atoi(idStr)
	for _, ide := range daftarIde {
		if ide.ID == id {
			fmt.Printf("Komentar untuk ide \"%s\":\n", ide.Judul)
			if len(ide.Komentar) == 0 {
				fmt.Println("Belum ada komentar untuk ide ini.\n")
			} else {
				for i, komentar := range ide.Komentar {
					fmt.Printf("%d. %s\n", i+1, komentar)
				}
				fmt.Println()
			}
			return
		}
	}
	fmt.Println("❌ ID tidak ditemukan.\n")
}

func cariIde(kataKunci string) {
	adaIde := false
	for _, ide := range daftarIde {
		if strings.Contains(strings.ToLower(ide.Judul), strings.ToLower(kataKunci)) ||
			strings.Contains(strings.ToLower(ide.Deskripsi), strings.ToLower(kataKunci)) {
			fmt.Printf("[ID: %d] %s\nKategori: %s | Status: %s\nDeskripsi: %s\n👍 %d 👎 %d | Komentar: %d\n\n",
				ide.ID, ide.Judul, ide.Kategori, ide.Status, ide.Deskripsi, ide.VotePositif, ide.VoteNegatif, len(ide.Komentar))
			adaIde = true
		}
	}
	if !adaIde {
		fmt.Println("❌ Tidak ada ide yang cocok dengan kata kunci tersebut.\n")
	}
}

func main() {
	for {
		fmt.Println("=== IdeaBox - Pengelolaan Ide Startup ===")
		fmt.Println("1. Lihat Daftar Ide")
		fmt.Println("2. Tambah Ide Baru")
		fmt.Println("3. Ubah Status Ide")
		fmt.Println("4. Komentar / Diskusi")
		fmt.Println("5. Voting Ide")
		fmt.Println("6. Hapus Ide")
		fmt.Println("7. Filter Berdasarkan Kategori")
		fmt.Println("8. Lihat Komentar")
		fmt.Println("9. Cari Ide")
		fmt.Println("10. Keluar")

		pilih := input("Pilih menu: ")

		switch pilih {
		case "1":
			lihatIde()
		case "2":
			tambahIde()
		case "3":
			ubahStatus()
		case "4":
			tambahKomentar()
		case "5":
			voteIde()
		case "6":
			hapusIde()
		case "7":
			filterIde()
		case "8":
			lihatKomentar()
		case "9":
			kataKunci := input("Masukkan kata kunci pencarian: ")
			cariIde(kataKunci)
		case "10":
			fmt.Println("👋 Terima kasih, sampai jumpa!")
			return
		default:
			fmt.Println("❌ Menu tidak tersedia.\n")
		}
	}
}
