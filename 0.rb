#! /bin/ruby
#参考：http://d.hatena.ne.jp/E_Mattsan/20121122/1353588273
require 'benchmark'

def subsequence a
  result = [[]]
  a.each do |i|
    result.size.times do |j|
      result.push(result[j].dup.push(i))
    end
  end
  result
end

x = gets.chomp.split("")
y = gets.chomp
Benchmark.bm 10 do |r|
	r.report "subs" do
		p (subsequence x).select {|s| y.include?(s.join)}.map{|s| s.join}
	end
end
