package main

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
	"strings"
)

// ExampleDownload : Example code for how to use this package for download video.
func Download(videoID string) bool {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		fmt.Println(err)
		return false
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stream.Close()

	file, err := os.Create("video.mp4")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Example usage for playlists: downloading and checking information.
func ExamplePlaylist(playlistID string) {
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		panic(err)
	}

	/* ----- Enumerating playlist videos ----- */
	header := fmt.Sprintf("Playlist %s by %s", playlist.Title, playlist.Author)
	println(header)
	println(strings.Repeat("=", len(header)) + "\n")

	for k, v := range playlist.Videos {
		fmt.Printf("(%d) %s - '%s'\n", k+1, v.Author, v.Title)
	}

	/* ----- Downloading the 1st video ----- */
	entry := playlist.Videos[0]
	video, err := client.VideoFromPlaylistEntry(entry)
	if err != nil {
		panic(err)
	}
	// Now it's fully loaded.

	fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)

	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = io.Copy(file, stream)

	if err != nil {
		panic(err)
	}

	println("Downloaded /video.mp4")
}
