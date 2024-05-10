# GalaxyEncoder
Codificador de datos
Fue diseñado principalmente para mantener la integridad de datos simples en UTF-8 solo para traspaso de informacion con caracteres comunes, principalemnte lo diseñe de forma que obteniendo un texto de entrada y una cantidad de bytes a elección entre 1 y 10 se generan de forma aleatoria y usando eso como clave inicial hace un bitflip de cada cada byte en la cadena de texto original, esta secuencia sucede de forma que se listan todas las posiciones en 1 dentro de cada byte a usar como clave, una vez teniendo estas posiciones se invierten en dichas posiciones cada byte del texto de entrada y dando un resultado en hexadecimal, de esta forma me asguro de manejar este tipo de string en lugar de preocuparme si el texto de entrada tenia algun caracter que pueda ser interpretado de otra forma en el transcurso del manejo de dicho string, ya sea por la base de datos o el codigo usado, ademas me aseguro de reintegrar la informacion de forma correcta para volver a representarla en pantalla una vez sea decodificada, la clave que se genero se agrega al texto resultante codificado lo que implica que su reconstruccion es simplemente obtener la clave y hacer el flipbit a la inversa. sin embargo la posicion en como se contruye el texto codificiado es facilmente alterable haciendo que si se necesita puede ordanrse en la forma que se quiera para obtener funciones mas personalizadas a su necesidad o directamente no insertando la clave y almacenandola en otro sitio.

A futuro me gustaria integrar mas funciones para UTF-16 y UTF-32 abarcando mas caracteres segun se necesiten.

## Instalar galaxyencoder

Para instalar usando Go:

```go
go get github.com/GaryJAMZ/GalaxyEncoder/go
```
Para agregar en TS:
```bash
npm i galaxyencoder
```
