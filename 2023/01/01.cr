sum_of_values = 0

numbers = [
  "one",
  "two",
  "three",
  "four",
  "five",
  "six",
  "seven",
  "eight",
  "nine"
]

numbers_to_int = {
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9
}

input = "./day_1_input.txt" 
# input = "./day_1_sample.txt"

File.each_line(input) do |line|
  number_indexes = numbers.map do |n|
    [n, line.index(n), line.rindex(n)]
  end

  first_number = number_indexes.min_by { |n| n[1].nil? ? Int32::MAX : n[1].as(Int32) }
  last_number = number_indexes.max_by { |n| n[2].nil? ? Int32::MIN : n[2].as(Int32) }

  first_value_found = false
  # puts line
  line.each_char_with_index do |c, i|
    if c.to_i?
      if first_number[1].nil? || i < first_number[1].as(Int32)
        # puts c
        sum_of_values += c.to_i * 10
      else
        # puts first_number[0]
        sum_of_values += numbers_to_int[first_number[0].as(String)] * 10
      end
      first_value_found = true
      break
    end
  end  
  unless first_value_found
    # puts first_number[0]
    sum_of_values += numbers_to_int[first_number[0].as(String)] * 10
  end

  last_value_found = false
  line.reverse.each_char_with_index do |c, i|
    if c.to_i?
      if last_number[2].nil? || (line.size - i) > last_number[2].as(Int32)
        # puts c
        sum_of_values += c.to_i
      else
        # puts last_number[0]
        sum_of_values += numbers_to_int[last_number[0].as(String)]
      end
      last_value_found = true
      break
    end
  end  
  unless last_value_found
    # puts last_number[0]
    sum_of_values += numbers_to_int[last_number[0].as(String)]
  end
  # puts ""
end


puts sum_of_values
