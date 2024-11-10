FROM oven/bun

RUN mkdir /app
WORKDIR /app
COPY . .

RUN bun install

CMD ["bun", "start"]