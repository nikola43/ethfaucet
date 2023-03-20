package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TransferEth(fromPk string, to string, value *big.Int) (error, string) {
	client, err := ethclient.Dial("https://rpc.ankr.com/eth_goerli")
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	privateKey, err := crypto.HexToECDSA(fromPk)
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("publicKeyECDSA")
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey"), ""
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	toAddress := common.HexToAddress(to)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err.Error())
		return err, ""
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return nil, signedTx.Hash().Hex()
}
