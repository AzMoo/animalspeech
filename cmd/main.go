// You can edit this code!
// Click here and start typing.
package main

import (
	"github.com/AzMoo/animalspeech/domain"
	"github.com/AzMoo/animalspeech/repo"
	"github.com/AzMoo/animalspeech/usecase"
)

func main() {
	fileRecorder := repo.FileRecorder{}

	catRepo := repo.CatRepo{}
	dogRepo := repo.DogRepo{}

	catRecorder := usecase.CatRecorder{
		CatFetcher: catRepo,
		Recorder:   fileRecorder,
	}
	catRecorder.RecordSpeech("123")

	dogSpeechRecorder := usecase.SpeechRecorder{
		SpeakerFetcher: func(id string) usecase.Speaker {
			return dogRepo.Fetch(id)
		},
		Recorder: fileRecorder,
	}
	dogSpeechRecorder.RecordSpeech("123")

	catSpeechRecorder := usecase.SpeechRecorder{
		SpeakerFetcher: func(id string) usecase.Speaker {
			return catRepo.Fetch(id)
		},
		Recorder: fileRecorder,
	}
	catSpeechRecorder.RecordSpeech("123")

	genericCatSpeechRecorder := usecase.GenericSpeechRecorder[domain.Cat]{
		SpeakerFetcher: catRepo,
		Recorder:       fileRecorder,
	}
	genericCatSpeechRecorder.RecordSpeech("123")

	genericDogSpeechRecorder := usecase.GenericSpeechRecorder[domain.Dog]{
		SpeakerFetcher: dogRepo,
		Recorder:       fileRecorder,
	}
	genericDogSpeechRecorder.RecordSpeech("123")
}
