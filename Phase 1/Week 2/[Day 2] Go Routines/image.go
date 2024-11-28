package main

import (
	"fmt"
	"time"
)

func processImage(imageURL string) {
	fmt.Printf("Processing image %s\n", imageURL)

	//simulate image processing
	time.Sleep(3 * time.Second)
	fmt.Printf("Image processing completed %s\n", imageURL)
}

func lain() {
	images := []string{
		"https://example.com/image1.jpg",
		"https://example.com/image2.jpg",
		"https://example.com/image3.jpg",
		"https://example.com/image4.jpg",
		"https://example.com/image5.jpg",
	}

	for _, imageURL := range images {
		go processImage(imageURL)
	}

	fmt.Println("Image processing started, main application continues ...")

	time.Sleep(5 * time.Second)

	fmt.Println("All image processing completed")
}
