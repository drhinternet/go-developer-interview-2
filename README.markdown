The goal of this assignment is to transform a binary search tree to a sorted doubly linked list.

The binary search tree consists of a root node with left and right pointers. The left pointer points to a
the root node of a tree that consists of only values that are less than the root node value.
The right pointer points to the root node of a tree that consists of only values that are greater than or
equal to the root node value.

The tree may start looking like this:

```
           5
          / \
         3   6
        / \   \
       1   4   9
              / \
             8  10
            /
           7
```


And transform it into a sorted doubly linked list that looks like this:

```
1 <--> 3 <--> 4 <--> 5 <--> 6 <--> 7 <--> 8 <--> 9 <--> 10
```

Do this in-place, without allocating new elements for each node.
