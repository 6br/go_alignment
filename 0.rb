#! /bin/ruby
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

#x = gets.chomp.split("")
#y = gets.chomp
xall = 'jkahkncjknewrkfiljsklhlsfhskujejjflwjklnvmxcnvlcsdjfjjsdljdslfjsljfhedjwljkshfuejcklsjs'
yall = 'klvnwoihwoihtewkllnxcnvsmvmjsdnjkjnshuvhsuiujeijwiodkakcopjnsdvsvbfsfvxjhsduifjkshskfrf'

last = 20

x = xall[0..last].split("")
y = yall[0..last]

Benchmark.bm 1 do |r|
	r.report "subs" do
		p (subsequence x).select {|s| y.include?(s.join)}.map{|s| s.join}
	end
end

