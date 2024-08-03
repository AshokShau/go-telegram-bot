# Go Telegram Bot

This is a simple Telegram bot written in Go using the `gotgbot` library. The bot responds to the `/start` command with a welcome message.

## Features
Template for future projects.

## Status
Active

## Prerequisites

- Go 1.22.5 or later
- Telegram bot token from [BotFather](https://core.telegram.org/bots#botfather)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Abishnoi69/go-telegram-bot.git
    cd go-telegram-bot
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Add your Telegram bot token to the `.env` file:
    ```sh
    cp sample.env .env && vi .env
    ```

## Usage

1. Run the bot:
    ```sh
    go run main.go
    ```

2. Start a chat with your bot on Telegram and send the `/start` command.

## Project Structure

- `main.go`: Entry point. Initializes the bot and starts polling for updates.
- `Telegram/modules/start.go`: Handler for the `/start` command.
- `Telegram/modules/loadModules.go`: Loads command handlers into the dispatcher.

## Contributing

Submit issues or pull requests for bugs or improvements.

## License

Licensed under the MIT License.

## Thanks
- [gotgbot](https://github.com/PaulSonOfLars/gotgbot) for the library.
- [Abishnoi69](https://github.com/Abishnoi69) for the template.
- [Go](https://golang.org/) for the language.