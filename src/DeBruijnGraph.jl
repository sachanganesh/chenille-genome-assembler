module DeBruijnGraph

	import Base
	export KmerGraph, new_graph, add_node, add_nodes

	mutable struct KmerEntry
		kmer::String
		frequency::Int64
		predecessors::Array{Int64}
		successors::Array{Int64}
	end

	mutable struct KmerGraph
		lookup::Dict{String, Int64}
		nodes::Array{KmerEntry}
		num_nodes::Int64
	end

	function new_graph()::KmerGraph
		return KmerGraph(Dict{String, Int64}(), KmerEntry[], 0)
	end

	function preceding_kmers(graph::KmerGraph, kmer::String, prec_kmer_id::Void=nothing)::Array{Int64}
		prec_kmer_ids = Int64[]

		if prec_kmer_id != nothing
			push!(prec_kmer_ids, prec_kmer_id)
		end

		base = kmer[2 : length(kmer)]
		for nt in String["A", "C", "G", "T"]
			prec_kmer = nt * base

			if (prec_kmer_id != nothing && graph.nodes[prec_kmer_id].kmer != prec_kmer) || haskey(graph.lookup, prec_kmer)
				push!(prec_kmer_ids, graph.lookup[prec_kmer])
			end
		end

		return prec_kmer_ids
	end

	function add_node(graph::KmerGraph, kmer::String, prec_kmer_id::Void=nothing)::Int64
		prec_kmers = preceding_kmers(graph, kmer, prec_kmer_id)
		kmer_entry = nothing

		if haskey(graph.lookup, kmer)
			kmer_id = graph.lookup[kmer]
			kmer_entry = graph.nodes[kmer_id]
		else
			kmer_id = graph.num_nodes + 1
			graph.num_nodes += 1

			push!(graph.nodes, KmerEntry(kmer, 0, Int64[], Int64[]))
			kmer_entry = graph.nodes[kmer_id]
		end

		kmer_entry.frequency += 1

		for prec_id in prec_kmers
			if prec_id in graph.nodes[kmer_id].predecessors
				push!(graph.nodes[kmer_id].predecessors, prec_id)
			end

			if kmer_id in graph.nodes[prec_id].successors
				push!(graph.nodes[prec_id].successors, kmer_id)
			end
		end

		return kmer_id
	end

	function add_nodes(graph::KmerGraph, kmers::Array{String})
		for i = 1 : length(kmers)
			prev_id = nothing

			if i == 1
				prev_id = add_node(graph, kmers[i], nothing)
			else
				prev_id = add_node(graph, kmers[i], prev_id)
			end
		end

		return graph
	end

end
