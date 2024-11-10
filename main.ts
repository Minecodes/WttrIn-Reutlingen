import { Bot } from "@skyware/bot";
import * as dotenv from "dotenv";
import { CronJob } from "cron";
import * as process from "process";
import axios from "axios";

dotenv.config();

const bot = new Bot({
    emitEvents: false
});


async function main() {
  await bot.login({
    identifier: process.env.BSKY_USERNAME,
    password: process.env.BSKY_PASSWORD,
  });
  const { data } = await axios.get(
    `https://wttr.in/${process.env.BOT_CITY}?format=4&lang=de`,
    {
      headers: {
        "User-Agent": "curl/8.7.1",
      },
    }
  );
  await bot.post({
    text: `Dies ist das Wetter für ${process.env.BOT_CITY}\n${data}`,
  });
  console.log(`Weather report got posted`);
}

/**
 * Cron formating: Minute Hour Day Month DayOfWeek
 *
 * 0 7 * * * - Every day at 7:00 AM
 */
//const scheduleExpressionMinute = "* * * * *";
const scheduleExpression = "0 7 * * *";

const job = new CronJob(scheduleExpression, main); // change to scheduleExpressionMinute for testing

job.start();