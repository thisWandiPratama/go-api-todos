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
