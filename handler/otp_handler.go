package handler

import (
	"net/http"
	"user-crud/dto"
	"user-crud/usecase"

	"github.com/gin-gonic/gin"
)

type OTPHandler struct {
	otpUsecase *usecase.OTPUsecase
}

func NewOTPHandler(otpUsecase *usecase.OTPUsecase) *OTPHandler {
	return &OTPHandler{otpUsecase: otpUsecase}
}

func (h *OTPHandler) SendOTP(c *gin.Context) {
	var req dto.SendOTPRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.otpUsecase.SendOTP(req.MobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.OTPResponse{
		Message: "OTP sent successfully",
		Expires: 120, // 2 minutes
	})
}

func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := h.otpUsecase.VerifyOTP(req.MobileNumber, req.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}