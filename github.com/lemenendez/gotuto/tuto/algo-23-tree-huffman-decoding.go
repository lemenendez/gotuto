/*
Huffman coding assigns variable length codewords to fixed length input characters based on their frequences.More frequent characters are assigned shorted codewords and less frequent characters are assigned longer codewords.All edges along the path to a character contain a code digit. If they are on the left side on the tree, they will be a 0 (zero). If on the right, they'll be a 1 (one). Only the leaves will contain a letter and its frequency count. All the nodes will contain a null instead of acharacter, and the count of the frequency of all its descendant characters.

For instance, consider the string ABRACADABRA. There are a total of 11 characters in the string.
This number should match the count in the ultimately determinated root of the tree. Our frequencies are A = 5, B = 2, R = 2, C =1, and D = 1. The two smallest frequencieas are for C and D, both equal to 1, so we'll create a tree with them.  The root node will contain the sume of te counts of tis descendants, in this case 1 + 1 = 2. The left node will be the first character encountered, C, and the right will contain D. Next we have 3 items with a characer count of 2: the tree we just created, the character B and the character R.

Te tree came first, so it will go on the left of our new root node. B will go on the right. Repeat until the tree is complete, then fill in the 1's and 0's for the edges. The finished graph looks like:


		ϕ,11
	0 /		   \
  A,5          ϕ, 6
			0 /     \ 1
		   R,2       ϕ,4
				  0	/    \ 1
				   ϕ     B,2
                 /   \
				C,1  D,1

Input

ϕϕϕϕ

				ϕ,5
			 0 /    \ 1
		     ϕ,2	ϕ,3
		   0 /  \ 1
		    ϕ     ϕ
		   B,1   C, 1

		s="1001011"

Sample output

S="1001011"
Processing the string from left to right.
S[0]='1' : we move to the right child of the root. We encounter a leaf node with value 'A'. We add 'A' to the decoded string.
We move back to the root.

S[1]='0' : we move to the left child.
S[2]='0' : we move to the left child. We encounter a leaf node with value 'B'. We add 'B' to the decoded string.
We move back to the root.

S[3] = '1' : we move to the right child of the root. We encounter a leaf node with value 'A'. We add 'A' to the decoded string.
We move back to the root.

S[4]='0' : we move to the left child.
S[5]='1' : we move to the right child. We encounter a leaf node with value C'. We add 'C' to the decoded string.
We move back to the root.

 S[6] = '1' : we move to the right child of the root. We encounter a leaf node with value 'A'. We add 'A' to the decoded string.
We move back to the root.

Decoded String = "ABACA"
https://en.wikipedia.org/wiki/Huffman_coding

*/

package main

func main() {

}
