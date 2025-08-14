FROM elixir:1.18 AS build

WORKDIR /app

ENV MIX_ENV=prod

COPY mix.exs mix.lock ./
RUN mix deps.get --only prod

COPY . .

RUN mix compile
# RUN mix assets.deploy

CMD ["mix", "phx.server"]
