package basecommands

import (
	"context"
	"fmt"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/server/user/actor/prompt"
)

var UserLoggedOutCommands iactor.ICommandGroup

var loginQuestions iactor.IInterrogatory
var registerQuestions iactor.IInterrogatory

func init() {

	// build the login prompts
	q_uname := prompt.NewQuestion("username").Text("What are you known as?").Get()
	q_upass := prompt.NewQuestion("password").Text("What is your secret password?").Get()
	loginQuestions = prompt.NewInterrogatory().Intro("Let's get you logged in.").AddQuestions(q_uname, q_upass).OnFinished(onLoginSubmit).Get()

	// build the register prompts
	q_name := prompt.NewQuestion("username").Text("What shall you be known by?").Get()
	q_pass := prompt.NewQuestion("password").Text("What shall your secret password be?").Get()
	registerQuestions = prompt.NewInterrogatory().Intro("").AddQuestions(q_name, q_pass).OnFinished(onRegisterSubmit).Get()

	// the logged out command set
	UserLoggedOutCommands = command.NewCommandGroup("user")

	// build the register command
	register := command.NewCommand().Name("register").Desc("registers an account").OnExec(register).Get()

	// build the login command
	login_c := command.NewCommand().Name("login").Desc("logs in to your account").OnExec(login).Get()

	UserLoggedOutCommands.RegisterCommands(register, login_c)
}

func login(actor iactor.IActor, msg string) {
	actor.StartQuestioning(loginQuestions)
}

func onLoginSubmit(actor iactor.IActor, qr iactor.IQuestionResult) {
	r_map := *qr.GetResult()
	uname, ok := r_map["username"]
	if !ok {
		actor.ErrorMessage("I lost your name in a fumble. Lets try again.")
		actor.StartQuestioning(loginQuestions)
		return
	}
	pw, ok := r_map["password"]
	if !ok {
		actor.ErrorMessage("I lost your password somehow. Lets try again.")
		actor.StartQuestioning(loginQuestions)
		return
	}
	dbuser, err := db.Store.Query.GetUserByName(context.Background(), uname)
	if err != nil {
		actor.Errorf("I couldn't find you, %s. Sorry.", uname)
		actor.StartQuestioning(loginQuestions)
		return
	}
	if dbuser.Password != pw {
		// Todo: encrypt/decrypt pw
		actor.Errorf("I'm not entirely sure you are who you say you are.")
		actor.StartQuestioning(loginQuestions)
		return
	}
	actor.MessageSimplef("Welcome back, %s.", uname)
	actor.SetUser(&dbuser)
	actor.SetCommandGroup(UserLoggedInCommands)
	actor.EndQuestioning()
}

func register(actor iactor.IActor, msg string) {
	actor.StartQuestioning(registerQuestions)
}

func onRegisterSubmit(actor iactor.IActor, qr iactor.IQuestionResult) {

	if qr == nil || actor == nil {
		fmt.Printf("Problem on register submit; nil parameter.")
		return
	}

	r_map := *qr.GetResult()
	uname, ok := r_map["username"]
	if !ok {
		actor.ErrorMessage("I lost your name in a fumble. Lets try again.")
		actor.StartQuestioning(registerQuestions)
		return
	}
	pw, ok := r_map["password"]
	if !ok {
		actor.ErrorMessage("I lost your password somehow. Lets try again.")
		actor.StartQuestioning(registerQuestions)
		return
	}
	unique := db.Store.UniqueName(uname)
	if !unique {
		actor.ErrorMessage(fmt.Sprintf("I already know someone named %s. Do you have something else you go by?", uname))
		actor.StartQuestioning(registerQuestions)
		return
	}
	being, err := db.Store.NewBeingInStartRoom(uname)
	if err != nil {
		actor.ErrorMessage("I couldn't find a body for you to inhabit, sorry!")
		return
	}
	cuserparams := dbg.CreateUserParams{
		Name:     uname,
		Password: pw,
		Level:    dbg.MudUserlevelPlayer,
		Being:    being.ID,
	}
	dbuser, err := db.Store.Query.CreateUser(context.Background(), &cuserparams)
	if err != nil {
		fmt.Println(err.Error())
		actor.ErrorMessage("For some reason I couldn't remember you.")
		db.Store.DeleteBeingEntry(being.ID)
		actor.StartQuestioning(registerQuestions)
		return
	}
	db.Store.SetBeingOwner(being, dbuser.ID)
	actor.SetUser(&dbuser)
	actor.SetCommandGroup(UserLoggedInCommands)
	actor.MessageSimple(fmt.Sprintf("It's a pleasure to meet you, %s!", uname))
	actor.EndQuestioning()

}
