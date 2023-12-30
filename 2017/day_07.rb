File.open("day_7_input.txt", "r") do |infile|
    parent = {}
    root_nodes = {}
    all_nodes = {}
    while (line = infile.gets)
      result = line.match(/(\w*)\s*\(\d*\)(\s*->\s*([\w|\s|,]*))*/)
      words = []
      parent = result[1].strip
      
      if all_nodes.has_key?parent
        parent_hash = all_nodes[parent]
      else 
        parent_hash = all_nodes[parent] = root_nodes[parent] = {}
      end
      
      unless result[3].nil? 
        children = result[3].gsub(",", '').split(" ").map(&:strip)
        children.each do |word|
          if root_nodes.has_key?(word)
            parent_hash[word] = root_nodes[word]
            root_nodes.delete(word)
          else
            parent_hash[word] = all_nodes[word] = {}
          end
        end
      end
    end
    puts root_nodes.keys
end
