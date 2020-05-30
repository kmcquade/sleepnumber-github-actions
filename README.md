# sleepnumber-github-actions

Use GitHub actions and my Sleep Number bed to wake me up in the morning

## Instructions

Fork the repo.

Go to GitHub settings and set the following as secrets:

* `SLEEPIQ_USER`: The email for your Sleep Number login
* `SLEEPIQ_PASS`: The password for your Sleep Number login

... and I'm still working out the rest of the details. Not sure if I will set the cron job parameters in the GitHub actions workflow or if I can set it up separately as environment variables.