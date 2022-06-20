FROM debian:buster-slim AS final

# 运行程序
CMD ["chmod", "+x", "./bootstrap.sh"]
CMD ["sh","./bootstrap.sh"]