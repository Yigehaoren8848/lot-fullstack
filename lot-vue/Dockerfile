FROM nginx:1.20.1

# 创建一个新目录
RUN mkdir -p /front/elm
COPY nginx.conf /etc/nginx/nginx.conf
# 复制当前Dockerfile所在目录下的elm/static到新目录
COPY elm/static /front/elm/static
COPY elm/index.html /front
# 暴露Nginx监听的端口（默认为80）
EXPOSE 8088

# 可选：如果你需要运行其他启动命令，可以使用CMD指令
# CMD ["nginx", "-g", "daemon off;"]
