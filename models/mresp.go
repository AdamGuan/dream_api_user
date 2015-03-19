package models

type MResp struct {
	responseNo  int
	responseMsg string
}

type MFindPwdResp struct {
	responseNo  int
	responseMsg string
	password string
}

type MUserExistsResp struct {
	responseNo  int
	responseMsg string
	exists string
}

type MUserLoginResp struct {
	responseNo  int
	responseMsg string
	F_uid string
	F_phone_number string
	F_crate_datetime string
	F_modify_datetime string
}

type MUserInfoResp struct {
	responseNo  int
	responseMsg string
	F_uid string
	F_phone_number string
	F_crate_datetime string
	F_modify_datetime string
}

type MModifyPhoneResp struct {
	responseNo  int
	responseMsg string
}