package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := setupConfig()
	if err != nil {
		fmt.Println("[SkaiaCraft] Error loading config:", err)
		return
	} else {
		fmt.Println("[SkaiaCraft] Config loaded")
	}

	// split version string into array
	version := strings.Split(conf.Minecraft.Version, ".")
	// get second element of array
	versionInt, _ := strconv.Atoi(version[1])

	if err != nil {
		fmt.Println("[SkaiaCraft] Error converting version:", err)
		return
	}
	fmt.Println("[SkaiaCraft] Minecraft version:", conf.Minecraft.Version)

	// create dataDir if not exist
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		fmt.Println("[SkaiaCraft] Creating data directory")
		os.Mkdir(dataDir, 0755)
	}
	// if version is 1.13 or higher, download java 17
	if versionInt >= 13 {
		err := downloadJava(17)
		if err != nil {
			fmt.Println("[SkaiaCraft] Error downloading Java:", err)
			return
		}
	} else {
		err := downloadJava(8)
		if err != nil {
			fmt.Println("[SkaiaCraft] Error downloading Java:", err)
			return
		}
	}
}
