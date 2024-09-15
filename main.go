package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	var videoID string
	fmt.Println("Enter a YouTube URL | ID : ")
	fmt.Scanln(&videoID)

	client := youtube.Client{}
	video, err := client.GetVideo(videoID)

	vidTitle := fmt.Sprintf("Playlist %s by %s", video.Title, video.Author)

	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create(vidTitle + ".mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
