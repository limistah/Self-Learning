def find_lowest_cost_node(costs):
    lowest_cost = float("inf")
    lowest_cost_node = None
    for node in costs: # Go through each node.
        cost = costs[node]
        if cost < lowest_cost:
            lowest_cost = cost
            lowest_cost_node = node
        return lowest_cost_node