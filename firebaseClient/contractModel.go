package firebaseClient

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// initiator is the party that scanned the qr
// recipient is the party whose qr was scanned
// recipient will have higher hierarchy

type Contract struct {
	InitiatorId string
	RecipientId string
}

func createContractId() (string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	id := base64.StdEncoding.EncodeToString(bytes)

	return id[:(len(id) - 1)], nil
}

func AddContract(contract *Contract, contractId *string) *string {
	if client == nil {
		CreateClient()
	}

	id, idErr := createContractId()

	if idErr != nil {
		log.Fatal("Error while trying to create clientId: ", idErr)
		log.Print("Error while trying to create clientId: ", idErr)
		return contractId
	}

	_, error := client.Collection("Contracts").Doc(id).Set(ctx, map[string]interface{}{
		"contractId":  id,
		"initiatorId": contract.InitiatorId,
		"recipientId": contract.RecipientId,
	})

	if error != nil {
		log.Fatal("Error while trying to access contract collection: ", error)
		log.Print("Error while trying to access contract collection: ", error)
		*contractId = ""

		return contractId
	}

	*contractId = id
	return contractId
}
