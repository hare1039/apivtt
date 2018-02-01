package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"hash/fnv"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}

func DownloadSub(uri string) error {
	out, err := os.Create(hash(uri) + filepath.Ext(uri))
	defer out.Close()

	if err != nil {
		return err
	}
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)

	return err
}

func Convert(c *gin.Context) {
	src := c.Query("src")
	fmt.Println("getting from " + src)
	err := DownloadSub(src)
	defer os.Remove(hash(src) + filepath.Ext(src))
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := exec.Command(
		"ffmpeg",
		"-hide_banner",
		"-loglevel",
		"panic",
		"-i", hash(src)+filepath.Ext(src),
		"-f", "webvtt",
		"-").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	c.Render(http.StatusOK, render.Data{
		ContentType: "mime/vtt",
		Data:        result,
	})
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	r := gin.Default()
	r.GET("/convert-to-vtt", Convert)
	r.Run(":39001")
}
