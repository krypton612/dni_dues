<p align="center">
  <a href="https://example.com/">
       <img src="https://www.ecipartners.com/wp-content/uploads/2021/12/IoT-review.jpg" alt="Logo" width=72 height=72>

  </a>

  <h3 align="center">Logo</h3>

  <p align="center">
    Short description
    <br>
    <a href="https://github.com/krypton612/dni_dues/issues/new?template=bug.md">Report bug</a>
    ·
    <a href="https://github.com/krypton612/dni_dues/issues/new?template=feature.md&labels=feature">Request feature</a>
  </p>
</p>


## Table of contents

- [Quick start](#quick-start)
- [Status](#status)
- [Start](#start)
- [Creators](#creators)
- [Thanks](#thanks)
- [Copyright and license](#copyright-and-license)

## Quick start

Some text

- you must first compile the script in go that contains the program
- Second, it should generate a token using the credentials that only the owner of the project can grant
- third party could do the queries without problems

## Status

This repository allows you to search for information on students from the Bolivian military college of engineering


```text
├── go.mod
├── go.sum
├── main.go
└── README.md
```
## Start
```
Usage of ./main:
  -file string
    	Envia la ruta de un archivo json alternativo (default "None")
  -generate string
    	Genera lo necesario para realizar las consultas [json or string] (default "None")
  -search string
    	Busca un cliente en la api "Nombre apellido_pat apellido_mat" (default "None")
```
./main -generate string

```json

{
  "token": "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJlbWkiLCJleHAiOjE2NTkzMTMwODJ9.2v4DrLXNeypdeI3pQaCOZOWNwVXzm30wSaVQR6cjBSh_2MRyWx0AKKrQUsJNZG0YB422wXvYt73DUMINZ57Ixg",
  "username": "emi",
  "publicId": "7KZ7Yp70vwU4tCRe",
  "sessionId": "-1120421761202207211591",
  "moduleId": "91",
  "orgId": "9ozphlqx",
  "merchantId": "redenlace_429777",
  "logo": "logo-emi.png",
  "code": 100
}    
```
./main -search "efrain colque calizaya"

```json
{
  "code": 100,
  "items": [
    {
      "codError": 0,
      "mensaje": "",
      "cuenta": [
        {
          "cuenta": "13130204;1;102/20;106",
          "descServicio": "Recaudacion deuda EMI en BS.",
          "detalle": "CIENCIAS BASICAS-Oficina Central EMI",
          "moneda": 1,
          "nombre": "COLQUE CALIZAYA EFRAIN",
          "servicio": 1
        },
        {
          "cuenta": "13130204;1;101/22;101",
          "descServicio": "Recaudacion deuda EMI en BS.",
          "detalle": "ING. SISTEMAS-Oficina Central EMI",
          "moneda": 1,
          "nombre": "COLQUE CALIZAYA EFRAIN",
          "servicio": 1
        },
        {
          "cuenta": "13130204;1;102/21;130",
          "descServicio": "Recaudacion deuda EMI en BS.",
          "detalle": "INGLES-Oficina Central EMI",
          "moneda": 1,
          "nombre": "COLQUE CALIZAYA EFRAIN",
          "servicio": 1
        },
        {
          "cuenta": "13130204;4;4MOODLE-1;441",
          "descServicio": "Recaudacion deuda EMI en BS.",
          "detalle": "OTROS CURSOS ESPECIALES-Unidad Academica - Cochabamba",
          "moneda": 1,
          "nombre": "COLQUE CALIZAYA EFRAIN",
          "servicio": 1
        }
      ],
      "fechaOperativa": 20220721,
      "nroOperacion": 71878
    }
  ]
}
```

## Creators

**Creator 1**

- <https://github.com/krypton612>

## Thanks

thank you, to all who support me, I will not let you down

## Copyright and license

Code and documentation copyright 2011-2022 the authors. Code released under the [MIT License](https://github.com/krypton612/dni_dues/blob/master/LICENSE).

Enjoy :metal:
