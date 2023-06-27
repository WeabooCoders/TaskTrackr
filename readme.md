
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
  POST api/v1/signup
```

| body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**
 `password` | `string` | **Required**  
  `email` | `email` | **Required** |


#### sigin

```http
  POST api/v1/sigin
```

| body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `email` | **Required**
 `password` | `string` | **Required**  


#### create task

```http
  POST api/v1/task
```

###### tambahkan header Authorization dan nilai nya adalah token <br>


| body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `title` | `string` | **Required**
 `description` | `string` | **Required**  
 `status` | `string` | **Required**  

#### Get task

```http
  GET api/v1/task
```

###### tambahkan header Authorization dan nilai nya adalah token <br>



## Features

- signup and sigin
- create task and get task

