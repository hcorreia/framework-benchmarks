defmodule ElixirPhoenixWeb.ApiController do
  use ElixirPhoenixWeb, :controller

  @hostname :inet.gethostname() |> elem(1) |> to_string()

  def home(conn, _params) do
    json(conn, %{
      result: "Ok",
      hostname: @hostname
    })
  end

  def db(conn, _params) do
    conn
    |> put_status(501)
    |> json(%{
      result: "Not implemented!",
      hostname: @hostname
    })
  end

  def health(conn, _params) do
    json(conn, %{
      result: "Ok",
      hostname: @hostname
    })
  end

  def chaos(conn, _params) do
    Req.get(Application.get_env(:elixir_phoenix, :chaos_endpoint))
    |> case do
      {:ok, %Req.Response{status: 200} = resp} ->
        json(conn, %{
          result: "Ok",
          hostname: @hostname,
          sleep_time: resp.body["sleep_time"]
        })

      _ ->
        conn
        |> put_status(503)
        |> json(%{
          result: "Chaos service is unavailable.",
          hostname: @hostname
        })
    end
  end
end
