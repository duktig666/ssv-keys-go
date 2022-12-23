// description:
// @author renshiwei
// Date: 2022/12/15 17:19

package shares

import (
	"encoding/json"
	"github.com/duktig666/ssv-keys-go/keystorev4"
	"github.com/duktig666/ssv-keys-go/ssv"
	"github.com/pkg/errors"
	"io/ioutil"
)

func keystoreShare() (string, error) {
	keystoreFile, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		return "", errors.Wrap(err, "read keystore fail.")
	}

	// get keystore
	keystore := string(keystoreFile)

	// parse keystore to get public key
	var keystoreStruct *keystorev4.Keystorev4
	err = json.Unmarshal(keystoreFile, &keystoreStruct)
	if err != nil {
		return "", errors.Wrap(err, "keystore parse fail.")
	}
	pubkey := keystoreStruct.Pubkey

	// crypt the keystore
	skBytes, err := keystorev4.DecryptFromJson(keystore, password)
	if err != nil {
		return "", errors.Wrap(err, "keystore decrypt fail")
	}

	// get operator list
	if len(operatorIdList) != len(operatorPubkeyList) {
		return "", errors.New("operator and operator-ids are inconsistent.")
	}

	count := len(operatorIdList)
	operators := make([]*ssv.Operator, 0, count)

	for i := 0; i < count; i++ {
		operators = append(operators, &ssv.Operator{
			Id:        int(operatorIdList[i]),
			PublicKey: operatorPubkeyList[i],
		})
	}

	shareRes, err := ssv.KeystoreShareV2ForJson(pubkey, "v2", ssvAmount, skBytes, operators)
	if err != nil {
		return "", err
	}

	return shareRes, nil
}

func output(data []byte) error {
	err := ioutil.WriteFile(outputPath, data, 0777)
	if err != nil {
		return errors.Wrap(err, "write file error.")
	}
	return nil
}
