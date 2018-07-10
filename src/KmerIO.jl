module KmerIO

	include("DeBruijnGraph.jl")
	import .DeBruijnGraph

	export parseFASTQ

	function graph_from_fastq(filepath::String)
		graph = nothing

		open(filepath, "r") do f
			for (i, shortread) in enumerate(eachline(f))
				if (i - 2) % 4 == 0
					kmers = kmers_from_shortread(shortread, 13)
					# println(length(kmers))
				# 	graph = DeBruijnGraph.add_nodes(graph, kmers)
				end
			end
		end

		return graph
	end

	function kmers_from_shortread(shortread::String, k::Int64)::Array{String}
		num_kmers = length(shortread) - k + 1
		kmers = Array{String}(num_kmers)

		for i = 1 : num_kmers
			kmers[i] = clean_kmer(shortread[i : i + k - 1])
		end

		return kmers
	end

	function clean_kmer(kmer::String)::String
		str = Array{Char}(length(kmer) + 1)

		for (i, ch) in enumerate(kmer)
			if ch in "ACGT"
				setindex!(str, ch, i)
			else
				setindex!(str, 'A', i)
			end
		end

		return String(str)
	end

end
