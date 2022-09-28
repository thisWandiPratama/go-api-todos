package user

type RegisterUserInput struct {
	Nama     string `json:"nama" binding:"required"`
	Alamat   string `json:"alamat" binding:"required"`
	Telepon  string `json:"telepon" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type FormCreateUserInput struct {
	Nama            string `json:"nama" binding:"required"`
	Alamat          string `json:"alamat" binding:"required"`
	Telepon         string `json:"telepon" binding:"required"`
	JenisKelamin    string `json:"jenis_kelamin" binding:"required"`
	Role            string `json:"role" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	GoogleSignId    string `json:"google_sign_id" binding:"required"`
	SyaratKetentuan string `json:"syarat_ketentuan" binding:"required"`
	Error           error
}

type FormUpdateUserInput struct {
	ID           int
	Nama         string `json:"nama" binding:"required"`
	Alamat       string `json:"alamat" binding:"required"`
	Telepon      string `json:"telepon" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
	Error        error
}
