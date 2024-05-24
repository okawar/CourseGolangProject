FROM alpine:latest
RUN apk add --no-cache go
COPY . /golang_pr
WORKDIR /golang_pr
RUN rm golang_pr.exe
RUN go build golang_pr
EXPOSE 8090
CMD ["./golang_pr"]