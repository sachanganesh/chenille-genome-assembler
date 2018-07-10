module Velour

	include("KmerIO.jl")
	using .KmerIO
	using ProfileView

	function main()
		fragments = ["data/staphylococcus_aureus/frag1.fastq", "data/staphylococcus_aureus/frag2.fastq"]
		# graph = KmerIO.DeBruijnGraph.new_graph()

		for fragment in fragments
			KmerIO.graph_from_fastq(fragment)
			# println(graph.num_nodes)
			break
		end
	end

end
