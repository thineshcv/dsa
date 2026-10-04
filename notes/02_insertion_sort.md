#insertion_sort #sorting #algorithm


# insertion sort

Problem
**Input** : A sequence of *n* numbers `<a1,a2,...,an>`
**Output** : A permutation (reordering) of the input sequence such that `<a1' <= a2' <= ... <= an'`

## How insertion works
The core idea
- for each element
- take it as `key`
- Compare it with the elements to its element
- shift larger elements one position to its left
- insert `key` into the correct position

example:
```
[5] 2 4 6 1 3
 ↓
[2, 5] 4 6 1 3
 ↓
[2, 4, 5] 6 1 3
 ↓
[2, 4, 5, 6] 1 3
 ↓
[1, 2, 4, 5, 6] 3
 ↓
[1, 2, 3, 4, 5, 6]
```

## Psedo
``` Psedo
Insertion-Sort(A,n)
for i = 2 to n 
  key = A[i]
  j = i - 1
  while j > 0 and A[j] > key
    A[j+1] = A[j]
    j = j - | Column1 
  A[j+1] = key
```
```

```

Python implementation - ![[../code_snippets/python/01_insertion.sort.py]]
```
```
