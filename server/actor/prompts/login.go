package prompts

import (
	"context"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/server/actor/base"
)

var LoginQuestions *base.Interrogator

func init() {

	ilogin := base.NewInter().Intro("Let's get you logged in.")
	qnamelogin := base.NewQuestion().Key("username").Text("What's your name?").Validator(usernameValidator)
	qpasswordlogin := base.NewQuestion().Key("password").Text("What's your password?").Validator(passwordValidator)
	ilogin.AddQuestion(qnamelogin.Get())
	ilogin.AddQuestion(qpasswordlogin.Get())
	ilogin.OnFinished(onLoginSubmit)
	LoginQuestions = ilogin.Get()

}

func onLoginSubmit(actor *base.Actor, qr *base.QuestionResult) {
	r_map := qr.GetResult()
	uname, ok := r_map["username"]
	if !ok {
		actor.ErrorMessage("I lost your name in a fumble. Lets try again.")
		actor.StartQuestioning(LoginQuestions)
		return
	}
	pw, ok := r_map["password"]
	if !ok {
		actor.ErrorMessage("I lost your password somehow. Lets try again.")
		actor.StartQuestioning(LoginQuestions)
		return
	}
	dbuser, err := db.Store.Query.GetUserByName(context.Background(), uname)
	if err != nil {
		actor.Errorf("I couldn't find you, %s. Sorry.", uname)
		actor.StartQuestioning(LoginQuestions)
		return
	}
	if dbuser.Password != pw {
		actor.Errorf("I'm not entirely sure you are who you say you are.")
		actor.StartQuestioning(LoginQuestions)
		return
	}
	actor.LoadUser(&dbuser)
	actor.SetCommandGroup("user", LoggedInCommands)
	actor.MessageSimplef("Welcome back, %s.", uname)
	actor.EndQuestioning()
}
