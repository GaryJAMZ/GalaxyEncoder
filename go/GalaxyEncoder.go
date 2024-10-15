package GalaxyEncoder

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func GalaxyHexEncoder(textToEncode string, lenKey int) string {
	return fmt.Sprintf("%x", GalaxyEncoder([]byte(textToEncode), lenKey))
}
func GalaxyHexDecoder(encodedText string, lenKey int) string {
	var arrayBytes []byte
	for i := 0; i < len(encodedText); i += 2 {
		byteValue, _ := strconv.ParseUint(encodedText[i:i+2], 16, 8)
		arrayBytes = append(arrayBytes, byte(byteValue))
	}
	return string(GalaxyDecoder(arrayBytes, lenKey))
}
func GalaxyEncoder(bytesToEncode []byte, lenKey int) []byte {
	if lenKey < 1 {
		lenKey = 1
	} else if lenKey > 10 {
		lenKey = 10
	}

	var enodedBytes []byte
	clave := make([]byte, lenKey)
	rand.Read(clave)
	fmt.Println(clave)
	for _, b := range bytesToEncode {
		enodedBytes = append(enodedBytes, bynaryChangerBytes(clave, b, false))
	}
	finalBytes, _ := buildBytes(lenKey, enodedBytes, clave, false)
	return finalBytes
}
func GalaxyDecoder(encodedBytes []byte, lenKey int) []byte {
	if lenKey < 1 {
		lenKey = 1
	} else if lenKey > 10 {
		lenKey = 10
	}
	encodedBytes, clave := buildBytes(lenKey, encodedBytes, nil, true)
	var decodedBytes []byte
	for _, b := range encodedBytes {
		decodedBytes = append(decodedBytes, bynaryChangerBytes(clave, b, true))
	}
	return decodedBytes
}
func bynaryChangerBytes(claves []byte, databyte byte, dencode bool) byte {
	indextoinvert := []int{}
	for _, clave := range claves {
		var positions []int
		for i, char := range fmt.Sprintf("%08b", clave) {
			if char == '1' {
				positions = append(positions, i)
			}
		}
		indextoinvert = append(indextoinvert, positions...)
		if dencode {
			for i, j := 0, len(indextoinvert)-1; i < j; i, j = i+1, j-1 {
				indextoinvert[i], indextoinvert[j] = indextoinvert[j], indextoinvert[i]
			}
		}
	}
	byteToString := fmt.Sprintf("%08b", databyte)
	bits := []rune(byteToString)
	for _, index := range indextoinvert {
		if index >= 0 && index < 8 {
			if bits[index] == '1' {
				bits[index] = '0'
			} else {
				bits[index] = '1'
			}
		}
	}
	newByteString := string(bits)
	newbyte, err := strconv.ParseUint(newByteString, 2, 8)
	if err != nil {
		fmt.Println("Error converting binary string to byte:", err)
		return 0
	}

	return byte(newbyte)
}
func buildBytes(bytesToUse int, bytesEncoded []byte, clave []byte, decode bool) ([]byte, []byte) {
	switch true {
	case bytesToUse <= 4 && decode:
		return bytesEncoded[:len(bytesEncoded)-bytesToUse], bytesEncoded[len(bytesEncoded)-bytesToUse:]
	case bytesToUse <= 4:
		return append(bytesEncoded, clave...), nil
	case bytesToUse <= 7 && decode:
		return bytesEncoded[bytesToUse-4 : len(bytesEncoded)-4], append(bytesEncoded[len(bytesEncoded)-4:], bytesEncoded[:bytesToUse-4]...)
	case bytesToUse <= 7:
		bytes := append(clave[4:], bytesEncoded...)
		return append(bytes, clave[:4]...), nil
	case bytesToUse > 7 && decode:
		midelbytes := bytesEncoded[:3]
		firstbytes := bytesEncoded[len(bytesEncoded)-4:]
		lastbytes := bytesEncoded[len(bytesEncoded)-(bytesToUse-3) : len(bytesEncoded)-4]

		return bytesEncoded[3 : len(bytesEncoded)-(bytesToUse-3)], append(append(firstbytes, midelbytes...), lastbytes...)
	case bytesToUse > 7:
		bytes := append(clave[4:7], bytesEncoded...)
		bytes = append(bytes, clave[7:]...)
		return append(bytes, clave[:4]...), nil
	}
	return nil, nil
}
