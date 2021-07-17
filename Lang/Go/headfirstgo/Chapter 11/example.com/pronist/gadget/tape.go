package gadget

type TapePlayer struct {
	Batteries string
}

type Player interface {
	Play(string)
	Stop()
}

func (t TapePlayer) Play(song string) {}
func (t TapePlayer) Stop()            {}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {}
func (t TapeRecorder) Record()          {}
func (t TapeRecorder) Stop()            {}
