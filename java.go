package main

import (
	"fmt"
	"os"
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

	// if dataDir/java{version}.zip does not exist, download it
	if _, err := os.Stat(dataDir + "/java" + versionStr + ".zip"); os.IsNotExist(err) {
		fmt.Println("[SkaiaCraft] Downloading Java")
		err := download(dataDir+"java"+versionStr+".zip", url)
		if err != nil {
			return err
		}
	}
	// if dataDir/java/{version} does not exist, unzip it
	if _, err := os.Stat(dataDir + "/java/" + versionStr); os.IsNotExist(err) {
		fmt.Println("[SkaiaCraft] Unzipping Java", versionStr)
		err := unzip(dataDir+"java"+versionStr+".zip", dataDir+"/java/"+versionStr)
		if err != nil {
			return err
		}
	}

	return nil
}
