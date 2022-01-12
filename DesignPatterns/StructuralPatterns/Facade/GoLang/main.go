package main

import (
	"fmt"
	"log"
)

//ABSTRACTION
type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}
func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletDebitNotifaction()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitbalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotifaction()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

//ACCOUNT SECTION
type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}
func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account name is Incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

//SECURITY CODE SECTION
type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}
func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("Security Code Verified")
	return nil
}

//WALLET SECTION
type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}
func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet Ballance added successfully")
	return
}
func (w *wallet) debitbalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

//LEDGER
type ledger struct {
}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

//NOTIFICATIOn
type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}
func (n *notification) sendWalletDebitNotifaction() {
	fmt.Println("Sending wallet debit notificaiton")
}

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 100000)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 10000)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
