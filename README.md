# WttrIn-Reutlingen
A Bluesky bot that posts the weather forecast of Reutlingen at 7:00 am every day.

---------------------------------

## Requirements
- [bun](https://bun.sh)
- [NodeJS](https://nodejs.org)
- A Bluesky account (please create a new account for the bot)

---------------------------------

## Installation
1. Clone this repository
2. Install the dependencies
```bash
bun install
```
3. Create a new Bluesky account and a app password
4. Create a new file called `.env` and add the following content:
```env
BSKY_USERNAME=username.bsky.network
BSKY_PASSWORD=password
BOT_CITY=Your_City
```
5. Replace `username.bsky.network` with your Bluesky username, `password` with your Bluesky app password and `Your_City` with your city (e.g. Reutlingen)
6. Run the bot
```bash
bun start
```
or use
```bash
docker compose up -d
```

## Change the language
To change the language of the weather forecast, you can change the `lang` parameter in the `main.ts` file. The default language is `en`. Also, you have to manually change the text in the line 28 of this file.

----------------------------------------------------------------

## License
 - This project is licensed under the BSD 3-Clause License - see the LICENSE file for details
 - [wttr.in](https://github.com/chubin/wttr.in/blob/master/LICENSE) is licensed under the Apache License 2.0
 - [dotenv](https://github.com/motdotla/dotenv/blob/master/LICENSE) is licensed under the BSD 2-Clause License
 - [node-cron](https://github.com/kelektiv/node-cron/blob/main/LICENSE) is licensed under the MIT License
 - [@skyware/bot](https://github.com/skyware-js/bot/blob/main/LICENSE) is licensed under the Mozilla Public License 2.0