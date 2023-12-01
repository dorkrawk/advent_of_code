sum_of_values = 0

File.each_line("./day_1_input.txt") do |line|

  line.each_char do |c|
    if c.to_i?
      sum_of_values += c.to_i * 10
      break
    end
  end  
  line.reverse.each_char do |c|
    if c.to_i?
      sum_of_values += c.to_i
      break
    end
  end  
end

puts sum_of_values
