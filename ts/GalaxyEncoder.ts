export function GalaxyEncoder(text: string, bytesToUse: number): string {
    const clave = Array.from({ length: bytesToUse }, () => Math.floor(Math.random() * 256));
    const results = buildText(text.split('').map((char) => {
        return binaryChanger(clave, char.charCodeAt(0), false).toString(16).padStart(2, '0'); // Convierte a hexadecimal
    }).join(''), clave, bytesToUse,false);
    return results[0];
}
export function GalaxyDecoder(textencode: string, bytesToUse: number): string {
    const results = buildText(textencode, [], bytesToUse, true);
    let clavehex: number[] = [];
    for (let i = 0; i < results[1].length; i += 2) {
        clavehex.push(parseInt(results[1].slice(i, i + 2), 16));
    }
    var decodedTextInHex: string = "";
    for (let i = 0; i < results[0].length; i += 2) {
        decodedTextInHex += String.fromCharCode(binaryChanger(clavehex, parseInt(results[0].slice(i, i + 2), 16), true));
    }
    return decodedTextInHex;
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
function buildText(text: string, clave: number[], bytesToUse: number,decode: boolean): [string, string]{
    switch (true) {
        case bytesToUse <= 4 && decode:
            return [text.slice(0, -(bytesToUse * 2)), text.slice(-(bytesToUse * 2))];
        case bytesToUse <= 4:
            return [text + Array.from(clave, (byte) => byte.toString(16).padStart(2, '0')).join(''), ""];
        case bytesToUse <= 7 && decode:
            return [text.slice((bytesToUse*2-8), -8), text.slice(-8) + text.slice(0, (bytesToUse*2-8))]
        case bytesToUse <= 7:
            return [Array.from(clave.slice(4), (byte) => byte.toString(16).padStart(2, '0')).join('') + text + Array.from(clave.slice(0, 4), (byte) => byte.toString(16).padStart(2, '0')).join(''), ""];
        case bytesToUse <= 10 && decode:
            return [text.slice(6, -(bytesToUse*2-6)), text.slice(-8) + text.slice(0, 6) + text.slice(-(bytesToUse*2-6), -8)];
        case bytesToUse <= 10:
            return [Array.from(clave.slice(4, 7), (byte) => byte.toString(16).padStart(2, '0')).join('') + text + Array.from(clave.slice(7), (byte) => byte.toString(16).padStart(2, '0')).join('')  + Array.from(clave.slice(0, 4), (byte) => byte.toString(16).padStart(2, '0')).join(''), ""];
    }
    return ["", ""];
}