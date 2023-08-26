# WttrIn-Reutlingen
Ein Mastodon/Misskey Bot, der die Wettervohersage mithilfe von wttr.in sendet.

---------------------------------

Um den Bot zu konfigurieren, benenne `config.example.json` in `config.json` um, fülle diese mit den Werten für die spezifische Platform aus und aktiviere diese.<br/>
Wenn du die Sprache des Bots ändern möchtest, änder einfach die Werte in der `main.go` Datei.<br/>
Wenn der Bot jeden Tag laufen soll, erstelle eine Cron Datei und kompiliere den Bot zu einer ausführbaren Datei.

Kompilieren:
```bash
go build main.go
```

Ausführen (Linux/MacOS):
```bash
./main
```

Ausführen (Windows/Powershell):
```powershell
.\main.exe
```

---------------------------------

Lizensen:
- [wttr.in](https://github.com/chubin/wttr.in) - Apache 2.0
- [go-mastodon](https://github.com/mattn/go-mastodon) - MIT
- [go-misskey](https://github.com/mattn/go-misskey) - GPL-3.0