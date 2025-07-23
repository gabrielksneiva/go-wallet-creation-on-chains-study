package network

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Conecta ao nó Ethereum via RPC
func ConnectToEthereumNode(rpcURL string) *ethclient.Client {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Erro ao conectar à rede Ethereum: %v", err)
	}
	return client
}

// Retorna o saldo (ETH) de um endereço
func GetBalance(client *ethclient.Client, address common.Address) *big.Float {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	balanceWei, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Fatal(err)
	}
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balanceWei), big.NewFloat(1e18))
	return ethValue
}

// Retorna o nonce atual do endereço
func GetNonce(client *ethclient.Client, address common.Address) uint64 {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		log.Fatal(err)
	}
	return nonce
}

// Retorna o gas price sugerido pela rede
func GetSuggestedGasPrice(client *ethclient.Client) *big.Int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}
