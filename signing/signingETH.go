package signing

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SignSimpleTransaction(
	privateKey *ecdsa.PrivateKey,
	toAddress common.Address,
	valueInEth float64,
	nonce uint64,
	gasPriceGwei int64,
	chainID *big.Int,
) *types.Transaction {

	value := big.NewInt(int64(valueInEth * 1e18)) // ETH -> wei
	gasLimit := uint64(21000)                     // envio simples
	gasPrice := big.NewInt(gasPriceGwei * 1e9)    // Gwei -> wei
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transação assinada! Hash:", signedTx.Hash().Hex())
	return signedTx
}
