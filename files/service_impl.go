package files

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-pg/pg/v10"

	"lending-service/account"
	"lending-service/utility/wraped_error"
)

type serviceImpl struct {
	repository Repository
}

func newService(repository Repository) serviceImpl {
	return serviceImpl{
		repository: repository,
	}
}

func (svc serviceImpl) proceedSaveFile(file multipart.File, fileHeader *multipart.FileHeader, account account.Account) (uploadResponse UploadResponseDto, errWrap *wraped_error.Error) {
	byteFile, err := io.ReadAll(file)
	if err != nil {
		return uploadResponse, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	contentType := fileHeader.Header.Get("Content-Type")
	base64Image := base64.StdEncoding.EncodeToString(byteFile)

	entity := File{
		AccountId:   account.Id,
		Type:        "PROFILE_PICTURE",
		ContentType: contentType,
		File:        base64Image,
	}

	idFile, err := svc.repository.Insert(entity)
	if err != nil {
		return uploadResponse, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	uploadResponse = UploadResponseDto{
		IdFile:    idFile,
		AccountId: account.Id,
		Type:      contentType,
	}

	return uploadResponse, nil
}

func (svc serviceImpl) proceedGetFile(queryParam FileQueryParam) (byteFile []byte, contentType string, errWrap *wraped_error.Error) {
	entity, err := svc.repository.Get(queryParam)
	if err != nil {
		if err == pg.ErrNoRows {
			return byteFile, contentType, wraped_error.WrapError(fmt.Errorf("file not found"), http.StatusBadRequest)
		}
		return byteFile, contentType, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	byteFile, err = base64.StdEncoding.DecodeString(entity.File)
	if err != nil {
		return byteFile, contentType, wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	contentType = entity.ContentType
	return byteFile, contentType, nil
}
