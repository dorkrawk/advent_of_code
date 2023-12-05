input = "./day_4_input.txt" 
# input = "./day_4_sample.txt"

class Card
  @winners : Array(Int32)
  @numbers : Array(Int32)

  def initialize(winners, numbers)
    @winners = winners.map { |n| n.to_i32 }
    @numbers = numbers.map { |n| n.to_i32 }
  end

  def card_winners
    @winners & @numbers
  end

  def score
    return 0 if card_winners.size == 0
    2 ** (card_winners.size - 1)
  end
end

points_total = 0
card_counts = Hash(Int32, Int32).new
card_index = 1

File.each_line(input) do |line|
  start_index = line.index(":") || 0
  card_text = line[start_index + 2..]
  winners_numbers = card_text.split(" | ")
  card = Card.new(winners_numbers[0].split(" ").reject(""), winners_numbers[1].split(" ").reject(""))  

  points_total += card.score
  if card_counts.has_key?(card_index)
    card_counts[card_index] += 1
  else
    card_counts[card_index] = 1
  end

  winners_count = card.card_winners.size
  (card_index + 1..card_index+winners_count).each do |card_id|
    if card_counts.has_key?(card_id)
      card_counts[card_id] += card_counts[card_index]
    else
      card_counts[card_id] = card_counts[card_index]
    end
  end
  card_index += 1
end

# puts points_total
puts card_counts.values.sum
