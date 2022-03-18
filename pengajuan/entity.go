package pengajuan

type Pengajuan struct {
	NoPengajuan    int
	NoAnggota      int
	Petugas        string
	TglPinjam      string
	TujuanPinjam   string
	Jaminan        string
	LamaPinjam     int
	JumlahPinjaman int
	Catatan        string
	SukuBunga      int
	Nama           string
	StatusDraf     int
	Persetujuan    int
}

type BuktiJaminan struct {
	Id          int
	IdPengajuan int
	Bukti       string
}

type JaminanBarang struct {
	Id                     int
	IdPengajuan            int
	JenisBarangJaminan     string
	NamaBarangJaminan      string
	KondisiBarangJaminan   string
	DeskripsiBarangJaminan string
}

type JaminanTanah struct {
	Id                     int
	IdPengajuan            int
	Latitude               string
	Longitude              string
	Altitude               string
	ArahLokasi             string
	NamaBarangJaminan      string
	KondisiBarangJaminan   string
	DeskripsiBarangJaminan string
}

type Notifikasi struct {
	Id          int
	IdPengajuan int
	Petugas     string
	TglSurvey   string
	Pukul       string
}
