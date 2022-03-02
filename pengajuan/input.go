package pengajuan

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateJaminanPengajuan struct {
	IdPengajuan            int    `json:"id_pengajuan" binding:"required"`
	JenisBarangJaminan     string `json:"jenis_barang_jaminan" binding:"required"`
	NamaBarangJaminan      string `json:"nama_barang_jaminan" binding:"required"`
	KondisiBarangJaminan   string `json:"kondisi_barang_jaminan" binding:"required"`
	DeskripsiBarangJaminan string `json:"deskripsi_barang_jaminan" binding:"required"`
}

type AddBuktiJaminan struct {
	IdPengajuan int    `json:"id_pengajuan" binding:"required"`
	Bukti       string `json:"avatar_bukti" binding:"required"`
	StatusDraf  int    `json:"status_draf" binding:"required"`
}

type AddJaminanTanah struct {
	IdPengajuan            int    `json:"id_pengajuan" binding:"required"`
	Latitude               string `json:"latitude" binding:"required"`
	Longitude              string `json:"longitude" binding:"required"`
	Altitude               string `json:"altitude" binding:"required"`
	ArahLokasi             string `json:"arah_lokasi" binding:"required"`
	NamaBarangJaminan      string `json:"nama_barang_jaminan" binding:"required"`
	KondisiBarangJaminan   string `json:"kondisi_barang_jaminan" binding:"required"`
	DeskripsiBarangJaminan string `json:"deskripsi_barang_jaminan" binding:"required"`
}

type SearchJaminan struct {
	Jaminan   string `json:"jaminan"`
	TglPinjam string `json:"tgl_pinjam"`
}

type AddNotifikasi struct {
	IdPengajuan int    `json:"id_pengajuan" binding:"required"`
	Petugas     string `json:"petugas" binding:"required"`
}
