package videostreaming

import (
	"fmt"

	"os/exec"
)

const (
	sampleRateHertz = 16000
	languageCode    = "en-US"
)

// this function takes an m3u8 video stream URI and the number of seconds to transcribe, i.e.
// https://manifest.googlevideo.com/api/manifest/hls_playlist/expire/1677607017/ei/Cez9Y_yNIJKU8gPJo7KQDw/ip/2a05:1141:1ed:9300:a0b4:e577:7ae:c583/id/v4xqUHoB1GE.1/itag/91/source/yt_live_broadcast/requiressl/yes/ratebypass/yes/live/1/sgoap/gir%3Dyes%3Bitag%3D139/sgovp/gir%3Dyes%3Bitag%3D160/hls_chunk_host/rr5---sn-1gi7znes.googlevideo.com/playlist_duration/30/manifest_duration/30/spc/H3gIhtwdAhvffxsZvLtIHXyjgjg-oKE/vprv/1/playlist_type/DVR/initcwndbps/3071250/mh/Ed/mm/44/mn/sn-1gi7znes/ms/lva/mv/m/mvi/5/pl/43/dover/11/pacing/0/keepalive/yes/fexp/24007246/mt/1677584917/sparams/expire,ei,ip,id,itag,source,requiressl,ratebypass,live,sgoap,sgovp,playlist_duration,manifest_duration,spc,vprv,playlist_type/sig/AOq0QJ8wRgIhAI_nPte5TC8P0GFLv1NPCvd-k8gpg1BsKbFefCLFvVmNAiEAxZ0ooO89RRPiN6zmZiedqZPNIVlHHbuU7urwlEw_RTY%3D/lsparams/hls_chunk_host,initcwndbps,mh,mm,mn,ms,mv,mvi,pl/lsig/AG3C_xAwRQIhAJ2UxYZa7GY8zb3ed3gsDt0ng-ACDFyvZelqASvsyOEoAiAtkh_wOVhfcWpZZ7o0SgW8JnQxTMIGsEAddiG5TT5CdA%3D%3D/playlist/index.m3u8
func M3U8VideoStreamToFile(m3u8VideoStreamURI string, fileName string, fileType string, secondsString string) error {
	// write to flac file
	cmd := exec.Command("ffmpeg", "-y", "-i", m3u8VideoStreamURI, "-t", secondsString, "-f", fileType, "-ar", fmt.Sprint(sampleRateHertz), fileName)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
