package go_api_file_picker

import (
	"io/fs"
	"net/http"
	"path/filepath"
)

type FilePicker struct {
}

type FileInfo struct {
}

func (picker *FilePicker) GetFileList(responseWriter http.ResponseWriter, request *http.Request) {
	var directory = request.URL.Query().Get("directory")
	filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {

		return nil
	})
}
