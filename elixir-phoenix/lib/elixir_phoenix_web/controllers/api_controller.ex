defmodule ElixirPhoenixWeb.ApiController do
  use ElixirPhoenixWeb, :controller

  def home(conn, _params) do
    json(conn, %{
      result: "Ok",
      hostname: :inet.gethostname() |> elem(1) |> to_string(),
    })
  end

  def db(conn, _params) do
    json(conn, %{
      result: "Not implemented!",
      hostname: :inet.gethostname() |> elem(1) |> to_string(),
    })
  end

  def health(conn, _params) do
    json(conn, %{
      result: "Ok",
      hostname: :inet.gethostname() |> elem(1) |> to_string(),
    })
  end

  def chaos(conn, _params) do
    json(conn, %{
      result: "Not implemented!",
      hostname: :inet.gethostname() |> elem(1) |> to_string(),
    })
  end
end
