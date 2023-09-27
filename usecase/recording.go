package usecase

import "github.com/AzMoo/animalspeech/domain"

type recorder interface {
	Record(speech string)
}

// CatRecorder is a usecase that fetches a Cat and uses a recorder to record
// its speech. Since the point of the usecase is to fetch a Cat and record its
// speech, it makes sense to depend explicitly on the Cat entity defined in the
// domain. There shouldn't be any reason to mock or arbitrarily replace the Cat
// entity, so we don't need to define an interface for it.
//
// We may want to mock or arbitrarily replace the catFetcher to change where we're
// fetching a Cat from, so we define an interface for it and inject it.
type CatRecorder struct {
	CatFetcher catFetcher
	Recorder   recorder
}

type catFetcher interface {
	Fetch(id string) domain.Cat
}

func (i CatRecorder) RecordSpeech(id string) {
	cat := i.CatFetcher.Fetch(id)
	i.Recorder.Record(cat.Speak())
}

// SpeechRecorder is a more generic usecase that is intended to fetch any
// entity that can Speak and record its speech. Since multiple entities can
// Speak, we define an interface for a Speaker, and inject a function fetches
// the Speaker. There's a few reasons this isn't ideal though.
//
// It can be easy to introduce non-trivial logic when converting to a Speaker, in which
// case we're introducing complexity outside the usecase. Also by injecting a constructor
// function we're introducing a level of indirection which can be confusing. We can see evidence of
// this getting a bit out of control in app/platform/inspection/legacy/app/folder_access.go:8
// in which we inject a lot of these types of functions, and it can be difficult to follow.
//
// We also need to export the Speaker interface, so we can use it in the
// layer instantiating the usecase. It's not a huge problem, because we
// already depend on the usecase layer to instantiate it, but the more we export
// the more opportunity there is for misuse and accidental coupling.
type SpeechRecorder struct {
	SpeakerFetcher func(id string) Speaker
	Recorder       recorder
}

type Speaker interface {
	Speak() string
}

func (i SpeechRecorder) RecordSpeech(id string) {
	s := i.SpeakerFetcher(id)
	i.Recorder.Record(s.Speak())
}

// GenericSpeechRecorder is a generic usecase that uses generics to define
// the type of speaker it's expecting a fetcher for. This allows us to inject
// a fetcher for any type that satisfies the Speaker interface, without needing
// to define an interface for it explicitly, or inject a constructor function.
//
// It avoids the problems in the SpeechRecorder usecase, in that the speaker
// interface no longer needs to be exported, and because we can inject
// a fetcher instead of a constructor function, we don't run the risk of introducing
// complex logic outside the usecase.
type GenericSpeechRecorder[T speaker] struct {
	SpeakerFetcher genericSpeakerFetcher[T]
	Recorder       recorder
}

type genericSpeakerFetcher[T speaker] interface {
	Fetch(id string) T
}

type speaker interface {
	Speak() string
}

func (i GenericSpeechRecorder[T]) RecordSpeech(id string) {
	s := i.SpeakerFetcher.Fetch(id)
	i.Recorder.Record(s.Speak())
}
