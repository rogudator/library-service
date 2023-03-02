### Library Serivce
Данный сервис предоставляет доступ к данным из базы данных MySQL при помощи gRPC. Сервис запускается по-умолчанию на localhost:50051, если не указано иначе в configs/config.yaml.
Есть два метода gRPC: GetAuthorsOfBook (возвращает список авторов книги) и GetBooksByAuthor (возвращает список книг автора).
### "Ассортимент" библиотеки
В библиотеке есть:
- "Philosopher's Stone" by J.K. Rowling
- "Chamber of Secrets" by J.K. Rowling
- "Prisoner of Azkaban" by J.K. Rowling
- "Goblet of Fire" by J.K. Rowling
- "Order of the Phoenix" by J.K. Rowling
- "Half-Blood Prince" by J.K. Rowling
- "Deathly Hallows" by J.K. Rowling
- "Chemistry" by 
- "Advanced Chemistry" by
- "Organic Chemistry" by
- "Anna Karenina" by Leo Tolstoy
- "Russian Folktales"
А также сейчас "разобраны" книги от автора Mark Twain, но надеемся в ближайшем будущем расширить ассортимент.
## Как запустить сервис на своем компьютере
1. Первым делом нужно скопировать гит репозиторий в подходящию для вас папку.
```
git clone https://github.com/rogudator/library-service.git
```
2. И перейти в папку сервиса.
```
cd library-service
```
3. Для безопасности пароль от базы данных хранится в переменной окружения. db.env можно создать вот так:

UNIX/MAC: 
```
echo 'MYSQL_PASSWORD=securepassword'>db.env
```
Windows:
```
echo MYSQL_PASSWORD=securepassword>db.env
```
4. Теперь нужно сбилдить образ MySQL.
```
make build-db
```
5. После этого можно запустить контейнер базы данных.
```
make db
```
6. Загрузить зависимости.
```
go mod dowload
```
7. Запустить проект локально.
```
make local
```