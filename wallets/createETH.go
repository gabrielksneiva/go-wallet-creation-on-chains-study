package wallets

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func CreateWallet() *Wallet {
	// -------------------------------------
	// 1. Gerar uma chave privada (256 bits)
	// -------------------------------------

	// Uma chave privada é apenas um número inteiro de 256 bits (32 bytes)
	// Cada byte é uma sequência de 8 bits (valores de 0 a 255)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// NOTA: Ethereum usa a curva SECP256k1, mas aqui usamos P256 apenas como exemplo
	// Em produção, você precisaria usar SECP256k1 (exige biblioteca externa ou implementação manual)

	// Convertendo a chave privada para hexadecimal para facilitar visualização
	privateKeyBytes := privateKey.D.Bytes() // D é o valor inteiro da chave privada
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	fmt.Println("🔐 Private Key:", privateKeyHex)

	// -------------------------------------
	// 2. Gerar a chave pública (X, Y)
	// -------------------------------------

	// A chave pública no ECDSA é um ponto (X, Y) na curva elíptica
	// A curva SECP256k1 é uma curva elíptica definida por uma equação matemática
	publicKeyX := privateKey.PublicKey.X.Bytes()
	publicKeyY := privateKey.PublicKey.Y.Bytes()

	// Concatenamos X e Y para formar a chave pública completa
	// A chave pública não é criptografada, mas é derivada de forma unidirecional da chave privada
	publicKeyBytes := append(publicKeyX, publicKeyY...)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)
	fmt.Println("🔓 Public Key (uncompressed):", publicKeyHex)

	// -------------------------------------
	// 3. Gerar o endereço Ethereum (Keccak256)
	// -------------------------------------

	// O endereço Ethereum é os últimos 20 bytes do hash Keccak256 da chave pública
	// Keccak256 é uma função de hash semelhante ao SHA3, mas com padding diferente
	// Ela pega uma entrada (bytes) e gera um "resumo" único de 32 bytes (256 bits)
	hash := sha3.NewLegacyKeccak256() // Usamos a versão legacy do Keccak256 usada pelo Ethereum
	hash.Write(publicKeyBytes)        // Geramos o hash da chave pública
	hashed := hash.Sum(nil)           // nil significa que não temos bytes extras para adicionar

	// O endereço Ethereum é os últimos 20 bytes do hash
	address := hashed[12:] // Pegamos os 20 últimos bytes (160 bits)
	addressHex := hex.EncodeToString(address)
	fmt.Println("🏠 Ethereum Address: 0x" + addressHex)

	return &Wallet{
		PrivateKey: privateKey,
		Address:    common.HexToAddress("0x" + addressHex), // Convert
	}
}
