package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/spf13/viper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	sapphire "github.com/oasisprotocol/sapphire-paratime/clients/go"

	"fetcher/creditData"
)

func setupViper() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("fetcher.conf")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Viper read error: %s\n", err)
	}
}

func main() {
	setupViper()

	expiryOffset := int64(viper.Get("expiryOffset").(int))

	// Prepare private key
	key, err := crypto.HexToECDSA(viper.Get("privateKey").(string))
	if err != nil {
		log.Fatalf("Private key error: %s\n", err)
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("Invalid private key!")
	}

	ctx := context.Background()
	conn, err := ethclient.DialContext(ctx, viper.Get("gateway").(string))
	if err != nil {
		log.Fatalf("Dial error: %s\n", err)
	}

	sender := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("Sender address: %s\n", sender)

	backend, err := sapphire.WrapClient(*conn, func(digest [32]byte) ([]byte, error) {
		// Pass in a custom signing function to interact with the signer
		return crypto.Sign(digest[:], key)
	})

	if err != nil {
		log.Fatalf("Failed to create Sapphire backend signer: %s\n", err)
	}

	defer conn.Close()
	log.Println("Created Sapphire backend signer client")

	contractAddress := common.HexToAddress(viper.Get("contractAddress").(string))
	cd, err := creditData.NewCreditData(contractAddress, backend)
	if err != nil {
		log.Fatalf("Failed to create contract instance: %s\n", err)
	}

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	scoreReqs := make(chan *creditData.CreditDataRequestScore)
	scoreSub, err := cd.WatchRequestScore(watchOpts, scoreReqs)
	if err != nil {
		log.Fatalf("Failed to create RequestScore watcher: %s\n", err)
	}
	defer scoreSub.Unsubscribe()

	reportReqs := make(chan *creditData.CreditDataRequestReport)
	reportSub, err := cd.WatchRequestReport(watchOpts, reportReqs)
	if err != nil {
		log.Fatalf("Failed to create RequestReport watcher: %s\n", err)
	}
	defer reportSub.Unsubscribe()

	for {
		select {
		case err := <-scoreSub.Err():
			log.Fatalf("RequestScore watch error: %s\n", err)

		case err := <-reportSub.Err():
			log.Fatalf("RequestReport watch error: %s\n", err)

		case event := <-scoreReqs:
			log.Printf("Received RequestScore event for %s\n", event.Whom)
			score := uint16(rand.Int31n(800))
			log.Printf("Score: %d\n", score)
			txOpts := backend.Transactor(sender)
			trans, err := cd.StoreCreditScore(txOpts, event.Whom, score, big.NewInt(time.Now().Unix()+expiryOffset))
			if err != nil {
				log.Printf("Error writing score for %s: %s\n", event.Whom, err)
			} else {
				log.Printf("Done writing; hash: %+v\n", trans.Hash())
			}

		case event := <-reportReqs:
			log.Printf("Received RequestReport event for %s\n", event.Whom)
			report := fmt.Sprintf("Report #%d", rand.Int31n(100000))
			log.Printf("Received report: %s\n", report)
			txOpts := backend.Transactor(sender)
			trans, err := cd.StoreCreditReport(txOpts, event.Whom, report, big.NewInt(time.Now().Unix()+expiryOffset))
			if err != nil {
				log.Printf("Error writing report for %s: %s\n", event.Whom, err)
			} else {
				log.Printf("Done writing; hash: %+v\n", trans.Hash())
			}
		}
	}
}
