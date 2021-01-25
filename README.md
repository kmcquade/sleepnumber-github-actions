# sleepnumber-github-actions

Use GitHub actions and my Sleep Number bed to wake me up in the morning. Tilts the bed to "Watch TV" position and helps me wake up without any alarms. Previously, I would set 2 hours of alarms to get out of bed. Now, I save those two hours of sleep and wake up very gently in the morning with this trick.

![Test wakeup](https://github.com/kmcquade/sleepnumber-github-actions/workflows/Test%20wakeup/badge.svg)

## Instructions

Fork the repo.

Go to GitHub settings and set the following as secrets:

* `SLEEPIQ_USER`: The email for your Sleep Number login
* `SLEEPIQ_PASS`: The password for your Sleep Number login

The following GitHub action files do this magic:
* `.github/workflows/test.yml`: This will execute on every push. This is nice so you can just push to master and do testing, without having to do some crazy curl commands to trigger the action manually
* `.github/workflows/weekdays730.yml`: This wakes me up at 730am and 8am Eastern by tilting the bed. Note that the times are set via Cron syntax. GitHub actions uses UTC instead of eastern so I had to adjust the Cron syntax accordingly.

This is SO much better than my previous iteration, which required the use of Jenkins on Raspberry Pi.

## Credits

Shoutout to Dan Penn for his Golang package that powers this thing: https://github.com/danpenn/SleepIQ

