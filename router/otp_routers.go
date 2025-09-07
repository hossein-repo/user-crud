package router

import (
	"user-crud/config"
	"user-crud/handler"
	"user-crud/repository"
	"user-crud/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterOTPRoutes(r *gin.Engine, cfg *config.Config) {
	otpRepo := repository.NewOTPRepository()
	otpUsecase := usecase.NewOTPUsecase(cfg, otpRepo)
	otpHandler := handler.NewOTPHandler(otpUsecase)

	otpGroup := r.Group("/otp")
	{
		otpGroup.POST("/send", otpHandler.SendOTP)
		otpGroup.POST("/verify", otpHandler.VerifyOTP)
	}

	// شروع cleanup routine
	go otpUsecase.CleanupExpiredOTPs()
}
