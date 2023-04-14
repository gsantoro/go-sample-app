# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian11:2f8ff53ec98f06e513f92b703bcd94cf47449cda

# more info at https://docs.github.com/en/packages/learn-github-packages/connecting-a-repository-to-a-package#connecting-a-repository-to-a-container-image-using-the-command-line
LABEL org.opencontainers.image.source=https://github.com/gsantoro/go-sample-app

COPY build/hello /app
CMD ["/app"]