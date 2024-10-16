# GalaxyEncoder

**GalaxyEncoder** fue diseñado principalmente para mantener la integridad de los datos y agregar un nivel de ofuscación. Inicialmente, solo trabajaba con strings, pero ahora permite procesar arrays de bytes directamente.

## ¿Cómo funciona?
GalaxyEncoder toma un array de bytes como entrada, y usando una clave de longitud variable (entre 1 y 10 bytes), genera aleatoriamente una secuencia de bytes. Estos bytes son utilizados como clave inicial para hacer un "bitflip" (inversión de bits) en los bytes originales del array.

### Detalles del proceso:
1. **Generación de la clave**: Se genera una clave aleatoria con la longitud que el usuario elige (de 1 a 10 bytes).
2. **Aplicación de la clave**: Para cada byte de la clave, se identifican las posiciones donde los bits están en 1.
3. **Inversión de bits**: Se invierten los bits en las mismas posiciones de los bytes del array de entrada.
4. **Resultado**: La clave generada se añade al array de bytes resultante, lo que permite revertir el proceso más adelante.

Este método asegura que los datos estén codificados, pero que puedan ser decodificados fácilmente recuperando la clave y aplicando el proceso de bitflip de manera inversa.

## Update 2.0
- Se ha mejorado el código para que ahora maneje arrays de bytes en lugar de strings, ampliando su uso a cualquier tipo de dato que pueda representarse como bytes.
- Para mantener compatibilidad con la versión anterior, se introdujeron las funciones **GalaxyHexEncoder** y **GalaxyHexDecoder**, que trabajan de manera similar a las funciones de la versión 1. Estas permiten codificar y decodificar strings en formato hexadecimal, preservando la funcionalidad previa.

## Instalación

### Go:
Para instalar GalaxyEncoder en Go:

```bash
go get github.com/GaryJAMZ/GalaxyEncoder/go

```
Para agregar en TS:
```bash
npm i galaxyencoder
```
### DIFERENCIAS ENTRE GO Y TS
- En Go, la función recibe el array de bytes ([]byte).
- En TypeScript, se usa un array de numeros (number[]), que representaran la cadena de bytes, esto con el fun de manejar de forma mas sencilla.
- En cualquier caso existe GalaxyHexEncoder y GalaxyHexDecoder, las cuales reciben string y retornan string, (al devolver un string en hexadecimal la cadena a codificar dobla su tamaño)