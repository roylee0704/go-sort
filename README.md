# go-sort
Sorting Algorithms in Golang.


## Insertion Sort
> For i = 1,2,.., n
    Insert A[i] into sorted array A[0:i-1]
    by pairwise swap down to the correct position.

```sh
[5][2][4][6][1][3]
    |
  (key)
```

## Merge Sort
> Given an Array of size n,
  Splits into 2 Arrays of size n/2,
  Sort each of the 2 Arrays,
  Merge them up into a sorted Array of size n.
