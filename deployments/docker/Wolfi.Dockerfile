# https://github.com/GoogleContainerTools/distroless
FROM cgr.dev/chainguard/wolfi-base

RUN adduser -D nonroot

USER nonroot

# more info at https://docs.github.com/en/packages/learn-github-packages/connecting-a-repository-to-a-package#connecting-a-repository-to-a-container-image-using-the-command-line
LABEL org.opencontainers.image.source=https://github.com/gsantoro/go-sample-app

COPY build/hello /app
CMD ["/app"]