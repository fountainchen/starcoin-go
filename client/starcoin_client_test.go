package client

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/starcoinorg/starcoin-go/types"

	owcrypt "github.com/blocktree/go-owcrypt"
)

func TestHttpCall(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")
	var result interface{}

	result, err := client.GetNodeInfo()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetTransactionByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetTransactionInfoByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetTransactionEventByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetBlockByHash("0x9e635ae64903409378f5146ff89bfea52a61326ffcbf4191fa63cce642cfc2ea")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetBlockByNumber(2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetBlocksFromNumber(2, 10)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetResource("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetState("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetGasUnitPrice()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	call := ContractCall{
		"0x00000000000000000000000000000001::Token::market_cap",
		[]string{"0x00000000000000000000000000000001::STC::STC"},
		[]string{},
	}
	result, err = client.CallContract(call)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

}

func TestBalance(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")
	var result *ListResource

	result, err := client.GetResource("0x79f75dc7cb6812760e1afba01dc9380e")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result.GetBalanceOfStc())
}

func TestWsCall(t *testing.T) {
	client := NewStarcoinClient("ws://localhost:9870")

	//c, err := client.NewTxnSendRecvEventNotifications("0xb75994d55eae88219dc57e7e62a11bc0")
	c, err := client.NewPendingTransactionsNotifications()

	if err != nil {
		t.Error(err)
	}

	data := <-c

	fmt.Println(data)

	c1, err := client.NewTxnSendRecvEventNotifications("0xb75994d55eae88219dc57e7e62a11bc0")

	if err != nil {
		t.Error(err)
	}

	data1 := <-c1

	fmt.Println(data1)

}

func TestSubmitTransaction(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")
	privateKeyString := "7ddee640acc92417aee935daccfa34306b7c2b827a1308711d5b1d9711e1bdac"
	privateKeyBytes, _ := hex.DecodeString(privateKeyString)
	privateKey := types.Ed25519PrivateKey(privateKeyBytes)
	addressArray := ToAccountAddress("b75994d55eae88219dc57e7e62a11bc0")

	result, err := client.TransferStc(addressArray, privateKey, ToAccountAddress("ab4039861ca47ec349b64ddb862293bf"), serde.Uint128{
		High: 0,
		Low:  100000,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func TestNodeInfo(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")
	var result interface{}

	result, err := client.GetNodeInfo()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}

func TestSign(t *testing.T) {
	privateKeyString := "587737ebefb4961d377a3ab2f9ceb37b1fa96eb862dfaf954a4a1a99535dfec0"
	publicKeyString := "32ed52d319694aebc5b52e00836e2f7c7d2c7c7791270ede450d21dbc90cbfa1"

	privateKey, _ := hex.DecodeString(privateKeyString)
	publicKey, _ := hex.DecodeString(publicKeyString)

	publicKeyGen, _ := owcrypt.GenPubkey(privateKey, owcrypt.ECC_CURVE_ED25519_NORMAL)

	message := "Example `personal_sign` message"

	signBytes, _, _ := owcrypt.Signature(privateKey, nil, []byte(message), owcrypt.ECC_CURVE_ED25519_NORMAL)
	signString := hex.EncodeToString(signBytes)

	messageBytes, _ := hex.DecodeString("f7abb31497be2d952de2e1c64e2ce3edae7c4d9f5a522386a38af0c76457301eb75994d55eae88219dc57e7e62a11bc0070000000000000002000000000000000000000000000000010f5472616e73666572536372697074730f706565725f746f5f706565725f76320107000000000000000000000000000000010353544303535443000210b75994d55eae88219dc57e7e62a11bc010a0860100000000000000000000000000809698000000000001000000000000000d3078313a3a5354433a3a5354439e1d000000000000fe\n")

	signBytes, _, _ = owcrypt.Signature(privateKey, nil, messageBytes, owcrypt.ECC_CURVE_ED25519_NORMAL)
	signString = hex.EncodeToString(signBytes)

	fmt.Println(owcrypt.Verify(publicKeyGen, nil, messageBytes, signBytes, owcrypt.ECC_CURVE_ED25519_NORMAL))
	fmt.Println(publicKey)
	fmt.Println(publicKeyGen)
	fmt.Println(signString)
}

func TestDeployContract(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")
	privateKeyString := "7ddee640acc92417aee935daccfa34306b7c2b827a1308711d5b1d9711e1bdac"
	privateKeyBytes, _ := hex.DecodeString(privateKeyString)
	privateKey := types.Ed25519PrivateKey(privateKeyBytes)

	//code,_ := ioutil.ReadFile("/Users/fanngyuan/Documents/workspace/starcoin_java/src/test/resources/contract/MyCounter.mv")
	//fmt.Println(code)
	code := []byte{161, 28, 235, 11, 2, 0, 0, 0, 9, 1, 0, 4, 2, 4, 4, 3, 8, 25, 5, 33, 12, 7, 45, 78, 8, 123, 32, 10, 155, 1, 5, 12, 160, 1, 81, 13, 241, 1, 2, 0, 0, 1, 1, 0, 2, 12, 0, 0, 3, 0, 1, 0, 0, 4, 2, 1, 0, 0, 5, 0, 1, 0, 0, 6, 2, 1, 0, 1, 8, 0, 4, 0, 1, 6, 12, 0, 1, 12, 1, 7, 8, 0, 1, 5, 9, 77, 121, 67, 111, 117, 110, 116, 101, 114, 6, 83, 105, 103, 110, 101, 114, 7, 67, 111, 117, 110, 116, 101, 114, 4, 105, 110, 99, 114, 12, 105, 110, 99, 114, 95, 99, 111, 117, 110, 116, 101, 114, 4, 105, 110, 105, 116, 12, 105, 110, 105, 116, 95, 99, 111, 117, 110, 116, 101, 114, 5, 118, 97, 108, 117, 101, 10, 97, 100, 100, 114, 101, 115, 115, 95, 111, 102, 248, 175, 3, 221, 8, 222, 73, 216, 30, 78, 253, 158, 36, 192, 57, 204, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1, 7, 3, 0, 1, 0, 1, 0, 3, 13, 11, 0, 17, 4, 42, 0, 12, 1, 10, 1, 16, 0, 20, 6, 1, 0, 0, 0, 0, 0, 0, 0, 22, 11, 1, 15, 0, 21, 2, 1, 2, 0, 1, 0, 1, 3, 14, 0, 17, 0, 2, 2, 1, 0, 0, 1, 5, 11, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 18, 0, 45, 0, 2, 3, 2, 0, 0, 1, 3, 14, 0, 17, 2, 2, 0, 0, 0}

	moduleId := types.ModuleId{
		Address: ToAccountAddress("b75994d55eae88219dc57e7e62a11bc0"),
		Name:    "xxxx",
	}

	scriptFunction := types.ScriptFunction{
		Module:   moduleId,
		Function: "init",
		TyArgs:   []types.TypeTag{},
		Args:     [][]byte{},
	}
	client.DeployContract(ToAccountAddress("b75994d55eae88219dc57e7e62a11bc0"), privateKey, scriptFunction, code)
}
