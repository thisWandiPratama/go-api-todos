package mahasiswa

type UserFormatter struct {
	Id     int     `json:"id"`
	Npm    string  `json:"npm"`
	Nama   string  `json:"nama"`
	NmMk   string  `json:"nm_mk"`
	Ntugas float64 `json:"ntugas"`
	Nquiz  float64 `json:"nquiz"`
	Nuts   float64 `json:"nuts"`
	Nuas   float64 `json:"nuas"`
	Total  float64 `json:"total"`
}

func FormatMahasiswa(mahasiswa Mahasiswa) UserFormatter {
	formatter := UserFormatter{
		Id:     mahasiswa.Id,
		Npm:    mahasiswa.Npm,
		Nama:   mahasiswa.Nama,
		NmMk:   mahasiswa.NmMk,
		Ntugas: mahasiswa.Ntugas,
		Nquiz:  mahasiswa.Nquiz,
		Nuts:   mahasiswa.Nuts,
		Nuas:   mahasiswa.Nuts,
		Total:  mahasiswa.Total,
	}

	return formatter
}
