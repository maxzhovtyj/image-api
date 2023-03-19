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

## **Недоліки/Покращення**

1. В чергу відправляється слайс байтів зображення.

   *Які можуть бути недоліки такого підходу?*

   *Як можна було б це реалізувати по-іншому?*

2. При відвантаженні зображення користувачу не повертається його ID. Як потім отримати це зображення, не використовуючи `/images-list` роут?
3. В статично-типізованих мовах програмування додавати тип змінної до назви - це погана звичка.
4. Робота з чергою не абстрагована, немає інтерфейсу та реалізації (про це ти написав у README).
5. Навіщо існує окремий сервіс `Publisher`? Чому б не використовувати інтерфейс `rabbitmq.Publisher`? Черга не може одночасно відноситись до шару транспорту і сервісу.
6. Чому стискання зображення знаходиться в шарі `repository`? Стискання - це бізнес-логіка, яка повинна знаходитись в сервісі. Необхідно додати інтерфейс `Resizer` з методом для оптимізації, який повинен викликатись в сервісі.

   Вигляд пакету, де в `resizer.go` знаходиться інтерфейс, а в інших файлах - реалізації, які імлементують цей інтерфейс. Таким чином в застосунку ми використовуємо лише інтерфейс і можемо з легкістю підміняти імлементації.

7. В шарі domain знаходяться помилки, які використовуються на різних шарах. Проте, в domain повинні знаходитись лише критичні бізнес-сутності, такі як *Image* та можливі якості зображень (зараз вони оголошуються в `app.go`).
   *Чи правильно те, що одні і ті ж помилки шейряться між різними шарами?*
8. Валідація на розмір зображення знаходиться в шарі контроллера, чи правильно це?

## **Переваги:**

1. Використана чиста архітектура, хоча потребує деяких покращень.
2. Є розуміння, що черга - це шар транспорту.
3. Виділений окремий пакет для роботи з файловою системою і є шар repository, який би використовував цей пакет і реалізовував отримання/зберігання зображень.
4. Є `.gitignore`, тому у репозиторій потенційно не зможуть потрапити файли, які там не повинні там бути.
5. Описаний `docker-compose`.
6. Є невеликий `makefile`.
7. Є описаний README з інструкцією для запуску.
8. Є конфігураційний файл з потрібними значеннями.
9. Немає непотрібних коментарів.
10. Код не залитий одним комітом, дотримано принципу атомарності. Всі меседжі зрозумілі і несуть в собі сенс. Для стандартизації меседжів можна використовувати, наприклад: [https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/).

## **Література**

1. [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
2. [SOLID Go Design by Dave Cheney](https://dave.cheney.net/2016/08/20/solid-go-design)
3. [A primer on the clean architecture pattern and its principles](https://www.techtarget.com/searchapparchitecture/tip/A-primer-on-the-clean-architecture-pattern-and-its-principles)
4. [Clean code by Robert Martin](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship/dp/0132350882)
5. [Practical Go by Dave Cheney](https://dave.cheney.net/practical-go/presentations/qcon-china.html)