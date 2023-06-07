package videototext

import (
	"log"
	"squawkmarketbackend/open_ai"
	"squawkmarketbackend/streamfinder"
	"squawkmarketbackend/videostreaming"
)

func YoutubeVideoIdToText(youtubeVideoId string) error {
	// get m3u8 video stream url
	videoStreamUrl, err := streamfinder.GetStreamUrlFromYoutubeVideoId(youtubeVideoId)
	if err != nil {
		return err
	}

	// filename for now is more or less an intermediate, may want to change later
	fileName := "audio.mp3"

	// convert video stream to audio file
	err = videostreaming.M3U8VideoStreamToFile(videoStreamUrl, fileName, "mp3", "10")
	if err != nil {
		return err
	}

	// convert audio file to text
	// open ai cost is $0.006 per minute of audio - i.e. $0.36 per hour of audio
	text, err := open_ai.SpeechToText(fileName)
	if err != nil {
		return err
	}

	// print text to log for now
	log.Println(*text)
	return nil
}

// recursive function to continuously transcribe a youtube video
// transcript prints out to log
func ContinuouslyTranscribeYoutubeStream(youtubeVideoId string) {
	log.Println("transcribing for " + youtubeVideoId)
	err := YoutubeVideoIdToText(youtubeVideoId)
	if err != nil {
		panic(err)
	}

	ContinuouslyTranscribeYoutubeStream(youtubeVideoId)
}
