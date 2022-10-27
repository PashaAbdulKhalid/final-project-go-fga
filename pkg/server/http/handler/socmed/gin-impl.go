package socmed

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/message"
	"github.com/PashaAbdulKhalid/final-project-go-fga/pkg/domain/socmed"
	"github.com/gin-gonic/gin"
)

type SocialMediaHdlImpl struct {
	socialMediaUsecase socmed.SocialMediaUsecase
}

func NewSocialMediaHandler(socialMediaUsecase socmed.SocialMediaUsecase) socmed.SocialMediaHandler {
	return &SocialMediaHdlImpl{socialMediaUsecase: socialMediaUsecase}
}

func (c *SocialMediaHdlImpl) CreateSocialMediaHdl(ctx *gin.Context) {
	log.Printf("%T - CreateSocialMediaHdl is invoked\n", c)
	defer log.Printf("%T - CreateSocialMediaHdl executed\n", c)

	var inputSocialMedia socmed.SocialMedia

	log.Println("binding body payload from request")
	if err := ctx.ShouldBindJSON(&inputSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed to bind payload",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PAYLOAD",
				"error_message": err.Error(),
			},
		})
		return
	}

	log.Println("calling create socialMedia usecase service")
	result, errMsg := c.socialMediaUsecase.CreateSocialMediaSvc(ctx, inputSocialMedia)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "social media has successfully created",
		"type":    "ACCEPTED",
		"data": gin.H{
			"id":         result.ID,
			"name":       result.Name,
			"url":        result.URL,
			"created_at": result.CreatedAt,
			"user_id":    result.UserID,
		},
	})
}

func (c *SocialMediaHdlImpl) GetSocialMediasHdl(ctx *gin.Context) {
	log.Printf("%T - GetSocialMediasIdHdl is invoked\n", c)
	defer log.Printf("%T - GetSocialMediasHdl executed\n", c)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	log.Println("calling get social medias usecase service")
	result, errMsg := c.socialMediaUsecase.GetSocialMediasSvc(ctx)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": fmt.Sprintf("social medias by user id %v is found", userId),
		"type":    "SUCCESS",
		"data":    result,
	})
}

func (c *SocialMediaHdlImpl) UpdateSocialMediaHdl(ctx *gin.Context) {
	log.Printf("%T - UpdateSocialMediaHdl is invoked\n", c)
	defer log.Printf("%T - UpdateSocialMediaHdl executed\n", c)

	log.Println("check social media id from path parameter")
	socmedIdParam := ctx.Param("socialMediaId")

	socmedId, err := strconv.ParseUint(socmedIdParam, 0, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "invalid params",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PARAMS",
				"error_message": "invalid params",
			},
		})
		return
	}

	log.Println("calling get social media by id usecase service")
	result, errMsg := c.socialMediaUsecase.GetSocialMediaByIdSvc(ctx, socmedId)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	stringUserId := ctx.Value("user").(string)
	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	log.Println("verify the social media belongs to")
	if result.UserID != userId {
		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
			Type:  "INVALID_SCOPE",
			Error: errors.New("cannot update the social media"),
		})
		return
	}

	var updatedSocialMedia socmed.SocialMedia

	log.Println("binding body payload from request")
	if err := ctx.ShouldBindJSON(&updatedSocialMedia); err != nil {
		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
			Type:  "INVALID_PAYLOAD",
			Error: errors.New("failed to bind payload"),
		})
		return
	}

	ctx.Set("socmedId", socmedId)

	updateResult, errMsg := c.socialMediaUsecase.UpdateSocialMediaSvc(ctx, updatedSocialMedia)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    01,
		"message": "social media has been successfully updated",
		"type":    "ACCEPTED",
		"data":    updateResult,
	})
}

func (c *SocialMediaHdlImpl) DeleteSocialMediaHdl(ctx *gin.Context) {
	log.Printf("%T - DeleteSocialMediaHdl is invoked\n", c)
	defer log.Printf("%T - DeleteSocialMediaHdl executed\n", c)

	log.Println("check socmedId from path parameter")
	socmedIdParam := ctx.Param("socialMediaId")

	socmedId, err := strconv.ParseUint(socmedIdParam, 0, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "invalid params",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PARAMS",
				"error_message": "invalid params",
			},
		})
		return
	}

	log.Println("calling get social media by id usecase service")
	result, errMsg := c.socialMediaUsecase.GetSocialMediaByIdSvc(ctx, socmedId)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	stringUserId := ctx.Value("user").(string)
	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	log.Println("verify the social media belongs to")
	if result.UserID != userId {
		message.ErrorResponseSwitcher(ctx, message.ErrorMessage{
			Type:  "INVALID_SCOPE",
			Error: errors.New("cannot delete the social media"),
		})
		return
	}

	log.Println("calling delete social media usecase service")
	errMsg = c.socialMediaUsecase.DeleteSocialMediaSvc(ctx, socmedId)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    01,
		"message": "social media has been successfully deleted",
		"type":    "ACCEPTED",
	})
}

