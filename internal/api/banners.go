package api

import (
	"bannerlord/internal/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserBanner(c *gin.Context) {
	const op = "api/handlers/banners.go/GetUserBanner"

	featureS, featureOK := c.GetQuery("feature_id")
	tagS, tagOK := c.GetQuery("tag_id")
	if !(featureOK && tagOK) {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrResp(NoEnoughInfo))
		return
	}

	feature, featureErr := strconv.Atoi(featureS)
	tag, tagErr := strconv.Atoi(tagS)
	if featureErr != nil || tagErr != nil {
		log.Printf(op + ": error with converting ids")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	banners, err := Cfg.Storage.GetBanners(feature, tag, c.GetBool("super_user"))
	if err != nil {
		log.Printf(op + ": error with quering banners")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, banners)
}

func GetBanners(c *gin.Context) {
	const op = "api/handlers/banners.go/GetBanners"

	featureS, featureOK := c.GetQuery("feature_id")
	tagS, tagOK := c.GetQuery("tag_id")
	var result []models.Banner
	var err error

	switch {
	case featureOK && tagOK:
		feature, featureErr := strconv.Atoi(featureS)
		tag, tagErr := strconv.Atoi(tagS)
		if featureErr != nil || tagErr != nil {
			log.Printf(op + "error with conv tag and feature ")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		result, err = Cfg.Storage.GetBanners(feature, tag, false)
		if err != nil {
			log.Printf(op+"error with querying banners: %v", err)
		}

	case featureOK:
		feature, featureErr := strconv.Atoi(featureS)
		if featureErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		result, err = Cfg.Storage.GetByFeature(feature)
		if err != nil {
			log.Printf(op+"error with querying banners: %v", err)
		}

	case tagOK:
		tag, tagErr := strconv.Atoi(featureS)
		if tagErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		result, err = Cfg.Storage.GetByTag(tag)
		if err != nil {
			log.Printf(op+"error with querying banners: %v", err)
		}

	default:
		result, err = Cfg.Storage.GetAll()
		if err != nil {
			log.Printf(op+"error with querying banners: %v", err)
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

func CreateBanner(c *gin.Context) {
	banner := models.BannerDTO{}
	err := c.BindJSON(&banner)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(banner)
}

func UpdateBanner(c *gin.Context) {
	const op = "api/banners.go/UpdateBanner"

	banner := models.BannerPatch{}
	err := c.BindJSON(&banner)
	if err != nil {
		log.Printf(op+"err with binding json: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idS := c.Param("id")
	id, _ := strconv.Atoi(idS)
	banner.ID = id
	if err = Cfg.Storage.UpdateBanner(&banner); err != nil {
		log.Printf(op+"error with updating %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
	return
}

func DeleteBanner(c *gin.Context) {
	const op = "DeleteBanner"
	idS, ok := c.GetQuery("id")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err = Cfg.Storage.DeleteBanner(id); err != nil {
		log.Printf(op+"delete error, %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
	return
}
