
## Explanation
[The Boyer-Moore-Horspool algorithm](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm) is a [String-searching algorithm](https://en.wikipedia.org/wiki/String-searching_algorithm).

It does a pre-processing on the substring by creating a table with the jumps to do to avoid unnecessary research.

The algorithm is divided into 2 parts, a phase of construction of the jump table and a research phase.

### Building of the jump table [[Table Code](../bmh/table.go)]

For the construction of the table, it is sufficient to count the offset to be made from the end of the substring.

For `tuvzzzzzzzx` the table give `t=10, u=9, v=8, z=1, other=11`.

If the character repeats itself, we take the last occurrence.

The last character is not counted, if it is found in the substring, the last character will be entered in the table.

ex: `tuvzxzzzzzx` => `t=10, u=9, v=8, x=6, z=1, other=11` (This example is different from the animation below `tuvzxzzzzzx` != `tuvzzzzzzzx`).

![alt tag](assets/image1.gif)

### Search for the substring [[Search Code](../bmh/search.go)]

For research, we compare starting at the end of the substring.

This allows you to make jumps.

example:
```
wwwwwwwwww|w|wwwwwwwwwwwwww  
tuvzxzzzzz|x|

The characters are not identical, and "w" is not in the table.
So no reason to compare the previous characters, so we jump from the length of the substring (11)

wwwwwwwwwwwwwwwwwwwwwww|w|w  
             tuvzxzzzzz|x|

``` 

![alt tag](assets/image2.gif)