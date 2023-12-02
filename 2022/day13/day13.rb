require 'json'

def part_2
  file = File.read("input/day13.txt")
  file = file.split("\n")
  data = file.map do |a|
    if a.empty?
      next
    end
    JSON.parse(a)
  end.compact

  data = data.push([[2]], [[6]])

  a = data.sort do |l, r|
    compare(l, r) ? -1 : 1
  end

  (a.index([[2]])+1) * (a.index([[6]])+1)
end

def part_1
  file = File.read("input/day13.txt")
  file = file.split("\n\n")
  data = file.map do |a|
    left = JSON.parse(a.split("\n")[0])
    right = JSON.parse(a.split("\n")[1])

    [left, right]
  end

  count = 0
  data.each_with_index do |pair, id|
    left = pair[0]
    right = pair[1]

    id += 1

    if compare(left, right)
      count += id
    end
  end

  count
end

def compare(l, r)
  if l.is_a?(Array) && r.is_a?(Array)
    if l.empty? && r.empty?
      return :next
    end

    if l.empty?
      return true
    end

    if r.empty?
      return false
    end

    if l[0].nil? && r[0]
      return true
    end

    if l[0] && r[0].nil?
      return false
    end

    a = compare(l[0], r[0])
    if a == :next
      return compare(l[1..-1], r[1..-1])
    else
      return a
    end
  end

  if l.is_a?(Integer) && r.is_a?(Integer)
    if l != r
      return l < r
    else
      return :next
    end
  end

  if l.is_a?(Array) && r.is_a?(Integer)
    return compare(l, [r])
  end

  if l.is_a?(Integer) && r.is_a?(Array)
    return compare([l], r)
  end
end

p part_1
p part_2
