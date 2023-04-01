# https://github.com/GoogleContainerTools/distroless
FROM debian:11.6

COPY build/hello /app
CMD ["/app"]