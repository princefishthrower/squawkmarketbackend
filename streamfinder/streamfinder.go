package streamfinder

import "os/exec"

// bloomberg global financial news - dp8PhLsUcFE
// march 4th nbc (nbc changes daily unfortunately ) - uVlW3q68wAQ
// to get an m3u8 video stream, you can use a tool like yt-dlp:
// yt-dlp --list-formats https://www.youtube.com/watch\?v\=dp8PhLsUcFE
// this will give a list of streams, best to take the smallest one for fastest transmission since we only want the audio
// yt-dlp -f 91 -g  https://www.youtube.com/watch\?v\=dp8PhLsUcFE
// this will give the m3u8 video stream URL
// seems like number 91 is fairly consistently the smallest one
func GetStreamUrlFromYoutubeVideoId(youtubeVideoId string) (string, error) {
	// use yt-dlp to get the stream url
	cmd := exec.Command("yt-dlp", "-f", "91", "-g", "https://www.youtube.com/watch?v="+youtubeVideoId)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
