FROM rust:1.89 AS build

WORKDIR /app

COPY . .

RUN cargo install --path .
RUN cargo build --release

FROM debian:stable-slim

COPY --from=build /app/target/release/rust-axum /server

CMD ["/server"]
