package pengajuan

type UserFormatter struct {
	NoPengajuan            int                     `json:"no_pengajuan"`
	NoAnggota              int                     `json:"no_anggota"`
	Petugas                string                  `json:"petugas"`
	TglPinjam              string                  `json:"tgl_pinjam"`
	TujuanPinjam           string                  `json:"tujuan_pinjam"`
	Jaminan                string                  `json:"jaminan"`
	LamaPinjam             int                     `json:"lama_pinjam"`
	JumlahPinjam           int                     `json:"jumlah_pinjaman"`
	Catatan                string                  `json:"catatan"`
	SukuBunga              int                     `json:"suku_bunga"`
	Nama                   string                  `json:"nama"`
	JenisBarangJaminan     string                  `json:"jenis_barang_jaminan"`
	NamaBarangJaminan      string                  `json:"nama_barang_jaminan"`
	KondisiBarangJaminan   string                  `json:"kondisi_barang_jaminan"`
	DeskripsiBarangJaminan string                  `json:"deskripsi_barang_jaminan"`
	BuktiJaminan           []FormatterBuktiJaminan `json:"bukti_jaminan"`
}

func FormatUser(user Pengajuan, bukti []FormatterBuktiJaminan) UserFormatter {
	formatter := UserFormatter{
		NoPengajuan:            user.NoPengajuan,
		NoAnggota:              user.NoAnggota,
		Petugas:                user.Petugas,
		TglPinjam:              user.TglPinjam,
		TujuanPinjam:           user.TujuanPinjam,
		Jaminan:                user.Jaminan,
		LamaPinjam:             user.LamaPinjam,
		JumlahPinjam:           user.JumlahPinjaman,
		Catatan:                user.Catatan,
		SukuBunga:              user.SukuBunga,
		Nama:                   user.Nama,
		JenisBarangJaminan:     user.JenisBarangJaminan,
		NamaBarangJaminan:      user.NamaBarangJaminan,
		KondisiBarangJaminan:   user.KondisiBarangJaminan,
		DeskripsiBarangJaminan: user.DeskripsiBarangJaminan,
		BuktiJaminan:           bukti,
	}

	return formatter
}

func FormatAlls(teachers []Pengajuan, bukti []FormatterBuktiJaminan) []UserFormatter {
	teachersFormatter := []UserFormatter{}

	for _, teac := range teachers {
		teacherFormatter := FormatUser(teac, bukti)
		teachersFormatter = append(teachersFormatter, teacherFormatter)

	}

	return teachersFormatter
}

type FormatterBuktiJaminan struct {
	Id          int    `json:"id"`
	IdPengajuan int    `json:"id_pengajuan"`
	Bukti       string `json:"avatar_bukti"`
}

func FormatBukti(user BuktiJaminan) FormatterBuktiJaminan {
	formatter := FormatterBuktiJaminan{
		Id:          user.Id,
		IdPengajuan: user.IdPengajuan,
		Bukti:       user.Bukti,
	}

	return formatter
}

func FormatBuktiAlls(teachers []BuktiJaminan) []FormatterBuktiJaminan {
	teachersFormatter := []FormatterBuktiJaminan{}

	for _, teac := range teachers {
		teacherFormatter := FormatBukti(teac)
		teachersFormatter = append(teachersFormatter, teacherFormatter)

	}

	return teachersFormatter
}
