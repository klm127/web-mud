package commands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/world/sound"
)

var VoiceCommands iactor.ICommandGroup

func init() {
	VoiceCommands = command.NewCommandGroup("voice")
	VoiceCommands.RegisterCommand(command.NewCommand().Name("say").Desc("uses your mouth to say something.").OnExec(being(say)).Get())
	VoiceCommands.RegisterCommand(command.NewCommand().Name("whisper").Desc("uses your mouth to whisper something.").OnExec(being(whisper)).Get())
	VoiceCommands.RegisterCommand(command.NewCommand().Name("yell").Desc("uses your mouth to yell something.").OnExec(being(yell)).Get())
}

func say(actor iactor.IActor, msg string) {
	a_sound := sound.New(actor.Being(), msg, 14)
	actor.Being().GetRoom().SoundEmit(a_sound)
}

func whisper(actor iactor.IActor, msg string) {
	a_sound := sound.New(actor.Being(), msg, 4)
	actor.Being().GetRoom().SoundEmit(a_sound)
}

func yell(actor iactor.IActor, msg string) {
	a_sound := sound.New(actor.Being(), msg, 20)
	actor.Being().GetRoom().SoundEmit(a_sound)
}
