package controllers

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/dejavuzhou/felix/ssh2ws/utils"
	"github.com/pkg/sftp"
	"log"
	"mime/multipart"
	"net/http"
	"path"
)

func SftpUp(w http.ResponseWriter, r *http.Request) {
	response := models.JsonResponse{HasError: true}
	idx := utils.GetQueryInt(r, "id", 0)
	mc, err := models.MachineFind(uint(idx))
	if err != nil {
		response.Message = err
		utils.ServeJSON(w, response)
		return
	}
	sftpClient := flx.NewSftpClient(mc)

	//file, header, err := this.GetFile("file")
	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("file")
	relativePath := r.URL.Query().Get("path") // get path. default is ""
	if err != nil {
		log.Println("Error: getfile err ", err)
		utils.Abort(w, "error", 503)
		return
	}
	defer file.Close()

	if err := UploadFile(relativePath, sftpClient, file, header); err != nil {
		log.Println("Error: sftp error:", err)
		utils.Abort(w, "message", 503)
	} else {
		w.Write([]byte("success")) // todo write file name back.
	}
}

// upload file to server via sftp.
/**
@desPath: relative path in remote server.
*/
func UploadFile(desPath string, client *sftp.Client, srcFile multipart.File, header *multipart.FileHeader) error {
	var fullPath string
	if wd, err := client.Getwd(); err == nil {
		fullPath = path.Join(wd, desPath)
		if _, err := client.Stat(fullPath); err != nil {
			return err // check path must exist
		}
	} else {
		return err
	}

	dstFile, err := client.Create(path.Join(fullPath, header.Filename))
	if err != nil {
		return err
	}
	defer srcFile.Close()
	defer dstFile.Close()

	_, err = dstFile.ReadFrom(srcFile)
	if err != nil {
		return err
	}
	return nil
}
