export function GalaxyHexEncoder(textToEncode: string, lenKey: number): string {
    const encoder = new TextEncoder();  
    const encoded = encoder.encode(textToEncode);
    const results = GalaxyEncoder(Array.from(encoded), lenKey);
    return results.map((char) => char.toString(16).padStart(2, '0')).join('');
}
export function GalaxyHexDecoder(encodedText: string, lenKey: number): string {
    const matchResult = encodedText.match(/.{1,2}/g);
    const encoded = matchResult ? matchResult.map((char) => parseInt(char, 16)) : [];
    const results = GalaxyDecoder(encoded, lenKey);
    const decoder = new TextDecoder();
    return decoder.decode(new Uint8Array(results));
}
export function GalaxyEncoder(bytesToEncode: number[], lenKey: number): number[] {
    lenKey = lenKey < 1 ? 1 : lenKey > 10 ? 10 : lenKey;
    const clave = Array.from({ length: lenKey }, () => Math.floor(Math.random() * 256));
    const results = buildText(bytesToEncode.map((char) => {
        return binaryChanger(clave, char, false);
    }), clave, lenKey, false);

    return results[0]
}
export function GalaxyDecoder(encodedBytes: number[], lenKey: number): number[] {
    lenKey = lenKey < 1 ? 1 : lenKey > 10 ? 10 : lenKey;
    console.log(encodedBytes)
    const results = buildText(encodedBytes, [], lenKey, true);
    let clavehex: number[] = results[1];
    console.log(clavehex)
    console.log(lenKey)
    var decodedText: number[] = [];
    for (let i = 0; i < results[0].length; i += 1) {
        decodedText.push(binaryChanger(clavehex, results[0][i], true));
    }
    return decodedText;
}
function binaryChanger(clave: number[],char: number, decode: boolean): number{
    var sortedClave = [...clave].sort((a, b) => b - a);
    let clavesBinarys: string[] = [];
    for (let i = 0; i < sortedClave.length; i++) {
        clavesBinarys.push(sortedClave[i].toString(2).padStart(8, '0'));
    }
    var charBinary: string = char.toString(2).padStart(8, '0');
    let indextoinvert: number[] = [];
    for (let i = 0; i < clavesBinarys.length; i++) {
        for (let j = 0; j < 8; j++) {
            clavesBinarys[i][j] === '1' ? indextoinvert.push(j) : null;
        }
    }
    decode? indextoinvert.slice().reverse(): indextoinvert;
    for (const position of indextoinvert) {
        charBinary = charBinary.slice(0, position) + (charBinary[position] === '1' ? '0' : '1') + charBinary.slice(position + 1);
    }
    return parseInt(charBinary, 2);
}
function buildText(text: number[], clave: number[], bytesToUse: number,decode: boolean): [number[], number[]]{
    switch (true) {
        case bytesToUse <= 4 && decode:
            return [text.slice(0, -bytesToUse), text.slice(-bytesToUse)];
        case bytesToUse <= 4:
            return [(text.push(...clave), text), []];
        case bytesToUse <= 7 && decode:
            return [text.slice((bytesToUse-4), -4), text.slice(-4).concat(text.slice(0, (bytesToUse-4)))];
        case bytesToUse <= 7:
            return [clave.slice(4).concat(text).concat(clave.slice(0, 4)), []];
        case bytesToUse <= 10 && decode:
            return [text.slice(3, -(bytesToUse-3)), text.slice(-4).concat(text.slice(0, 3)).concat(text.slice(-(bytesToUse-3), -4))];
        case bytesToUse <= 10:
            return [clave.slice(4, 7).concat(text).concat(clave.slice(7)).concat(clave.slice(0, 4)), []];
    }
    return [[], []];
}
