## 👋  Summary

The goal is to create an HTTP API for uploading, optimizing, and serving images.

1. The API should expose an endpoint for uploading an image. After uploading, the image should be sent for optimization via a queue (e.g. RabbitMQ) to prevent excessive system load in case of many parallel image uploads and increase the system durability.
2. Uploaded images should be taken from the queue one by one and optimized using the `github.com/h2non/bimg` go package (or `github.com/nfnt/resize` package). For each original image, three smaller-size image variants should be generated and saved, with 75%, 50%, and 25% quality.
3. The API should expose an endpoint for downloading an image by ID. The endpoint should allow specifying the image optimization level using query parameters (e.g. `?quality=100/75/50/25`).