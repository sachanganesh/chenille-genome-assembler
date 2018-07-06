#!/usr/bin/env bash

mkdir -p data/staphylococcus_aureus

echo "Downloading Staphylococcus_aureus fragment 1"
wget "http://gage.cbcb.umd.edu/data/Staphylococcus_aureus/Data.original/frag_1.fastq.gz" -O data/staphylococcus_aureus/frag1.fastq.gz

echo "Extracting Staphylococcus_aureus fragment 1"
gunzip -k data/staphylococcus_aureus/frag1.fastq.gz

echo "Downloading Staphylococcus_aureus fragment 2"
wget "http://gage.cbcb.umd.edu/data/Staphylococcus_aureus/Data.original/frag_2.fastq.gz" -O data/staphylococcus_aureus/frag2.fastq.gz

echo "Extracting Staphylococcus_aureus fragment 2"
gunzip -k data/staphylococcus_aureus/frag2.fastq.gz
