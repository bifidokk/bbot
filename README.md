# Simple Telegram Bot

![](https://github.com/bifidokk/bbot/workflows/build/badge.svg)

Just a simple telegram bot written in Golang. 

Environment variables:
* BBOT_TOKEN="telegrambottoken"
* BBOT_BASE_URL=":9000"
* BBOT_WEBHOOK_URL="https://domain.com"
* BBOT_WEBHOOK_PATH="/bbot/"
* BBOT_DB_DSN="host=postgres user=postgres password=postgres dbname=postgres port=5432"

For local development:
ngrok http 9000