
# sistem manajemen tugas

Ide ini saya dapat dari chat gpt. kerena saya kekurangan ide jadi saya tanya chat gpt.
## Installation

Install my-project

```bash
  cd Sistem-Manajemen-Tugas
  go mod tidy
  go run main.go
```

note: jangan lupa untuk membuat database terlebih dahulu sebelum menjalankan
    
## API Reference

#### signup

```http
  POST /v1/signup
```

| body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**
 `password` | `string` | **Required**  
  `email` | `email` | **Required** |


