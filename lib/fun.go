package lib

func CheckTx(txHash string) bool{
	if txHash != nil {
		return true
	}
	return false
}
