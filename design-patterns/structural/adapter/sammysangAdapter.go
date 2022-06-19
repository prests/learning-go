package main

type sammysangAdapter struct {
	sstv *SammysangTV
	// TODO: add a field for the SammysangTV reference
}

func (ss *sammysangAdapter) turnOn() {
	ss.sstv.setOnState(true)
	// TODO
}

func (ss *sammysangAdapter) turnOff() {
	ss.sstv.setOnState(false)
	// TODO
}

func (ss *sammysangAdapter) volumeUp() int {
	vol := ss.sstv.getVolume() + 1
	ss.sstv.setVolume(vol)
	return vol
	// TODO
}

func (ss *sammysangAdapter) volumeDown() int {
	vol := ss.sstv.getVolume() - 1
	ss.sstv.setVolume(vol)
	return vol
	// TODO
}

func (ss *sammysangAdapter) channelUp() int {
	channel := ss.sstv.getChannel() + 1
	ss.sstv.setChannel(channel)
	return channel
	// TODO
}

func (ss *sammysangAdapter) channelDown() int {
	channel := ss.sstv.getChannel() - 1
	ss.sstv.setChannel(channel)
	return channel
	// TODO
}

func (ss *sammysangAdapter) goToChannel(ch int) {
	ss.sstv.setChannel(ch)
	// TODO
}
