The goal of this assignment is to transform a tree that might look like this:

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
1 <--> 2 <--> 3 <--> 4 <--> 5 <--> 6 <--> 7 <--> 8 <--> 9 <--> 10
```

Do this in-place, without allocating new elements for each node.
