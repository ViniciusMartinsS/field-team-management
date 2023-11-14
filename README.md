# Field Team Management

This service is dedicated to overseeing field teams, offering a comprehensive platform for technical professionals to seamlessly create, update, and monitor their assigned tasks. For managers, it delivers a range of features, including the ability to view a consolidated list of tasks for all technicians, receive real-time notifications upon task completion, and efficiently remove tasks that are no longer required.

## API Documentation

<details>
  <summary><b>Authentication</b></summary>

  </br>

  > **Handles API Authentication**

  #### URL
  `/v1/auth`

  #### Method
  `POST`

  #### Data Params
  ```json
  {
      "email": "example@example.io",
      "password": "123456"
  }
  ```

  * `email` **Required**
  * `password` **Required**

  #### Success Response
  **HTTP Status Code** `200`
  ```json
  {
    "status": true,
    "result": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
  ```

  #### Error Response
  **HTTP Status Code** `401`
  ```json
  {
    "status": false,
    "error": "unauthorized"
  }
  ```

  **HTTP Status Code** `400`
  ```json
  {
    "status": false,
    "error": "malformed request"
  }
  ```

  **HTTP Status Code** `500`
  ```json
  {
    "status": false,
    "error": "internal server error"
  }
  ```

  #### Try it out
  ```bash
  curl --location 'http://localhost:8080/v1/auth' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "email": "example@example.io",
    "password": "123456"
  }'
  ```

  <sub>

  **⚠️ Credentials**

  ```
  Technician 01
    Email: joe.doe@example.com
    Password: 123456
    Already Created Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvZS5kb2VAZXhhbXBsZS5jb20iLCJyb2xlX2lkIjoyLCJ1c2VyX2lkIjoyfQ.yt8K34ETTo0mYHmZ0VvrzXljqkMsmFnVc1SEkBLgEMs

  Technician 02
    Email: janne.biu@example.com
    Password: S
    Already Created Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imphbm5lLmJpdUBleGFtcGxlLmNvbSIsInJvbGVfaWQiOjIsInVzZXJfaWQiOjN9.X5LKqEhQUzY_VbsoH6cxw3te86wssD-x1fVVgLfUCKQ

  Manager
    Email: jedi.carson@example.com
    Password: 123456
    Already Created Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImplZGkuY2Fyc29uQGV4YW1wbGUuY29tIiwicm9sZV9pZCI6MSwidXNlcl9pZCI6MX0.8mqJOj9oaHguRlZOV3x77mADXNDBmdUL4ptMFhdA2t4
  ```

  </sup>

</details>

<details>
  <summary><b>List Tasks</b></summary>

  </br>

  > **Shows all tasks of a technician**

  #### URL
  `/v1/tasks`

  #### Method
  `GET`

  #### Authorization
  `Bearer Token`

  * `token` **Required**

  #### Success Response
  **HTTP Status Code** `200`
  ```json
  {
      "status": true,
      "result": [
          {
              "id": 1,
              "summary": "hello",
              "date": "15/11/2023 18:00",
              "user_id": 2
          }
      ]
  }
  ```

**HTTP Status Code** `200`
  ```json
  {
      "status": true,
      "result": null
  }
  ```

  #### Error Response
  **HTTP Status Code** `400`
  ```json
  {
    "status": false,
    "message": "tasks not found"
  }
  ```

  **HTTP Status Code** `500`
  ```json
  {
    "status": false,
    "error": "internal server error"
  }
  ```

  #### Try it out
  ```bash
  curl --location 'http://localhost:8080/v1/tasks' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'
  ```

</details>

<details>
  <summary><b>Create Task</b></summary>

  </br>

  > **Creates a task for a specific technician**

  #### URL
  `/v1/tasks`

  #### Method
  `POST`

  #### Data Params
  ```json
  {
    "summary": "This is a new task",
    "date": "15/11/2023 15:04"
  }
  ```

  * `summary` **Required**
  * `date` **Optional - DD/MM/YYYY HH:MM**

  #### Authorization
  `Bearer Token`

  * `token` **Required**

  #### Success Response
  **HTTP Status Code** `200`
  ```json
  {
    "status": true,
    "result": {
        "id": 3,
        "summary": "This is a new task",
        "date": "15/11/2023 15:04",
        "user_id": 2
    }
  }
  ```

  #### Error Response
  **HTTP Status Code** `400`
  ```json
  {
    "status": false,
    "error": "malformed request"
  }
  ```

  **HTTP Status Code** `403`
  ```json
  {
    "status": false,
    "error": "not allowed to perform this action"
  }
  ```

  **HTTP Status Code** `500`
  ```json
  {
    "status": false,
    "error": "internal server error"
  }
  ```

  #### Try it out
  ```bash
  curl --location 'http://localhost:8080/v1/tasks/' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9s' \
  --data '{
    "summary": "This is a new task",
    "date": "15/11/2023 15:04"
  }'
  ```

</details>

<details>
  <summary><b>Update Task</b></summary>

  </br>

  > **Updates a task of technician**

  #### URL
  `/v1/tasks/:id`

  #### Method
  `PATCH`

  #### Data Params
  ```json
  {
    "summary": "Hello World",
    "date": "15/11/2023 18:04"
  }
  ```

  * `summary` **Optional**
  * `date` **Optional - DD/MM/YYYY HH:MM**

  #### Authorization
  `Bearer Token`

  * `token` **Required**

  #### Success Response
  **HTTP Status Code** `200`
  ```json
  {
    "status": true,
    "result": {
        "id": 1,
        "summary": "Hello World",
        "date": "15/11/2023 18:04",
        "user_id": 2
    }
  }
  ```

  #### Error Response
  **HTTP Status Code** `400`
  ```json
  {
    "status": false,
    "error": "malformed request"
  }
  ```

  **HTTP Status Code** `400`
  ```json
  {
    "status": false,
    "error": "tasks not found"
  }
  ```

  **HTTP Status Code** `403`
  ```json
  {
    "status": false,
    "error": "not allowed to perform this action"
  }
  ```

  **HTTP Status Code** `500`
  ```json
  {
    "status": false,
    "error": "internal server error"
  }
  ```

  #### Try it out
  ```bash
  curl --location --request PATCH 'http://localhost:8080/v1/tasks/1' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9' \
  --data '{
    "summary": "Hello World",
    "date": "2023/11/18 18:04"
  }'
  ```

</details>

<details>
  <summary><b>Delete Task</b></summary>

  </br>

  > **[MANAGER ONLY] Deletes task of a technician**

  #### URL
  `/v1/tasks/:id`

  #### Method
  `DELETE`

  #### Authorization
  `Bearer Token`

  * `token` **Required**

  #### Success Response
  **HTTP Status Code** `204`

  #### Error Response
  **HTTP Status Code** `400`
  ```json
  {
    "status": false,
    "error": "malformed request"
  }
  ```

  **HTTP Status Code** `403`
  ```json
  {
    "status": false,
    "error": "not allowed to perform this action"
  }
  ```

  **HTTP Status Code** `500`
  ```json
  {
    "status": false,
    "error": "internal server error"
  }
  ```

  #### Try it out
  ```bash
  curl --location --request DELETE 'http://localhost:8080/v1/tasks/3' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'
  ```

</details>

## Developer Guideline

**Clone Repository**
```
$ git@github.com:ViniciusMartinsS/field-team-management.git
```

### Docker

**Spin Up Application**
```
$ make docker-up
```

**Stop Application**
```
$ make docker-down
```

### Kubernetes Local Cluster (K3D & KUBECTL)

 > Please note that to run the upcoming steps, you must have **[K3D](https://k3d.io/v5.4.6/)** & **[KUBECTL](https://kubernetes.io/docs/reference/kubectl/)** installed

**Spin Up Application**
```
$ make kubernetes
```

**Application Readiness**
```
$ make kubernetes-pod-status
```
<sub> You can only request the API once all pods are **READY (1/1)** </sub>

**Stop Application**
```
$ make kubernetes-stop
```

### Local

**Copy and Change Configs**
```
$ cp run.local.api.sh.sample  run.local.api.sh
```

**Spin Up Application**
```
$ ./run.local.api.sh
```