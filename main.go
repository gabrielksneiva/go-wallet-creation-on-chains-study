package main

import (
	"fmt"
	"go-wallet-creation-on-chains-study/wallets"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// // Criar carteira
	// wallet1 := wallets.CreateWallet() // remetente
	// wallet2 := wallets.CreateWallet() // destinatário

	// fmt.Println("Remetente:", wallet1.Address.Hex())
	// fmt.Println("Destinatário:", wallet2.Address.Hex())

	// // Conectar à Sepolia
	// rpcURL := "https://sepolia.infura.io/v3/b509635471974105bb96970aa0c77aaa"
	// client := network.ConnectToEthereumNode(rpcURL)

	// // Consultar dados reais
	// balance := network.GetBalance(client, wallet1.Address)
	// nonce := network.GetNonce(client, wallet1.Address)
	// gasPrice := network.GetSuggestedGasPrice(client)

	// fmt.Println("💰 Saldo:", balance, "ETH")
	// fmt.Println("🔢 Nonce:", nonce)
	// fmt.Println("⛽ GasPrice:", gasPrice, "wei")

	// // Preparar destinatário
	// to := common.HexToAddress(wallet2.Address.Hex()) // coloque um destinatário válido

	// // Assinar com dados reais
	// tx := signing.SignSimpleTransaction(wallet1.PrivateKey, to, 0.01, nonce, gasPrice.Int64()/1e9, big.NewInt(11155111))

	// // Proximo passo: Enviar tx: client.SendTransaction(ctx, tx)
	// err := client.SendTransaction(ctx, tx)
	// if err != nil {
	// 	log.Fatalf("❌ Erro ao enviar a transação: %v", err)
	// }

	wallet := wallets.CreateWallet()
	fmt.Printf("Chave privada: %s\n", wallet.PrivateKey.D.Text(16))
}
