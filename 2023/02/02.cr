class Game
  getter :rounds

  def initialize(r = 0, g = 0, b = 0)
    @rounds = [] of Hash(String, Int32)
  end  

  def add_round(r = 0, g = 0, b = 0)
    @rounds << Hash{
      "red" => r,
      "green" =>  g,
      "blue" => b
    }
  end

  def max_color(color)
    @rounds.map do |round|
      round[color]
    end.max
  end

  def min_power
    max_color("red") * max_color("green") * max_color("blue")
  end

  def valid_with_limits?(r, g, b)
    max_color("red") <= r && max_color("green") <= g && max_color("blue") <= b
  end
end

input = "./day_2_input.txt" 
# input = "./day_2_sample.txt"

RED_LIMIT = 12
GREEN_LIMIT = 13
BLUE_LIMIT = 14

game_id = 1

id_sum = 0
min_sum = 0

File.each_line(input) do |line|
  start_index = line.index(":") || 0
  game_text = line[start_index + 2..]
  round_texts = game_text.split("; ")

  game = Game.new

  round_texts.each do |round_text|
    round = Hash(String, Int32).new
    round["red"] = 0
    round["green"] = 0
    round["blue"] = 0
    cubes = round_text.split(", ")
    cubes.each do |cube|
      cube_count = cube.split(" ")
      round[cube_count[1]] = cube_count[0].to_i32
    end
    game.add_round(round["red"], round["green"], round["blue"])
  end
  
  if game.valid_with_limits?(RED_LIMIT, GREEN_LIMIT, BLUE_LIMIT)
    id_sum += game_id
  end

  min_sum += game.min_power

  game_id += 1
end

puts min_sum
