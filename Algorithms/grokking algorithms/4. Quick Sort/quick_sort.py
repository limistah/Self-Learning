def quick_sort(items):
    if len(items) < 2:
      return items
    pivot_index = 0
    pivot_item = items[pivot_index]
    greater = [i for i in items[1:] if i >= pivot_item]
    lesser = [i for i in items[1:] if i <= pivot_item]
    return quick_sort(lesser) + [pivot_item] + quick_sort(greater)

print(quick_sort([10,5,2,3]))
