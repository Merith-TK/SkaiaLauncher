package main

import "fmt"

func main() {
	err := setupConfig()
	if err != nil {
		fmt.Println("[SkaiaCraft] Error loading config:", err)
		return
	} else {
		fmt.Println("[SkaiaCraft] Config loaded")
	}

	// err = downloadJava(8)
	// if err != nil {
	// 	fmt.Println("[SkaiaCraft] Error downloading Java:", err)
	// } else {
	// 	fmt.Println("[SkaiaCraft] Java downloaded")
	// }

	download("image.png", "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png")
}
