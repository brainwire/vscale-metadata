FROM            alpine
# Add the binary. As it is statically linked, no need to add libc or anything else.
ADD             vscale-metadata /
# Open server port
EXPOSE          8080
# Set the entrypoint as the binary, so `docker run <image>` will behave as the binary
ENTRYPOINT      ["/vscale-metadata"]
CMD             ["-p", "8080"]
