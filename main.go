package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	// If go comes new version, you can change default ver to new version
	versionFlag := flag.String("version", "1.22.2", "Go version to install")

	flag.Parse()

	dockerfileContent := fmt.Sprintf(`
    FROM alpine:latest

    RUN apk update && \
        apk add --no-cache curl

    WORKDIR /tmp

    RUN curl -LO https://golang.org/dl/go%s.linux-amd64.tar.gz && \
        tar -C /usr/local -xzf go%s.linux-amd64.tar.gz && \
        rm go%s.linux-amd64.tar.gz

    ENV PATH="/usr/local/go/bin:${PATH}"

    ENV GOPATH /go
    ENV PATH $PATH:/go/bin:$GOPATH/bin

    # If you want to use GCC for CGO, turn this on to 1. 
    ENV CGO_ENABLED 0

    CMD ["/bin/sh"]
    `, *versionFlag, *versionFlag, *versionFlag)

	err := ioutil.WriteFile("Dockerfile", []byte(dockerfileContent), 0644)
	if err != nil {
		fmt.Println("Error creating Dockerfile:", err)
		return
	}
	fmt.Println("Dockerfile created successfully.")

	fmt.Println("Building Docker image...")
	cmd := exec.Command("docker", "build", "-t", "godocker_image", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building Docker image:", err)
		return
	}
	fmt.Println("Docker image built successfully.")
}
