# Binary Heap


<a name="ds"><br/></a>
## Data Structure

A binary heap is a heap data struc­ture cre­ated using a binary tree.

### Binary heap has two rules:
- **Shape** prop­erty: Binary Heap has to be com­plete binary tree at all lev­els except the last level.
- **Heap** prop­erty: All nodes are either greater than equal to (Max-Heap) or less than equal to (Min-Heap) to each of its child nodes.

### Implementation:
- Use array to store the data.
- Start stor­ing from index 1, not 0.
- For any given node at posi­tion `[i]`:
  - its Left Child is at `[2*i]` if available.
  - its Right Child is at `[2*i+1]` if available.
  - its Par­ent Node is at `[i/2]` if avail­able.


<a name="heap"><br/></a>
## Heap Operations

### Insert
  - Add the ele­ment at the bot­tom leaf of the Heap.
  - Per­form the **Bubble-Up** operation: ~ O(log n)
    - Bubble-Up a.k.a up-heap, percolate-up, sift-up, trickle-up, heapify-up, or cascade-up
    - If inserted ele­ment is smaller than its par­ent node in case of Min-Heap OR greater than its par­ent node in case of Max-Heap, swap the ele­ment with its parent.
    - Keep repeat­ing the above step, if node reaches its cor­rect posi­tion, STOP.

### Extract-Min OR Extract-Max Operation
  - Take out the ele­ment from the root.( it will be min­i­mum in case of Min-Heap and max­i­mum in case of Max-Heap).
  - Take out the last ele­ment from the last level from the heap and replace the root with the element.
  - Per­form Sink-Down

### Delete
  - Find the index for the ele­ment to be deleted.
  - Take out the last ele­ment from the last level from the heap and replace the index with this element.
  - Per­form **Sink-Down** operation: ~ O(log n)
    - If replaced ele­ment is greater than any of its child node in case of Min-Heap OR smaller than any if its child node in case of Max-Heap, swap the ele­ment with its small­est child (Min-Heap) or with its great­est child (Max-Heap).
    - Keep repeat­ing the above step, if node reaches its cor­rect posi­tion, STOP.


<a name="see"><br/></a>
See http://algorithms.tutorialhorizon.com/binary-min-max-heap/
