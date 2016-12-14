class PriorityQueue<T>
{
  IComparer<T> comparer;
  T[] heap;
  public PriorityQueue() : this(null) { }
  public PriorityQueue(int capacity) : this(capacity, null) { }
  public PriorityQueue(IComparer<T> comparer) : this(16, comparer) { }
  public PriorityQueue(int capacity, IComparer<T> comparer)
  {
    this.comparer = (comparer == null) ? Comparer<T>.Default : comparer;
    this.heap = new T[capacity];
  }

  public int Count { get; private set; }
  public T Peek()
  {
    if (Count > 0) return heap[0];
    throw new InvalidOperationException("The queue is empty");
  }
  public T Pop()
  {
    var v = Peek();
    heap[0] = heap[--this.Count];
    if (Count > 0) shiftDown(0);
    return v;
  }
  public void Push(T v)
  {
    if (Count >= heap.Length) Array.Resize(ref heap, Count * 2);
    heap[Count] = v;
    shiftUp(Count++);
  }

  private void shiftDown(int n)
  {
    var v = heap[n];
    for (var j = n * 2; 0 <= j && j < Count; n = j, j *= 2)
    {
      var j1 = j + 1;
      if (j1 > 0 && j1 < Count && comparer.Compare(heap[j1], heap[j]) > 0) j++;
      if (j != n && comparer.Compare(v, heap[j]) >= 0) break;
      heap[n] = heap[j];
    }
    heap[n] = v;
  }

  private void shiftUp(int n)
  {
    var v = heap[n];
    for (var i = n/2; n > 0 && comparer.Compare(v, heap[i]) > 0; i/=2) {
      heap[n] = heap[i];
      n = i;
    }
    heap[n] = v;
  }
}
