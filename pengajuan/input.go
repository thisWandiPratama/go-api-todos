package pengajuan

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateJaminanPengajuan struct {
	NoPengajuan            int    `json:"no_pengajuan" binding:"required"`
	JenisBarangJaminan     string `json:"jenis_barang_jaminan" binding:"required"`
	NamaBarangJaminan      string `json:"nama_barang_jaminan" binding:"required"`
	KondisiBarangJaminan   string `json:"kondisi_barang_jaminan" binding:"required"`
	DeskripsiBarangJaminan string `json:"deskripsi_barang_jaminan" binding:"required"`
}

type AddBuktiJaminan struct {
	IdPengajuan int    `json:"id_pengajuan" binding:"required"`
	Bukti       string `json:"avatar_bukti" binding:"required"`
}
