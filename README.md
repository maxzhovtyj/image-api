## 👋  Summary

The goal is to create an HTTP API for uploading, optimizing, and serving images.

1. The API should expose an endpoint for uploading an image. After uploading, the image should be sent for optimization via a queue (e.g. RabbitMQ) to prevent excessive system load in case of many parallel image uploads and increase the system durability.
2. Uploaded images should be taken from the queue one by one and optimized using the `github.com/h2non/bimg` go package (or `github.com/nfnt/resize` package). For each original image, three smaller-size image variants should be generated and saved, with 75%, 50%, and 25% quality.
3. The API should expose an endpoint for downloading an image by ID. The endpoint should allow specifying the image optimization level using query parameters (e.g. `?quality=100/75/50/25`).

## Run application

Create .env file with the following values
```dotenv
HTTP_HOST=localhost
HTTP_PORT=8000

# Change AMQP_HOST=localhost if you run the app not in via docker-compose

AMQP_USER=guest
AMQP_PASSWORD=guest
AMQP_HOST=rabbitmq
AMQP_PORT=5672
```

Use make to run docker-compose
```shell
make appUp
```

OR

Run rabbitmq instance
```shell
docker run -d --name softcery-rabit -p 5672:5672 rabbitmq:3.10.7-management
```

And run command
```shell
make run
```

## Possible improvements

* Add tests
* Add more logs
* Use interfaces for consumer and publisher 
* Add support for other image extensions