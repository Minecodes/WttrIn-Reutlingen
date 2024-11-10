import { BskyAgent } from "@atproto/api";
import * as dotenv from "dotenv";
import { CronJob } from "cron";
import * as process from "process";
import axios from "axios";

dotenv.config();

// Create a Bluesky Agent
const agent = new BskyAgent({
  service: "https://bsky.social",
});

async function main() {
  await agent.login({
    identifier: process.env.BLUESKY_USERNAME!,
    password: process.env.BLUESKY_PASSWORD!,
  });
  await agent.post({
    text: "🙂",
  });
  console.log("Weather report got posted");
}

main();

/**
 * Cron formating: Minute Hour Day Month DayOfWeek
 * 
 * 0 7 * * * - Every day at 7:00 AM
 */
const scheduleExpressionMinute = "* * * * *";
const scheduleExpression = "0 7 * * *";

const job = new CronJob(scheduleExpression, main); // change to scheduleExpressionMinute for testing

job.start();
