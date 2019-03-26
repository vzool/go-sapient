package sapient

import "testing"

// Generated by original Sapient written in PHP
// Check https://github.com/paragonie/sapient
var testSealingSecretKeys = []string{
	"8adhixPUUtrISMaQVjkDnIZRlCB_PWb9cjIxWtuK9D4=",
	"lgeXCWVbEDeUD2yvv0w3gOWNfUA1cPUrm0f_MxAzWZc=",
	"bFzItK-qPpgWSKMGb6VRkuZW_dI_3p2zjhEEj97kDrk=",
	"9RG_rmltUAPaG6S-IrewMLPeAYRH4Qz0Wn8IVb_GfyU=",
	"i3kW7JmnAX2OcuGY5Ns51tCAZ-uMI7TmRDpRmhjKjTo=",
	"LWRumTjq_wYhP-ToWEkIRS2jLoNpl5AsEPNqtURf5rE=",
	"04o_ZP34Zqy1Znq6FlJhUNsjS7KlU9UWCrn-fHIrnCM=",
	"r-wUUUfEe-R-1mOY3KBqtwxs_RBh3go7TTcosUKUuGw=",
	"jeoD7DoLWPHEdVGw8rLYhCICUldChEkaoIXYDX3W77w=",
	"2c1w6AVg5nhotOL4t26kc_sgCGppfmjhW2MBMuh-EFs=",
}

func TestSealingSecretKey(t *testing.T) {

	// test for 10 times

	for i := 0; i < 10; i++ {

		key := NewSealingSecretKey(nil)

		verifySealingSecretKey(key, t)
	}
}

func TestSealingSecretKey_LoadKeys(t *testing.T) {

	for _, encoded := range testSealingSecretKeys {

		// Test: Load key at instance creation
		key := NewSealingSecretKey(encoded)

		// Test: Load key after instance creation
		key.Load(encoded)

		// Test: External checking and Byte by Byte match
		verifySealingSecretKey(key, t)
	}
}

func verifySealingSecretKey(key *SealingSecretKey, t *testing.T) {

	bytesSize := SealingSecretKeyBytesSize
	encodedSize := SealingSecretKeyEncodedSize

	if size := len(key.Bytes()); size < bytesSize || size > bytesSize {
		t.Errorf("Key Size(%v) not match constant value(%v)", size, bytesSize)
	}

	encoded := key.String()

	if size := len(encoded); size < encodedSize || size > encodedSize {
		t.Errorf("Key Size(%v) not match constant value(%v)", size, encodedSize)
	}

	decoded, err := Base64UrlDecode(encoded)

	if err != nil {

		t.Error("Failed to decode generated key", err.Error())
	}

	// Match bytes by bytes after encoding

	for i, b := range decoded {

		if key.Bytes()[i] != b {

			t.Errorf("Byte(%v) missmatch with original Byte(%v)", b, key.Bytes()[i])
		}
	}
}
