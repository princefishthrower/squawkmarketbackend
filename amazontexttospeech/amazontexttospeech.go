package amazontexttospeech

import (
	"io"
	"squawkmarketbackend/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"

	"fmt"
)

// TextToSpeech converts text to speech using Amazon Polly
func TextToSpeech(text string) ([]byte, error) {
	sess := session.Must(session.NewSession())
	svc := polly.New(sess)

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String("<speak><prosody rate=\"fast\">" + utils.EscapeForSSML(text) + "</prosody></speak>"),
		VoiceId:      aws.String("Amy"),
		TextType:     aws.String("ssml"),
		// british language code
		LanguageCode: aws.String("en-GB"),
	}

	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	mp3data, err := io.ReadAll(output.AudioStream)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return mp3data, nil
}
