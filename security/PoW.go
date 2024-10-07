package sec

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// Function to calculate PoW
func proofOfWork(publicKey string, difficulty int) (string, int) {
	var nonce int
	var hash string

	// Loop until we find a valid hash
	for {
		// Combine public key with nonce
		data := publicKey + strconv.Itoa(nonce)
		hashBytes := sha256.Sum256([]byte(data))
		hash = hex.EncodeToString(hashBytes[:])

		// Check if the hash has the required number of leading zeros
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			break
		}
		nonce++
	}

	return hash, nonce
}

// Verifying the solution
func verifyPoW(publicKey string, nonce int, difficulty int) bool {
	data := publicKey + strconv.Itoa(nonce)
	hashBytes := sha256.Sum256([]byte(data))
	hash := hex.EncodeToString(hashBytes[:])

	// Check for leading zeros
	return strings.HasPrefix(hash, strings.Repeat("0", difficulty))
}

func main() {
	// Example: public key and difficulty
	publicKey := "unicom-node-public-key"
	difficulty := 4 // Number of leading zeros

	// Perform proof of work
	hash, nonce := proofOfWork(publicKey, difficulty)
	fmt.Printf("Solved PoW: Hash=%s, Nonce=%d\n", hash, nonce)

	// Verify the PoW solution
	if verifyPoW(publicKey, nonce, difficulty) {
		fmt.Println("Proof of Work verified!")
	} else {
		fmt.Println("Verification failed.")
	}
}
