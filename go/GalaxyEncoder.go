package GalaxyEncoder

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func GalaxyEncoder(text string, bytesToUse int) (string, error) {
	var encodedText, encodedTextInHex string
	if bytesToUse > 10 || bytesToUse < 1 {
		return "", fmt.Errorf("bytes to use must be between 1 and 10")
	}
	clave := make([]byte, bytesToUse)
	rand.Read(clave)
	for _, char := range text {
		encodedTextInHex += fmt.Sprintf("%02x", binaryChanger(clave, fmt.Sprintf("%08b", char), false))
	}
	encodedText, _ = buildText(bytesToUse, encodedTextInHex, clave, false)
	return encodedText, nil
}
func GalaxyDecoder(textencode string, bytesUsed int) (string, error) {
	var hexText, decodedText string
	if bytesUsed > 10 || bytesUsed < 1 {
		return "", fmt.Errorf("bytes to use must be between 1 and 10")
	}
	hexText, clavesToAdd := buildText(bytesUsed, textencode, []byte{}, true)
	clave := []byte{}
	for _, claveHex := range clavesToAdd {
		for i := 0; i < len(claveHex); i += 2 {
			resp, _ := strconv.ParseInt(claveHex[i:i+2], 16, 64)
			clave = append(clave, byte(resp))
		}
	}
	for i := 0; i < len(hexText); i += 2 {
		hexByte := hexText[i : i+2]
		charcode, _ := strconv.ParseInt(hexByte, 16, 64)
		decodedText += string(binaryChanger(clave, fmt.Sprintf("%08b", charcode), true))
	}
	return decodedText, nil
}
func binaryChanger(claves []byte, charBinary string, dencode bool) int {
	for i := 0; i < len(claves)-1; i++ {
		for j := i + 1; j < len(claves); j++ {
			if claves[i] > claves[j] {
				claves[i], claves[j] = claves[j], claves[i]
			}
		}
	}
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
	for _, position := range indextoinvert {
		bit := "1"
		if charBinary[position] == '1' {
			bit = "0"
		}
		charBinary = charBinary[:position] + bit + charBinary[position+1:]
	}
	result, _ := strconv.ParseInt(charBinary, 2, 64)
	return int(result)
}
func buildText(bytesToUse int, text string, clave []byte, decode bool) (string, []string) {
	var clavesToAdd []string
	switch true {
	case bytesToUse <= 4 && decode:
		return text[:len(text)-(bytesToUse*2)], append(clavesToAdd, text[len(text)-(bytesToUse*2):])
	case bytesToUse <= 4:
		return text + fmt.Sprintf("%x", clave), nil
	case bytesToUse <= 7 && decode:
		return text[(bytesToUse-4)*2 : len(text)-8], append(clavesToAdd, text[len(text)-8:], text[:(bytesToUse-4)*2])
	case bytesToUse <= 7:
		return fmt.Sprintf("%x", clave[4:]) + text + fmt.Sprintf("%x", clave[:4]), nil
	case bytesToUse <= 10 && decode:
		return text[6 : len(text)-(8+(bytesToUse*2-14))], append(clavesToAdd, text[len(text)-8:], text[:6], text[len(text)-(8+(bytesToUse*2-14)):len(text)-8])
	case bytesToUse <= 10:
		return fmt.Sprintf("%x", clave[4:7]) + text + fmt.Sprintf("%x", clave[7:]) + fmt.Sprintf("%x", clave[:4]), nil
	}
	return "", nil
}
