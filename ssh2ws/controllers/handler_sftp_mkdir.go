package controllers

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"log"
	"strconv"
	"strings"
)

func SftpMkdir(c *gin.Context) {
	sftpClient, err := getSftpClient(c)
	if handleError(c, err) {
		return
	}
	fullPath := c.Query("path")
	if strings.HasPrefix(fullPath, "$HOME") {
		wd, err := sftpClient.Getwd()
		if handleError(c, err) {
			return
		}
		fullPath = strings.Replace(fullPath, "$HOME", wd, 1)
	}
	log.Println(fullPath)

	err = sftpClient.Mkdir(fullPath)
	if handleError(c, err) {
		return
	}
	jsonSuccess(c, "ok"+fullPath)
}

func getSftpClient(c *gin.Context) (*sftp.Client, error) {
	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	mc, err := models.MachineFind(uint(idx))
	if err != nil {
		return nil, err
	}
	return flx.NewSftpClient(mc)
}
