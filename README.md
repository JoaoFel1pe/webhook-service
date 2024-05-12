# Golang Webhook Service

## 🌐 Overview

This project aims to illustrate the creation of a webhook service using Golang for an imaginary payment gateway, with an API built in Flask providing support.

When a request is made on the payment gateway to initiate a payment, a response is sent. Additionally, the webhook service, written in Golang, is contacted to send a webhook request. Communication between the Golang service and the Flask API is facilitated through a Redis channel. This channel transmits data and directs the Golang service to send a request (webhook) to the URL provided in the payload.

## 📑 Table of Contents

- [Adding the Queuing Logic](#-adding-the-queuing-logic)
- [Architecture](#-architecture)
- [Setup and Installation](#-setup-and-installation)
- [Usage](#-usage)
- [Logs](#-logs)
- [Dependencies](#-dependencies)
- [Contributors](#-contributors)
- [License](#-license)

## ➕ Adding the Queuing Logic

Queuing logic is essential for managing data flow efficiently. In this project, we utilize channel-based queuing in Golang for its concurrency, capacity control, and blocking/non-blocking operation capabilities.

## 🏗️ Architecture

The architecture of the solution is straightforward:

- **API (Payment Gateway)**: Accepts requests on an endpoint and returns a payload. Additionally, it sends the payload through a Redis channel called `payments`.
  
- **Webhook Service**: Listens to the `payments` Redis channel. When data is received, it formats the payload to be sent to the URL indicated and implements a retry mechanism using Golang channel queuing and exponential backoff in case of request failures.

## 🛠️ Setup and Installation

To run this project, follow these steps:

1. Get a free webhook URL at [webhook.site](https://webhook.site).
   
2. Create a file named `.env` at the root of the project alongside `docker-compose.yaml`. Add the following content:

    ```
    REDIS_ADDRESS=redis:6379
    WEBHOOK_ADDRESS=<WEBHOOK_ADDRESS>
    ```

    Replace `<WEBHOOK_ADDRESS>` with the webhook URL obtained from [webhook.site](https://webhook.site).
   
3. Build and start the containers using Docker Compose:

    ```bash
    docker-compose up -d --build
    ```

4. Once the build is finished, track the logs of the webhook service using:

    ```bash
    docker-compose logs -f webhook
    ```

## 🚀 Usage

- Access `http://127.0.0.1:8000/payment` to start sending data to the webhook service through Redis.
  
## 📝 Logs

In case of success, here is a sample of what the logs will look like:

``` bash
api                                     | 172.19.0.1 - - [12/May/2024 15:00:03] "GET /payment HTTP/1.1" 200 -
d04fe2946a28_webhook-service-webhook-1  | 2024/05/12 15:00:04 delivered
```

## 📦 Dependencies
### Development
![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

### Frameworks
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white)

### Deployment
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## 👥 Contributors

João Felipe Barbosa da Silva (jfbs@cin.ufpe.br 📧)

## 📄 License

This project is licensed under the [MIT License](LICENSE).
