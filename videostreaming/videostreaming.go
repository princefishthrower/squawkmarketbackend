package videostreaming

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"
)

const (
	sampleRateHertz = 16000
	languageCode    = "en-US"
)

// this function requires an m3u8 video stream and the number of seconds to transcribe
// to get an m3u8 video stream, you can use a tool like yt-dlp:
// yt-dlp --list-formats https://www.youtube.com/watch\?v\=v4xqUHoB1GE
// this will give a list of streams, best to take the smallest one for fastest transmission since we only want the audio
// yt-dlp -f 91 -g  https://www.youtube.com/watch\?v\=v4xqUHoB1GE
// this will give the m3u8 video stream URL

// https://manifest.googlevideo.com/api/manifest/hls_playlist/expire/1677607017/ei/Cez9Y_yNIJKU8gPJo7KQDw/ip/2a05:1141:1ed:9300:a0b4:e577:7ae:c583/id/v4xqUHoB1GE.1/itag/91/source/yt_live_broadcast/requiressl/yes/ratebypass/yes/live/1/sgoap/gir%3Dyes%3Bitag%3D139/sgovp/gir%3Dyes%3Bitag%3D160/hls_chunk_host/rr5---sn-1gi7znes.googlevideo.com/playlist_duration/30/manifest_duration/30/spc/H3gIhtwdAhvffxsZvLtIHXyjgjg-oKE/vprv/1/playlist_type/DVR/initcwndbps/3071250/mh/Ed/mm/44/mn/sn-1gi7znes/ms/lva/mv/m/mvi/5/pl/43/dover/11/pacing/0/keepalive/yes/fexp/24007246/mt/1677584917/sparams/expire,ei,ip,id,itag,source,requiressl,ratebypass,live,sgoap,sgovp,playlist_duration,manifest_duration,spc,vprv,playlist_type/sig/AOq0QJ8wRgIhAI_nPte5TC8P0GFLv1NPCvd-k8gpg1BsKbFefCLFvVmNAiEAxZ0ooO89RRPiN6zmZiedqZPNIVlHHbuU7urwlEw_RTY%3D/lsparams/hls_chunk_host,initcwndbps,mh,mm,mn,ms,mv,mvi,pl/lsig/AG3C_xAwRQIhAJ2UxYZa7GY8zb3ed3gsDt0ng-ACDFyvZelqASvsyOEoAiAtkh_wOVhfcWpZZ7o0SgW8JnQxTMIGsEAddiG5TT5CdA%3D%3D/playlist/index.m3u8
func M3U8VideoStreamToText(m3u8VideoStream string, secondsString string) (*string, error) {
	file := "audio.flac"
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	// write to flac file
	cmd := exec.Command("ffmpeg", "-y", "-i", m3u8VideoStream, "-t", secondsString, "-f", "flac", "-ar", fmt.Sprint(sampleRateHertz), file)
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// read in the flac file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Receive the transcript from the API
	// Send audio data to the Speech API
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:          speechpb.RecognitionConfig_FLAC,
			SampleRateHertz:   sampleRateHertz,
			LanguageCode:      languageCode,
			AudioChannelCount: 2,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	}
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		fmt.Printf("Failed to create streaming recognizer: %v", err)
		return nil, err
	}

	// Get the transcript from the response
	transcript := ""
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Transcript: %v\n", alt.Transcript)
			transcript = alt.Transcript
		}
	}

	return &transcript, nil
}
