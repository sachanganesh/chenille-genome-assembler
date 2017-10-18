from random import randint


sample = "ATGGAAGTCGCGGAATC"


def get_reads(num_reads):
	reads = []

	for _ in range(num_reads):
		reads.append(fragment_read(sample, 6))


def fragment_read(seq, frag_len):
	ind = randint(0, len(seq) - 1)

	if ind + frag_len >= len(seq):
		return seq[ind:]
	else:
		return seq[ind : ind + frag_len]


def main():
	get_reads(20)


if __name__ == "__main__":
	main()
