package files

import (
	"mime/multipart"

	"lending-service/account"
	"lending-service/utility/wraped_error"
)

type Service interface {
	proceedSaveFile(file multipart.File, fileHeader *multipart.FileHeader, account account.Account) (uploadResponse UploadResponseDto, errWrap *wraped_error.Error)
	proceedGetFile(queryParam FileQueryParam) (byteFile []byte, contentType string, errWrap *wraped_error.Error)
}
