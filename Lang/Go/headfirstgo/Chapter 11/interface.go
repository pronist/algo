package main

type NoiseMaker interface {
	MakeSound()
}

type (
	Whistle string
	Horn    string
	Robot   string
)

func (w Whistle) MakeSound() {}
func (h Horn) MakeSound()    {}
func (r Robot) MakeSound()   {}

func (r Robot) Walk() {}

func main() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}
