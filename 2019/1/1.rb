def fuel_for_module(mass)
    base_mass = (mass / 3) - 2 
    if base_mass.positive?
        base_mass + fuel_for_module(base_mass)
    else
        0
    end
end

total_fuel = 0
File.readlines('modules').each do |m|
    # fuel = (m.to_i / 3) - 2
    fuel = fuel_for_module(m.to_i)
    total_fuel += fuel
end
puts total_fuel
