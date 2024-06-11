package main

import (
	"fmt"
	"image"
	imageprocess "images/image_process"
	"strings"
)

type Job struct {
	InputPath  string
	Image      image.Image
	OutPutPath string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, p := range paths {
			job := Job{
				InputPath:  p,
				OutPutPath: strings.Replace(p, "images", "images/output", 1)}
			job.Image = imageprocess.ReadImage(p)

			out <- job
		}

		close(out)

	}()
	return out
}

func resizeImage(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		for job := range input {

			job.Image = imageprocess.Resize(job.Image)

			out <- job
		}

		close(out)

	}()
	return out
}

func convertToGrayScale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {

			job.Image = imageprocess.GrayScale(job.Image)

			out <- job
		}

		close(out)

	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)

	go func() {
		for job := range input {
			imageprocess.WriteImage(job.OutPutPath, job.Image)
			out <- true
		}
		close(out)
	}()

	return out
}

func main() {
	imagePaths := []string{"images/image.jpg"}

	channel1 := loadImage(imagePaths)
	channel2 := resizeImage(channel1)
	channel3 := convertToGrayScale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("success!")
		} else {
			fmt.Println("failed!")
		}
	}
}
