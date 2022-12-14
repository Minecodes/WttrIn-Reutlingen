# WttrIn-Reutlingen
A Mastodon bot that sends a weather forecast with wttr.in

[DE](https://github.com/Minecodes/blob/main/README-DE.md)

---------------------------------

To configure the bot, just create a `config.json` file and fill it out with the values that you have out of the example file.<br/>
If you want to change the language of the bot, just change the values in the `main.go` file.<br/>
When the bot should run everyday, then you can do it with a cron file and by compiling the bot to an executable.

Compile:
```bash
go build main.go
```

---------------------------------

Licenses:
- [wttr.in](https://github.com/chubin/wttr.in) - Apache 2.0
- [go-mastodon](https://github.com/mattn/go-mastodon) - MIT