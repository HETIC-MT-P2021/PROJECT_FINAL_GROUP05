# Bot Discord - Carlos

## Prerequisite

- Golang v1.16

## Where am I ? 💡

This bot allows users in discord servers, edit and share medias from videos to audios

- To submit bug reports and feature suggestions, or track changes:
  https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/issues

Don't forget to read/use commits and PR conventions before doing any change !

## Start project

```sh
$ docker-compose up --build
```

```sh
$ docker exec -i db sh -c 'exec mysql -ugoproject -p"$MYSQL_PASSWORD"' < ./build/data/init.sql
```

To fill data in database, make the following HTTP request

**POST localhost:8080/api/v1/fixtures**

Now by executing the following HTTP request, data should be returned

**GET localhost:8080/api/v1/servers**

## Docs 📄

You can find our docs folder [here](https://github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP05/tree/main/docs)

## Swagger doc 📄

To generate swagger doc

```sh
$ swag init -g cmd/app/main.go
```

Go in [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to see API doc

## Features 📘 (incoming)

Extract a media from video or audio, which starts at X and ends at Y
Extract the audio track from a video, which starts at X and which ends at Y
Add a border to your discord avatar
Add a watermark to a picture
Reverse audio
Convert an image to a gif, adding a filter

Filters available (Lightning, Glitch, Rain, disco, 
- fadebegin (Fade to black at the start), 
- fadeend (Fade to black at the end), 
- fadeboth (Fade to black at the start and at the end), 
- carlos (it just posts a gif by Carlos) 😍)

## Main Libraries 📚 (incoming)

- [Go-Swagger](https://github.com/go-swagger/go-swagger)
- [Gin Gonic](https://github.com/gin-gonic/gin)

## Contributors ✨

<table>
  <tr>
    <td align="center"><a href="https://github.com/jibe0123"><img src="https://avatars.githubusercontent.com/u/13694014?s=400&u=979e9cdf62bcebe3e97740f83768fb41c8984a70&v=4" width="100px;" alt=""/><br /><sub><b>Jean Baptiste Agostin</b></sub></a><br /><a href="https://github.com/jibe0123" title="Developper">✏️</a>
    <td align="center"><a href="https://github.com/wyllisMonteiro"><img src="https://avatars2.githubusercontent.com/u/36091415?s=400&v=4" width="100px;" alt=""/><br /><sub><b>Wyllis Monteiro</b></sub></a><br /><a href="https://github.com/wyllisMonteiro" title="Developper">✏️</a>
    <td align="center"><a href="https://github.com/Para0234"><img src="https://avatars.githubusercontent.com/u/31258019?v=4" width="100px;" alt=""/><br /><sub><b>Benoit Covet</b></sub></a><br /><a href="https://github.com/Para0234" title="Developper">✏️</a></td>
    <td align="center"><a href="https://github.com/valmrt77"><img src="https://avatars0.githubusercontent.com/u/36480710?v=4" width="100px;" alt=""/><br /><sub><b>Valentin Moret</b></sub></a><br /><a href="https://github.com/valmrt77" title="Developper">✏️</a></td>
  </tr>
</table>

## License 📑

[MIT](https://github.com/HETIC-MT-P2021/CQRSES_GROUP4/blob/master/LICENSE)
