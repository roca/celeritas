package celeritas

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
)

func (c *Celeritas) WriteJSON(w http.ResponseWriter, status int, payload interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) WriteJXML(w http.ResponseWriter, status int, payload interface{}, headers ...http.Header) error {
	out, err := xml.MarshalIndent(payload, "", "   ")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) DownloadFile(w http.ResponseWriter, r *http.Request, pathToFile, filename string) error {
	fp := path.Join(pathToFile, filename)
	fileToServer := filepath.Clean(fp)
	w.Header().Set("Content-Type", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	http.ServeFile(w, r, fileToServer)
	return nil
}
