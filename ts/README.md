# GalaxyEncoder
Fue diseñado principalmente para mantener la integridad de datos y añadir una ofuscación a los mismos, orignalmente solo se manejaban strings, sin embargo ahora se permiten la insercion de cadenas de bytes directamente de forma que obteniendo un array de bytes como entrada y una longitud de clave a elección entre 1 y 10 se generan de forma aleatoria bytes y usando eso como clave inicial hace un bitflip de cada byte en los bytes originales, esta secuencia sucede de forma que se listan todas las posiciones en 1 dentro de cada byte a usar como clave, una vez teniendo estas posiciones se invierten en dichas posiciones de cada byte del array de entrada y dando un resultado de bytes encodeados, la clave que se genero se agrega a los bytes resultantes codificados lo que implica que su reconstruccion es simplemente obtener la clave y hacer el flipbit a la inversa.

## Update 2.0
Se cambiaron las funciones basicas para manejar arrays de bytes en lugar de string, esto amplia el uso de estas funciones permitiendo directamente codificar cualquier cadena de bytes, asi mismo la compatibilidad con las anteriores funciones y sus resultados se manejan por medio de las funciones llamadas GalaxyHexEncoder y GalaxyHexDecoder las cuales funcionan igual que las anteriores funcones de la version 1.

## Instalar galaxyencoder

```bash
npm i galaxyencoder
```
