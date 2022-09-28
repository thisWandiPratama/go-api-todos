package user

type UserFormatter struct {
	ID      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
	Avatar  string `json:"avatar"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:      user.ID,
		Nama:    user.Nama,
		Alamat:  user.Alamat,
		Telepon: user.Telepon,
		Avatar:  user.Avatar,
		Email:   user.Email,
		Token:   token,
	}

	return formatter
}
