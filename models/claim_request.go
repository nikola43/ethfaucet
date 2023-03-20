package models

type ClaimRequest struct {
	WalletAddress string `json:"wallet_address"`
	IPAddress     string `json:"ip_address"`
}
