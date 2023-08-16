FROM debian:bullseye-20230725

# more info at https://docs.github.com/en/packages/learn-github-packages/connecting-a-repository-to-a-package#connecting-a-repository-to-a-container-image-using-the-command-line
LABEL org.opencontainers.image.source=https://github.com/gsantoro/go-sample-app

COPY build/hello /app
CMD ["/app"]