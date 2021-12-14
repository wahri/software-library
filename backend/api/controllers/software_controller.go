package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"software_library/backend/api/models"
	"software_library/backend/api/responses"
	"software_library/backend/api/utils/formaterror"
	"strconv"
	"time"
)

func (server *Server) CreateSoftware(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()
	timeUpload := now.Unix()
	nameFile := r.FormValue("name")
	filename := fmt.Sprintf("%d-%s%s", timeUpload, nameFile, filepath.Ext(handler.Filename))

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// }

	software := models.Software{}

	// err = json.Unmarshal(body, &software)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	software.Name = r.FormValue("name")
	software.ZipFile = fileLocation
	software.LinkSource = r.FormValue("LinkSource")
	software.LinkPreview = r.FormValue("LinkPreview")
	software.LinkTutorial = r.FormValue("LinkTutorial")
	software.License = r.FormValue("License")
	software.Description = r.FormValue("Description")
	software.PreviewImage = r.FormValue("PreviewImage")

	productVersion, _ := strconv.ParseFloat(r.FormValue("ProductVersion"), 64)
	software.ProductVersion = productVersion

	softwareCreated, err := software.SaveSoftware(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, softwareCreated.ID))
	responses.JSON(w, http.StatusCreated, softwareCreated)
}
