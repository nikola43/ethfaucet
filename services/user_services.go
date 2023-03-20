package services

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	database "github.com/nikola43/ethfaucet/database"
	"github.com/nikola43/ethfaucet/models"
	"github.com/nikola43/ethfaucet/utils"
)

func Claim(claimRequest *models.ClaimRequest) (*models.ClaimResponse, error) {

	pk := "2408ee769f8ece93ad70e1da4cbfea8368e0946bcd07fcc306a8cbb4daf8f7b3"

	user := new(models.User)

	gormDBResult := database.GormDB.Where("wallet_address = ?", claimRequest.WalletAddress).Find(&user)
	if gormDBResult.Error != nil {
		fmt.Println("gormDBResult result")
		return nil, gormDBResult.Error
	}

	if user.ID > 0 {
		lastClaimDate := user.ClaimDate
		currentTime := time.Now()
		diff := currentTime.Sub(lastClaimDate)
		fmt.Printf("Seconds: %f\n", diff.Hours())

		isSameWallet := user.WalletAddress == claimRequest.WalletAddress
		mustWait := diff < time.Hour*24

		//mustWait := diff < time.Second*10
		if isSameWallet {
			if mustWait {
				return nil, errors.New("Already claimed, you must wait ")
			}
		}

		user.WalletAddress = claimRequest.WalletAddress
		user.IPAddress = claimRequest.IPAddress
		user.ClaimDate = currentTime

		result := database.GormDB.Model(&user).Update("claim_date", currentTime)
		if result.Error != nil {
			fmt.Println("result.Error result")
			return nil, result.Error
		}

		err, txHash := utils.TransferEth(pk, user.WalletAddress, big.NewInt(10000000000000000)) // 0.1 eth
		claimResponse := new(models.ClaimResponse)
		claimResponse.TxHash = txHash

		if err != nil {
			return nil, err
		}
		return claimResponse, nil

	} else {

		user.WalletAddress = claimRequest.WalletAddress
		user.IPAddress = claimRequest.IPAddress
		user.ClaimDate = time.Now()

		result := database.GormDB.Create(&user)
		if result.Error != nil {
			fmt.Println("result.Error")
			return nil, result.Error
		}

		claimResponse := new(models.ClaimResponse)
		err, txHash := utils.TransferEth(pk, user.WalletAddress, big.NewInt(10000000000000000)) // 0.1 eth
		claimResponse.TxHash = txHash
		if err != nil {
			return nil, err
		}
		return claimResponse, nil
	}
}
