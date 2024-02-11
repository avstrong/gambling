# Gambling Simulation Game - Alpha Version

Welcome to the alpha version of our simple Gambling simulation game. This project allows users to engage in a game where they choose between `true` or `false` to win. Players can create games, set the game's name, number of players, entry fee, and the currency for the entry fee. Winners take all the money or share it amongst those who chose correctly. This document outlines how to get started, use the API, and test the game's functionality.

## Features

- **User Registration**: Register users to participate in the game.
- **Wallet Operations**: Users can deposit and withdraw money in available currencies (USD, EUR) and check their balances.
- **Game Creation**: Create games with specific attributes like name, maximum players, entry fee, and currency.
- **Player Registration**: Register players for games.
- **Game Play**: Start the game and determine the winner.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system to run or build the application.

### Running the Application

To run the application directly:

```bash
go run main.go
```

To build and run the application:
```bash
make build-app
./bin/app start
```

Use make help to view a list of available commands.

### API Endpoints

The application provides the following RESTful endpoints:

* **Create Game: POST /api/v1/game**
* **Register Players: POST /api/v1/game/register**
* **Play Game: POST /api/v1/game/play**
* **Register User: POST /api/v1/user**
* **Deposit Funds: POST /api/v1/wallet/deposit**
* **Withdraw Funds: POST /api/v1/wallet/withdraw**
* **Check Balance: GET /api/v1/wallet/balance**

### Testing the Game

Follow these steps to test the game functionality:

#### Register a User:

**POST** `http://localhost:9092/api/v1/user`

**Body:**
```json
{
  "email": "user@gmail.com"
}
```

#### Deposit Funds:

**POST** `http://localhost:9092/api/v1/wallet/deposit`

**Body:**
```json
{
  "userID": "user@gmail.com",
  "amount": 12.00,
  "currency": "USD"
}
```

#### Check Balance:

**GET** `http://localhost:9092/api/v1/wallet/balance`

**Body:**
```json
{
  "userID": "user@gmail.com",
  "currency": "USD"
}
```

#### Withdraw Funds:

**POST** `http://localhost:9092/api/v1/wallet/withdraw`

**Body:**
```json
{
  "userID": "user@gmail.com",
  "amount": 10.00,
  "currency": "USD"
}
```

#### Create a Game:

**POST** `http://localhost:9092/api/v1/game`

**Body:**
```json
{
  "name": "game 1",
  "maxPlayers": 2,
  "entryFee": 2,
  "entryCurrency": "USD"
}
```

#### Register Players for the Game:

**POST** `http://localhost:9092/api/v1/game/register`

**Body:**
```json
{
  "gameID": "{gameID}",
  "userID": "user@gmail.com",
  "playerChoice": false
}
```

#### Start the Game:

**POST** `http://localhost:9092/api/v1/game/play`

**Body:**
```json
{
  "id": "{gameID}"
}
```

After the game concludes, you can check the balances of the winning players, along with their total wins and the date/time of their last win, paving the way for future expansions.

## Future Work
* Implement more currencies and payment methods.
* Introduce more comprehensive player statistics.
* Expand the game rules and modes.
