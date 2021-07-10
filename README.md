I have some programming ideas that needs a topic to practice. I think a music serive is a nice one.

Backend: server to (up/down) streaming music:
- API: gRPC
- [Middleware](https://github.com/grpc-ecosystem/go-grpc-middleware): interceptor + per API callOption.
> - secure header
> - logger
> - panic recovery
- [Authentication](https://grpc.io/docs/guides/auth/)
- Ignore cookie (which cause problems such as CSRF), use header or body.

Frontend: App to browse/listen music:
- [Clean architecture](https://github.com/ResoCoder/flutter-tdd-clean-architecture-course)
- App
- Flutter for Web
> - Call gRPC API: http/2 issue.
> - local storage to header or body.

Others:
- Test: unit, API, end-to-end, integration.
- CI/CD
- Scrum: 1-week sprint.

## Week 1: 7/11 - 7/17

Backend:

- DB: 
> - Postgres user.
> - Create music table.
> - Postgres with [sqlx](https://github.com/jmoiron/sqlx)
- API:
> - create music: title, artist, year, genres.
> - list all music.

Frontend:
- Setup dev environment: Windows 10.
- App:
> - create music
> - list all music
- For Web:
> - create music
- list all music

