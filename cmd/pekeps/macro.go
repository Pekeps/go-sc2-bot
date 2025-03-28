package main

func (bot *bot) macro() {
	bot.Hub.StepManagers(bot.loop)
}
