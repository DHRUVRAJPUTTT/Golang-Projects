package data

type OTPData struct {
	PhoneNumber string `json:"phoneNumber"`
}

type VerifyData struct {
	User *OTPData `json:"user,omitempty"`
	Code string   `json:"code,omitempty"`
}
