# WttrIn-Reutlingen
Ein Mastodon Bot, der die Wettervohersage mithilfe von wttr.in sendet.

---------------------------------

Um den Bot zu konfigurieren, erstelle einfach eine `config.json` Datei und fülle sie mit den in der Beispiels Config gezeigten Werten aus.<br/>
Wenn du die Sprache des Bots ändern möchtest, änder einfach die Werte in der `main.go` Datei.<br/>
Wenn der Bot jeden Tag laufen soll, erstelle eine Cron Datei und kompiliere den Bot zu einer ausführbaren Datei.

Kompilieren:
```bash
go build main.go
```

---------------------------------

Lizensen:
- [wttr.in](https://github.com/chubin/wttr.in) - Apache 2.0
- [go-mastodon](https://github.com/mattn/go-mastodon) - MIT