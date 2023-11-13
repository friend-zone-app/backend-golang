package authentication

import "time"

type (
	OtpStruct struct {
		userId    string
		code      string
		createdAt time.Time
	}

	OtpMap map[string]OtpStruct

	Otp struct {
		otp OtpMap
	}
)

func NewOtp() Otp {
	return Otp{make(OtpMap)}
}

func (handler Otp) GetOTP() OtpMap {
	return handler.otp
}

func (handler Otp) Get(userId string) (*OtpStruct, bool) {
	otp, found := handler.otp[userId]

	return &otp, found
}

func (handler Otp) Write(userId, code string) (*OtpStruct, bool) {
	now := time.Now()
	otp, found := handler.otp[userId]
	if !found {
		return &otp, found
	} else {
		otp.createdAt = now
		otp.code = code
		return &otp, found
	}
}

func (handler Otp) Delete(userId string) bool {
	delete(handler.GetOTP(), userId)
	return true
}

func (handler Otp) Register(userId, code string) {
	now := time.Now()
	otpstruct := OtpStruct{userId: userId, createdAt: now, code: code}
	handler.otp[userId] = otpstruct
	if len(userId) > 1 {
		handler.otp[userId[:1]] = otpstruct
	}
}

func (otp OtpStruct) GetTime() time.Time {
	return otp.createdAt
}
