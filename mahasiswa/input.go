package mahasiswa

type AddMahasiswa struct {
	Npm    string  `json:"npm" binding:"required"`
	Nama   string  `json:"nama" binding:"required"`
	NmMk   string  `json:"nm_mk" binding:"required"`
	Ntugas float64 `json:"ntugas" binding:"required"`
	Nquiz  float64 `json:"nquiz" binding:"required"`
	Nuts   float64 `json:"nuts" binding:"required"`
	Nuas   float64 `json:"nuas" binding:"required"`
}
