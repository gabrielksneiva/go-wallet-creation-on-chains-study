package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 1. Gera uma chave privada (secp256k1)
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Converte para formato hexadecimal (útil para armazenar)
	privateKeyHex := crypto.FromECDSA(privateKey)
	fmt.Println("🔐 Chave privada (hex):", common.Bytes2Hex(privateKeyHex))

	// 3. Extrai a chave pública
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("erro ao converter para ECDSA")
	}

	// 4. Converte chave pública para formato hexadecimal
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("🔓 Chave pública (hex):", common.Bytes2Hex(publicKeyBytes))

	// 5. Gera o endereço Ethereum (últimos 20 bytes do Keccak256 da chave pública)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("🏦 Endereço Ethereum:", address)
}
