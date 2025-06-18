Here is the standard notation for the cron to schedule the jobs...

┌──────────── second (0 - 59)
│ ┌────────── minute (0 - 59)
│ │ ┌──────── hour (0 - 23)
│ │ │ ┌────── day of month (1 - 31)
│ │ │ │ ┌──── month (1 - 12)
│ │ │ │ │ ┌── day of week (0 - 6) (Sunday=0)
│ │ │ │ │ │
* * * * * *

*/10 * * * * * will schedule the job every 10 seconds
