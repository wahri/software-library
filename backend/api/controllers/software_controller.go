package controllers

import (
	"fmt"
	"net/http"
	"software_library/backend/api/models"
	"software_library/backend/api/responses"
	"software_library/backend/api/utils/formaterror"
	upload "software_library/backend/api/utils/uploadfile"
	"strconv"
)

func (server *Server) CreateSoftware(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

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
	software.ZipFile, _ = upload.UploadFile(w, r, "ZipFile", "ZipFile")
	software.LinkSource = r.FormValue("LinkSource")
	software.LinkPreview = r.FormValue("LinkPreview")
	software.LinkTutorial = r.FormValue("LinkTutorial")
	software.License = r.FormValue("License")
	software.Description = r.FormValue("Description")
	software.PreviewImage, _ = upload.UploadFile(w, r, "PreviewImage", "PreviewImage")

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
