package main

import (
	"fmt"
	"strings"
)

var (
	javaUrl = "https://api.adoptopenjdk.net/v3/binary/latest/{version}/ga/windows/x64/jdk/hotspot/normal/adoptopenjdk?project=jdk"
)

// download java8
func downloadJava(version int) error {
	// convert version to string
	versionStr := fmt.Sprintf("%d", version)
	url := strings.Replace(javaUrl, "{version}", versionStr, -1)
	// download java
	fmt.Println("[SkaiaCraft] Downloading Java")
	err := download("java"+versionStr+".zip", url)
	if err != nil {
		return err
	}
	return nil
}
