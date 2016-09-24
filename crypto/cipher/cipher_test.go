package cipher

import (
	"fmt"
	"testing"
)

func TestCBCEnDecrypter(t *testing.T) {

	for i := 0; i < 100; i++ {

		plaintext := fmt.Sprintf("%d.%s", i, "exampleplaintext")
		plaintext = plaintext[:16]

		encrypted, err := CBCEncrypter([]byte(plaintext))
		if err != nil {
			t.Errorf("CBCEncrypter error:%s\n", err.Error())
			continue
		}

		decrypted, err := CBCDecrypter(encrypted)
		if err != nil {
			t.Errorf("CBCDecrypter error:%s\n", err.Error())
			continue
		}

		if decrypted != plaintext {
			t.Errorf("CBC Encrypt/Decrypt wrong. plaintext:%s decrypted:%s\n",
				plaintext, decrypted)
		}

		// t.Logf("encrypt len:%d :%s decrypted:%s err:%v\n",
		//	len(encrypted), encrypted, decrypted, err)
	}
}
