package main

import (
	"fmt"
)

// For the sake of ease, we've hardcoded these values
const (
	OddFor1 = 1.55
	OddFor0 = 1.85
)

var (
	PoolBalance = 10000000
	BetID       = 1
)

type LProvider struct {
	LpAccountId   int
	Amount        float
	Balance       float
	CurrencyDenom string
}

type Bettor struct {
	BettorAccountId int
	BetAmount       float
	BetValue        bool
	Balance         float
	Payout          float
	CurrencyDenom   string
}

func addBettor(bettorAccountId int) bool {
	return true
}

func placeBet(bettorAccountId int, betID int, betValue bool, betAmount float, currencyDenom string) bool {
	return true
}

func provideLiquidity(lpAccountId int, amount float, denom string) bool {
	return true
}

func handleResult(betID int, betValue bool) {
	return true
}

func getLpBalance() float {
	return 0
}

func getBalance(bettorAccountId int) (isAccountExist bool, balance int) {
	return 0
}

func main() {
	var allLPs []int
	var allBettors []int
	var allBets []int
	fmt.Println("Smart contract for Betting in Golang")
}

