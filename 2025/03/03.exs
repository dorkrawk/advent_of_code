input_file = "day_3_sample.txt"
# input_file = "day_3_input.txt"

defmodule AOC do
  def get_input(path) do
    path
    |> File.read!()
    |> String.trim()
    |> String.split("\n", trim: true)
    |> Enum.map(&String.graphemes/1)
  end

  def find_max_joltage(bank) do
    int_bank = Enum.map(&String.to_integer/1)
    [max1 | rest_of_bank] = int_bank
    max2 = 0
    last_index = Enum.count(rest_of_bank)
    Enum.each(Enum.with_index(rest_of_bank), fn {battery, i} ->
      if i != last_index && battery > max1 do
        max1 = battery
      else
        if battery > max2 do
          max2 = battery
        end
      end
    end)
    {max1, max2}
  end
end

input = AOC.get_input(input_file)
result = Enum.map(input, &AOC.find_max_joltage/1)
IO.inspect(result)
