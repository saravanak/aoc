require 'irb'

def abs_distance(distance, target, current, iteration)
  target > current + iteration ? iteration : (target - current + 1).abs
end

def remaining_distance(distance, target, current, iteration)
  target > current + iteration ? 0 : (iteration - (target - current)).abs
end

def move_01(distance, target, current, iteration)
  moved = abs_distance(distance[:x], target, current, iteration)
  { x: moved, y: distance[:y], moved: moved }
end

def move_02(distance, target, current, iteration)
  moved = remaining_distance(distance[:y], target, current, iteration)
  { x: distance[:x] , y: moved, moved: moved == 0 ? iteration : (target - current).abs }
end

def move_03(distance, target, current, iteration)
  moved = abs_distance(distance[:y], target, current, iteration)
  { x: iteration, y: moved, moved: moved }
end

def move_04(iteration, target, current)
  moved = remaining_distance(iteration, target, current)
  { x: moved , y: iteration, moved: moved == 0 ? iteration : (target - current).abs }
end

def move_05(iteration, target, current)
  moved = abs_distance(iteration, target, current)
  { x: moved , y: iteration, moved: moved}
end

def move_06(iteration, target, current)
  move_02(iteration, target, current)
end

def move_07(iteration, target, current)
  move_03(iteration, target, current)
end

def move_08(iteration, target, current)
  move_04(iteration, target, current)
end


target = 0
destination = ARGV[0].to_i
distance = {x: 0, y: 0, moved: 0}
moves = 0
one_round = [method(:move_01), method(:move_02), method(:move_03), method(:move_04), method(:move_05), method(:move_06), method(:move_07), method(:move_08)]
current = 1
while current <= destination do
  round_index = moves % 7
  iteration = moves / 7  + 1
  # if current == 3
  #   binding.irb
  # end
  method_to_call = one_round[round_index]
  distance = method_to_call.call(distance, destination, current, iteration)
  puts( "#{distance}: #{current}")
  current += distance[:moved]
  moves += 1
end

puts distance
puts(distance[:x] + distance[:y])

  # --5-|-4--
  # |       |
  # 6       3
  # |       |
  # --  x ---
  # |       |
  # 7       2
  # |       |
  # --8-|-1--
