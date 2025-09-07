package dto

type SendOTPRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type VerifyOTPRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
	Code         string `json:"code" binding:"required,numeric,len=6"`
}

type OTPResponse struct {
	Message string `json:"message"`
	Expires int64  `json:"expires_in"`
}