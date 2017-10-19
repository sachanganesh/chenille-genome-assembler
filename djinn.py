import argparse
from random import randint
from difflib import SequenceMatcher
from graphviz import Digraph


def get_reads(seq, read_len, num_reads):
	reads = []

	for _ in range(num_reads):
		reads.append(fragment_read(seq, read_len))

	return reads


def fragment_read(seq, frag_len):
	ind = randint(0, len(seq) - 1)

	if ind + frag_len >= len(seq):
		return seq[ind:]
	else:
		return seq[ind : ind + frag_len]


def get_coverage(seq_len, reads, num_reads):
	avg_read_len = 0
	for read in reads:
		avg_read_len += len(read)
	avg_read_len /= len(reads)

	return num_reads * (avg_read_len / seq_len)


def get_kmers(k, reads, read_len):
	kmer_set = set()

	for read in reads:
		if k > read_len:
			print("k for kmer is greater than read length")
			exit(1)
		elif k <= len(read):
			for i in range(len(read) - k + 1):
				kmer_set.add(read[i : i + k])

	return kmer_set


def find_next_kmers(kmer, kmer_set):
	followers = []

	for other in kmer_set:
		if kmer != other and kmer[1:] == other[:-1]:
			followers.append(other)

	return followers


def build_debruijn(kmer_set):
	graph = {}
	reverse_graph = {}

	for kmer in kmer_set:
		followers = find_next_kmers(kmer, kmer_set)
		graph[kmer] = set(followers)

		for follower in followers:
			if follower in reverse_graph:
				reverse_graph[follower].add(kmer)
			else:
				reverse_graph[follower] = set([kmer])

	for key in graph:
		if key not in reverse_graph:
			reverse_graph[key] = set()

	return graph, reverse_graph


def get_graph_ends(graph, reverse_graph):
	heads = []
	tails = []

	for key in graph:
		if len(reverse_graph[key]) == 0:
			heads.append(key)

	for key in reverse_graph:
		if len(graph[key]) == 0:
			tails.append(key)

	return heads, tails


def sieve_graph(graph):
	sieved_graph = {}

	for key in graph:
		if key in set.union(*graph.values()) or len(graph[key]) != 0:
			sieved_graph[key] = graph[key]

	return sieved_graph, graph


def simplify_degruijn(graph, reverse_graph):
	graph, _ = sieve_graph(graph)
	reverse_graph, _ = sieve_graph(reverse_graph)

	kmers = list(graph.keys())
	change = 1

	while change:
		graph_size = len(graph.keys())

		for kmer in kmers:
			word = kmer

			while word is not None and word in graph:
				if len(graph[word]) == 1 and len(reverse_graph[list(graph[word])[0]]) == 1:
					follower = list(graph[word])[0]

					s = SequenceMatcher(None, word, follower)
					alpha, beta, overlap_len = s.find_longest_match(0, len(word), 0, len(follower))
					new_word = word[:alpha] + word[alpha : alpha + overlap_len] + follower[beta + overlap_len:]

					# set new word in graph with appropriate chain
					graph[new_word] = graph[follower]

					# set new word in reverse_graph with appropriate chain
					reverse_graph[new_word] = reverse_graph[word]

					# get predecessors of word from reverse_graph
					word_predecessors = list(reverse_graph[word])
					follower_antecessors = list(graph[follower])

					# for each predecessor in graph, link new word in with existing chain
					for predecessor in word_predecessors:
						graph[predecessor].remove(word)
						graph[predecessor].add(new_word)

					# for each antecessor in reverse_graph, link new word in with existing chain
					for antecessor in follower_antecessors:
						reverse_graph[antecessor].remove(follower)
						reverse_graph[antecessor].add(new_word)

					# empty original graph nodes
					graph.pop(word, 0)
					graph.pop(follower, 0)

					# empty original reverse_graph nodes
					reverse_graph.pop(word, 0)
					reverse_graph.pop(follower, 0)

					word = new_word
				else:
					word = None

		change = graph_size - len(graph.keys())

	return graph, reverse_graph


def visualize_graph(graph):
	dot = Digraph(comment="de Bruijn graph for assembly")

	for kmer in graph:
		dot.node(kmer, kmer)

		for follower in graph[kmer]:
			dot.edge(kmer, follower)

	dot.attr(rankdir="LR")
	dot.render("debruijn.gv", view=True)


def prepare_arguments():
	parser = argparse.ArgumentParser(description="Djinn: Naive de novo genome assembler")
	parser.add_argument("read_len", metavar="L", type=int, help="length of reads")
	parser.add_argument("num_reads", metavar="N", type=int, help="number of reads")
	parser.add_argument("k", metavar="k", type=int, default=4, help="kmer length")

	args = parser.parse_args()
	return args.read_len, args.num_reads, args.k


def main():
	read_len, num_reads, k = prepare_arguments()

	sample = "ATGGAAGTCGCGGAATC"

	reads = get_reads(sample, read_len, num_reads)

	coverage = get_coverage(len(sample), reads, num_reads)

	kmers = sorted(get_kmers(k, reads, read_len))

	debruijn, rev_debruijn = build_debruijn(kmers)
	simp_debruijn, simp_rev_debruijn = simplify_degruijn(debruijn, rev_debruijn)

	print("Sequence:", sample)
	print("\nCoverage:", coverage)
	print("Assembled Contigs:")
	for key in simp_debruijn.keys():
		print("\t%s" % key)

	visualize_graph(simp_debruijn)

if __name__ == "__main__":
	main()
