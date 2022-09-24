# API Data Wilayah Indonesia
- Provinsi
- Kabupaten/kota
- Kecamatan
- Desa/Kelurahan

## DATA SOURCE

- [https://github.com/ArrayAccess/Indonesia-Postal-And-Area](https://github.com/ArrayAccess/Indonesia-Postal-And-Area)
- [Badan Pusat Statistik Indonesia (BPS)](https://bps.go.id/)

## TODO

- [ ] Fix Data Province
- [ ] Query Search (id, sub_id, postal)
- [ ] Auto Migration DB (Mysql, Postgres, MongoDB, ElasticSearch, Firebase)

## Path

```txt
[GET]   /provinces
[GET]   /regencies
[GET]   /districts
[GET]   /villages

Search :
    param ?name=
```

## Response

```json
{
  "total": 34,
  "data": [
    {
      "id": 11,
      "name": "ACEH",
      "meta_data": {
        "latitude": "4.0405070151375",
        "longitude": "96.649002928244",
        "postal": [
          "20000",
          "23000",
          "24000"
        ]
      }
    }
  ]
}
```