#input_file = "day_2_sample.txt"
input_file = "day_2_input.txt"

defmodule AOC do
  def get_input(path) do
    path
    |> File.read!()
    |> String.trim()
    |> String.replace(~r/[\r\n]+/, "")
    |> String.split(",", trim: true)
  end

  def get_range_values(range) do
    [first, last] = 
      range
      |> String.split("-")
      |> Enum.map(&String.to_integer/1)
    first..last
  end

  def test_invalid(id) do
    id_str = Integer.to_string(id)
    if rem(String.length(id_str), 2) != 0 do
      false  # early return for odd length numbers
    else
      {front, back} = String.split_at(id_str, div(String.length(id_str), 2))
      front == back
    end
  end
end

ranges = AOC.get_input(input_file)

answer_sum =
  ranges
  |> Enum.flat_map(&AOC.get_range_values/1)
  |> Enum.filter(&AOC.test_invalid/1)
  |> Enum.sum()

IO.puts(answer_sum)
