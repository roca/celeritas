package celeritas

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
)

func (c *Celeritas) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("body must only contain a single JSON object")
	}

	return nil
}

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

func (c *Celeritas) WriteXML(w http.ResponseWriter, status int, payload interface{}, headers ...http.Header) error {
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	http.ServeFile(w, r, fileToServer)
	return nil
}

// 404 NotFound Error
func (c *Celeritas) Error404(w http.ResponseWriter, r *http.Request) {
	c.ErrorStatus(w, http.StatusNotFound)
}

// 500 InternalServer Error
func (c *Celeritas) Error500(w http.ResponseWriter, r *http.Request) {
	c.ErrorStatus(w, http.StatusInternalServerError)
}

// 400 BadRequest Error
func (c *Celeritas) Error400(w http.ResponseWriter, r *http.Request) {
	c.ErrorStatus(w, http.StatusBadRequest)
}

// 401 Unauthorized Error
func (c *Celeritas) ErrorUnauthorized(w http.ResponseWriter, r *http.Request) {
	c.ErrorStatus(w, http.StatusUnauthorized)
}

// 403 Forbidden Error
func (c *Celeritas) ErrorForbidden(w http.ResponseWriter, r *http.Request) {
	c.ErrorStatus(w, http.StatusForbidden)
}

func (c *Celeritas) ErrorStatus(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
