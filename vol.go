package main

import (
	"io/ioutil"
	"log"
	"time"

	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
)

func main() {
	count := 1
	callUuidV4, _ := uuid.NewV4()

	nodeName := os.Getenv("NODE_NAME")
	podName := os.Getenv("POD_NAME")
	podNamespace := os.Getenv("POD_NAMESPACE")
	podIP := os.Getenv("POD_IP")

	serviceAccountName := os.Getenv("SERVICE_ACCOUNT")

	path := os.Getenv("VOL_PATH")
	if path == "" {
		path = "./"
	}

	filePath := fmt.Sprintf("%s/gv_%s_%s.txt", path, nodeName, podName)

	statusMessage := "OK"

	// touch a file
	_, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		statusMessage = err.Error()
	}

	// web API endpoint
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		instanceUuidV4, _ := uuid.NewV4()

		c.JSON(200, gin.H{
			"message":         "go-vol",
			"time":            time.Now(),
			"count":           count,
			"uuid_call":       instanceUuidV4.String(),
			"uuid_instance":   callUuidV4.String(),
			"client_ip":       c.ClientIP(),
			"version":         1,
			"version_msg":     "version 1",
			"node_name":       nodeName,
			"pod_name":        podName,
			"pod_namepage":    podNamespace,
			"pod_ip":          podIP,
			"service_account": serviceAccountName,
			"file":            filePath,
			"files":           listFiles(path),
			"status_message":  statusMessage,
		})
		count++
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func listFiles(path string) []string {
	// get a listing of files in the path
	//
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fileNames := make([]string, 0)

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	return fileNames
}
