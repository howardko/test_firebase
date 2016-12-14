package main

import (
	"wordsmap/common/util"
	"runtime"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"os"
	"encoding/json"
	"fmt"
	"wordsmap/notify"
	"wordsmap/common/Error"
)

// Version of wordsmap
type Version struct {
	ReleaseVersion string `json:"release_version"`
	ReleaseTime    string `json:"release_time"`
}

var version = Version{
	ReleaseVersion: "dev",
	ReleaseTime:    time.Now().String(),
}

func init() {
	versionPath := "conf/version.json"
	versionFile, err := os.Open(versionPath)
	if err != nil {
		fmt.Println("open config file failed. Use defulat config")
	} else {
		jsonParser := json.NewDecoder(versionFile)
		if err = jsonParser.Decode(&version); err != nil {
			fmt.Println("parsing config file failed.")
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	util.SetupConfig("conf/config.json")

	r := gin.New()
	r.NoRoute(func(c *gin.Context) {
		noRoute := Error.UnknownResource
		c.JSON(noRoute.Status, gin.H{"code":noRoute.Code, "message": noRoute.Message})
	})

	r.GET("/v0.1/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, version)
	})

	v0_1 := r.Group("/v0.1")
	notify.SetupRoute(v0_1)

	listenPort := util.Config.Server.Port
    	r.Run(listenPort)
}
