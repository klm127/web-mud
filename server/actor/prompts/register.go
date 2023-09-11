package prompts

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/server/actor/base"
	"github.com/pwsdc/web-mud/world"
)

var RegisterQuestions *base.Interrogator

var LoggedOutCommands *base.CommandSet
var LoggedInCommands *base.CommandSet

var unameRe regexp.Regexp

func init() {
	ib := base.NewInter().Intro("Let's register you an account.")
	qname := base.NewQuestion().Key("username").Text("What shall your name be?").Validator(usernameValidator)
	qpassword := base.NewQuestion().Key("password").Text("What shall your password be?").Validator(passwordValidator)
	ib.AddQuestion(qname.Get())
	ib.AddQuestion(qpassword.Get())
	ib.OnFinished(onRegisterSubmit)
	RegisterQuestions = ib.Get()
	unameRe = *regexp.MustCompile("^[a-zA-Z]+$")
}

func onRegisterSubmit(actor *base.Actor, qr *base.QuestionResult) {

	r_map := qr.GetResult()
	uname, ok := r_map["username"]
	if !ok {
		actor.ErrorMessage("I lost your name in a fumble. Lets try again.")
		actor.StartQuestioning(RegisterQuestions)
		return
	}
	pw, ok := r_map["password"]
	if !ok {
		actor.ErrorMessage("I lost your password somehow. Lets try again.")
		actor.StartQuestioning(RegisterQuestions)
		return
	}
	unique := db.Store.UniqueName(uname)
	if !unique {
		actor.ErrorMessage(fmt.Sprintf("I already know someone named %s. Do you have something else you go by?", uname))
		actor.StartQuestioning(RegisterQuestions)
		return
	}

	being := world.NewBeing(uname, "", arg.Config.World.StartRoom(), nil)
	if being == nil {
		actor.ErrorMessage("I couldn't find a body for you to inhabit, sorry!")
		return
	}
	cuserparams := dbg.CreateUserParams{
		Name:     uname,
		Password: pw,
		Level:    dbg.MudUserlevelPlayer,
		Being:    *being,
	}
	dbuser, err := db.Store.Query.CreateUser(context.Background(), &cuserparams)
	if err != nil {
		actor.ErrorMessage("For some reason I couldn't remember you.")
		world.DeleteBeing(*being)
		actor.StartQuestioning(RegisterQuestions)
		return
	}
	actor.LoadUser(&dbuser)
	actor.SetCommandGroup("user", LoggedInCommands)
	actor.MessageSimple(fmt.Sprintf("It's a pleasure to meet you, %s!", uname))
	actor.EndQuestioning()
}

func usernameValidator(text string) (bool, string) {
	text = strings.TrimSpace(text)
	if len(text) < 4 {
		return false, "Your name must be at least 4 characters."
	}
	if len(text) > 20 {
		return false, "Your name is too long. It can be, at most, 20 characters."
	}
	matches := unameRe.Match([]byte(text))
	if matches {
		return true, ""
	} else {
		return false, "Your name had some strange characters I didn't understand."
	}
}

func passwordValidator(text string) (bool, string) {
	if len(text) < 6 {
		return false, "Your password must be at least 6 characters."
	}
	return true, ""
}
