## breadth first search

Two questions are asked:

- Question type 1: Is there a path from node A to node B?
- Question type 2: What is the shortest path from node A to node B?

To impletement BFS, we could use a queue
Which is a data structure that let uses LIFO pattern, contrast to a stack that uses a FIFO pattern

We can also use Djiktra's algorithm to perform a search.
Djiktra works on Direct Acyclic Graph - graph with weigted nodes and have arrows to their next item

Dijktra is not good for graph with negative weight. Use Bellman-Ford algorithm instead.

To implement Dijktra maintain tree hash tables

- Graph
- Costs
- Parents

Breadth-first search is used to calculate the shortest path for an unweighted graph.
• Dijkstra’s algorithm is used to calculate the shortest path for a weighted graph.
• Dijkstra’s algorithm works when all the weights are positive.
• If you have negative weights, use the Bellman-Ford algorithm.