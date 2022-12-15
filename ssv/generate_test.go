// description:
// @author renshiwei
// Date: 2022/11/3 18:03

package ssv

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLen(t *testing.T) {
	str := "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBeGlXQm5ya3g5SVNXSFhDSWhJZG0Kek1sWnNOMFd2QldseHZIWUdSU0xnNlI0Q0J1dlFFSitKemUrN1dsVnRkaTNTM1ZIK2NWNXBTdzVTYnFvd2RZZgo0RnEyTWN6dWxSM1d3dU5pWjByUjFJc2Jwb0h6TUNYcUJwTzJOeHlMWVhIeHp3RG1ML3VSQlp5dW5zMnF4MUZuCllReUlmREIxSkI1OE1WVndWM1R0UnVCWXFJNTdIakpaUmw5c1gwdC93UmlhNG5CUUs1WmgwbGd1dzV6ZEYyeEsKQ3h2K05KdVZPODg5RHlMZDVjOGg4MmVSeFJHclJFR0lFaVJaMFhlUER6ME9DdkZaSFNSUEZ3NXNQODl3THBFbwplTlNjYzN2d05NcW8vV25TaEZNdDJkNzNNd1E4Ky9IMFdaQlRiZnhrSzY3YXJXVC9mc2RKOGcwN3BYYWZWV0h5ClB3SURBUUFCCi0tLS0tRU5EIFJTQSBQVUJMSUMgS0VZLS0tLS0K"
	// 612
	fmt.Println(len(str))
}

var (
	// SkPem is a operator private key
	SkPem = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAowE7OEbwyLkvrZ0TU4jjooyIFxNvgrY8Fj+WslyZTlyj8UDf\nFrYh5Un2u4YMdAe+cPf1XK+A/P9XX7OB4nf1OoGVB6wrC/jhLbvOH650ryUYopeY\nhlSXxGnD4vcvTvcqLLB+ue2/iySxQLpZR/6VsT3fFrEonzFTqnFCwCF28iPnJVBj\nX6T/HcTJ55IDkbtotarU6cwwNOHnHkzWrv7ityPkR4Ge11hmVG9QjROt56ehXfFs\nFo5MqSvqpYplXkI/zUNm8j/lqEdU0RXUr41L2hyKY/pVjsgmeTsN7/ZqACkHye9F\nbkV9V/VbTh7hWVLTqGSh7BY/D7gwOwfuKiq2TwIDAQABAoIBADjO3Qyn7JKHt44S\nCAI82thzkZo5M8uiJx652pMeom8k6h3SNe18XCPEuzBvbzeg20YTpHdA0vtZIeJA\ndSuwEs7pCj86SWZKvm9p3FQ+QHwpuYQwwP9Py/Svx4z6CIrEqPYaLJAvw2mCyCN+\nzk7A8vpqTa1i4H1ae4YTIuhCwWlxe1ttD6rVUYfC2rVaFJ+b8JlzFRq4bnAR8yme\nrE4iAlfgTOj9zL814qRlYQeeZhMvA8T0qWUohbr1imo5XzIJZayLocvqhZEbk0dj\nq9qKWdIpAATRjWvb+7PkjmlwNjLOhJ1phtCkc/S4j2cvo9gcS7WafxaqCl/ix4Yt\n5KvPJ8ECgYEA0Em4nMMEFXbuSM/l5UCzv3kT6H/TYO7FVh071G7QAFoloxJBZDFV\n7fHsc+uCimlG2Xt3CrGo9tsOnF/ZgDKNmtDvvjxmlPnAb5g4uhXgYNMsKQShpeRW\n/ay8CmWbsRqXZaLoI5br2kCTLwsVz2hpabAzBOr2YV3vMRB5i7COYSMCgYEAyFgL\n3DkKwsTTyVyplenoAZaS/o0mKxZnffRnHNP5QgRfT4pQkuogk+MYAeBuGsc4cTi7\nrTtytUMBABXEKGIJkAbNoASHQMUcO1vvcwhBW7Ay+oxuc0JSlnaXjowS0C0o/4qr\nQ/rpUneir+Vu/N8+6edETRkNj+5unmePEe9NBuUCgYEAgtUr31woHot8FcRxNdW0\nkpstRCe20PZqgjMOt9t7UB1P8uSuqo7K2RHTYuUWNHb4h/ejyNXbumPTA6q5Zmta\nw1pmnWo3TXCrze0iBNFlBazf2kwMdbW+Zs2vuCAm8dIwMylnA6PzNj7FtRETfBqr\nzDVfdsFYTcTBUGJ21qXqaV0CgYEAmuMPEEv9WMTo43VDGsaCeq/Zpvii+I7SphsM\nmMn8m6Bbu1e4oUxmsU7RoanMFeHNbiMpXW1namGJ5XHufDYHJJVN5Zd6pYV+JRoX\njjxkoyke0Hs/bNZqmS7ITwlWBiHT33Rqohzaw8oAObLMUq2ZqyYDtQNYa90vIkH3\n5yq1x00CgYEAs4ztQhGRbeUlqnW6Z6yfRJ6XXYqdMPhxuBxvNn/dxJ10T4W2DUuC\njSdpGXrY+ECYyXUwlXBqbaKx1K5AQD7nmu9J3l0oMkX6tSBj1OE5MabATrsW6wvT\nhkTPJZMyPUYhoBkivPUKyQXswrQV/nUQAsAcLeJShTW4gSs0M6weQAc=\n-----END RSA PRIVATE KEY-----\n"
	// EncryptedKeyBase64 SkPem in base64 format
	EncryptedKeyBase64 = "NW/6N5Ubo5T+oiT9My2wXFH5TWT7iQnN8YKUlcoFeg00OzL1S4yKrIPemdr7SM3EbPeHlBtOAM3z+06EmaNlwVdBiexSRJmgnknqwt/Ught4pKZK/WdJAEhMRwjZ3nx1Qi1TYcw7oZBaOdeTdm65QEAnsqOHk1htnUTXqsqYxVF750u8JWq3Mzr3oCN65ydSJRQoSa+lo3DikIDrXSYe1LRY5epMRrOq3cujuykuAVZQWp1vzv4w4V6mffmxaDbPpln/w28FKCxYkxG/WhwGuXR1GK6IWr3xpXPKcG+lzfvlmh4UiK1Lad/YD460oMXOKZT8apn4HL4tl9HOb6RyWQ=="
)

func TestGenerateKeysBase64(t *testing.T) {
	pk, sk, err := GenerateKeysBase64()
	require.NoError(t, err)
	fmt.Println("pk:", pk)
	fmt.Println("sk:", sk)
}

func TestGenerateKeysOutput(t *testing.T) {
	pkB, skB, err := GenerateKeys()
	require.NoError(t, err)
	fmt.Println("pk:", string(pkB))
	fmt.Println("sk:", string(skB))

	fmt.Println("pk:", base64.StdEncoding.EncodeToString(pkB))
	fmt.Println("sk:", base64.StdEncoding.EncodeToString(skB))
}

func TestGenerateKeys(t *testing.T) {
	_, skByte, err := GenerateKeys()
	require.NoError(t, err)
	sk, err := ConvertPemToPrivateKey(string(skByte))
	require.NoError(t, err)
	require.Equal(t, 2048, sk.N.BitLen())
	require.NoError(t, sk.Validate())
}

func TestDecodeKey(t *testing.T) {
	sk, err := ConvertPemToPrivateKey(SkPem)
	require.NoError(t, err)
	key, err := DecodeKey(sk, EncryptedKeyBase64)
	require.NoError(t, err)
	require.Equal(t, "626d6a13ae5b1458c310700941764f3841f279f9c8de5f4ba94abd01dc082517", key)
}

func TestExtractPublicKey(t *testing.T) {
	_, skByte, err := GenerateKeys()
	require.NoError(t, err)
	sk, err := ConvertPemToPrivateKey(string(skByte))
	require.NoError(t, err)
	pk, err := ExtractPublicKey(sk)
	require.NoError(t, err)
	require.NotNil(t, pk)
}

func TestPrivateKeyToByte(t *testing.T) {
	_, skByte, err := GenerateKeys()
	require.NoError(t, err)
	sk, err := ConvertPemToPrivateKey(string(skByte))
	require.NoError(t, err)
	b := PrivateKeyToByte(sk)
	require.NotNil(t, b)
	require.Greater(t, len(b), 1024)
}
