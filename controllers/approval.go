package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shunwuse/go-hris/constants"
	"github.com/shunwuse/go-hris/dtos"
	"github.com/shunwuse/go-hris/lib"
	"github.com/shunwuse/go-hris/models"
	"github.com/shunwuse/go-hris/services"
)

type ApprovalController struct {
	logger          lib.Logger
	approvalService services.ApprovalService
}

func NewApprovalController() ApprovalController {
	logger := lib.GetLogger()

	// Initialize services
	approvalService := services.NewApprovalService()

	return ApprovalController{
		logger:          logger,
		approvalService: approvalService,
	}
}

func (c ApprovalController) GetApprovals(ctx *gin.Context) {
	approvals, err := c.approvalService.GetApprovals()
	if err != nil {
		c.logger.Errorf("Error getting approvals: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting approvals",
		})
		return
	}

	approvalsResponse := make([]dtos.ApprovalResponse, 0)
	for _, approval := range approvals {
		approvalResponse := dtos.ApprovalResponse{
			ID:          approval.ID,
			CreatorName: approval.Creator.Name,
			Status:      string(approval.Status),
		}

		if approval.Approver != nil {
			approvalResponse.ApproverName = &approval.Approver.Name
		}

		approvalsResponse = append(approvalsResponse, approvalResponse)

	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": approvalsResponse,
	})
}

func (c ApprovalController) AddApproval(ctx *gin.Context) {
	token := ctx.MustGet(constants.JWTClaims).(services.TokenPayload)
	userID := token.UserID

	approval := models.Approval{
		CreatorID: userID,
		Status:    constants.ApprovalStatusPending,
	}

	err := c.approvalService.AddApproval(approval)
	if err != nil {
		c.logger.Errorf("Error adding approval: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error adding approval",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Approval added successfully",
	})
}