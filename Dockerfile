# docker build . -t ghcr.io/reddio-com/itachi:latest
FROM golang:1.22-bookworm as builder

RUN apt-get update && apt-get install -y curl gcc
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

RUN mkdir /build
COPY . /build
RUN cd /build && git submodule init && git submodule update --recursive --checkout && make build

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

COPY ./conf /conf
RUN mkdir /cairo_db /yu
COPY --from=builder /build/itachi /itachi

CMD ["/itachi"]
