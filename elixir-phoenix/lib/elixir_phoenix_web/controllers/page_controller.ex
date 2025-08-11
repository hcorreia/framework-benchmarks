defmodule ElixirPhoenixWeb.PageController do
  use ElixirPhoenixWeb, :controller

  def home(conn, _params) do
    render(conn, :home)
  end
end
