package server

import (
	"bytes"
	"github.com/drone/routes"
	"github.com/masayukioguni/go-webp-server/resize"
	"github.com/masayukioguni/go-webp-server/webp"

	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

type Server struct {
	ImageHost string
}

func (s *Server) IsValidImagePath(imagePath string) bool {
	if imagePath == "" {
		return false
	}

	ext := filepath.Ext(imagePath)

	if ext == ".jpg" {
		return true
	}
	return false
}

func (s *Server) WebpHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	imagePath := params.Get(":imagePath")

	if !s.IsValidImagePath(imagePath) {
		http.Error(w, "imagePath not specified or invalid", http.StatusBadRequest)
		return
	}

	imageURL := s.ImageHost + imagePath

	resp, err := http.Get(imageURL)

	if err != nil {
		http.Error(w, "http.Get() "+err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "http.Get() Statuscode:"+resp.Status, resp.StatusCode)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "ioutil.ReadAll() "+err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := webp.Decode(bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "webp.Decode() "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = webp.Encode(w, m, &webp.Options{false, 50})
	if err != nil {
		http.Error(w, "webp.Encode() "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) ResizeHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	widthParam := params.Get(":width")
	heightParam := params.Get(":height")

	width, err := strconv.ParseUint(widthParam, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v strconv.Atoi(%v)", err.Error(), widthParam)
		return
	}

	if width < 0 || width > 640 {
		http.Error(w, "height not specified or invalid", http.StatusBadRequest)
		return
	}

	height, err := strconv.ParseUint(heightParam, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v strconv.Atoi(%v)", err.Error(), heightParam)
		return
	}

	if height < 0 || height > 480 {
		http.Error(w, "height not specified or invalid", http.StatusBadRequest)
		return
	}

	imagePath := params.Get(":imagePath")

	if !s.IsValidImagePath(imagePath) {
		http.Error(w, "imagePath not specified or invalid", http.StatusBadRequest)
		return
	}

	imageUrl := s.ImageHost + imagePath

	resp, err := http.Get(imageUrl)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s http.Get(%s)", err.Error(), imageUrl)
		return
	}

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		fmt.Fprintf(w, "%s %v", resp.Status, imageUrl)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	m, err := resize.Resize(bytes.NewBuffer(body), uint(width), uint(height))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = webp.Encode(w, m, &webp.Options{false, 50})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
}

func (s *Server) Route(m *routes.RouteMux) {
	m.Get("/webp/:imagePath(.*)", s.WebpHandler)
	m.Get("/:width/:height/:imagePath(.*)", s.ResizeHandler)
}

func (s *Server) Run() {

	mux := routes.New()
	s.Route(mux)

	http.Handle("/", mux)

	http.ListenAndServe(":8088", nil)

}

func NewServer(imageHost string) *Server {
	server := &Server{
		ImageHost: imageHost,
	}
	return server
}
