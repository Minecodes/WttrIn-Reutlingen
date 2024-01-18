# WttrIn-Reutlingen
A Mastodon bot that sends a weather forecast with wttr.in

[DE](https://git.minecodes.de/thies/WttrIn-Reutlingen/src/branch/main/README-DE.md)

---------------------------------

## Requirements
- [Go](https://golang.org/)
- A Mastodon account (with an application)

## Application permissions
- read:accounts
- write:media (isn't used at the moment because wttr.in disabled the png support)
- write:statuses

---------------------------------

## Installation
To configure the bot, just rename the `config.example.json` file to `config.json` and fill out the values with the values of your mastodon application (Settings -> Development -> New application) and start it.<br/>
If you want to change the language of the bot, just change the values in the `main.go` file.<br/>
If the bot should run everyday, then you can do it with a cron file, and by compiling the bot to an executable.

Compile:
```bash
go build main.go
```

Run (Linux/MacOS):
```bash
./main
```

Run (Windows/Powershell):
```powershell
.\main.exe
```

---------------------------------

## Licenses
- [wttr.in](https://github.com/chubin/wttr.in) - Apache 2.0
- [go-mastodon](https://github.com/mattn/go-mastodon) - MIT