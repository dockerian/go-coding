using System;
using System.Collections;
using System.Collections.Generic;

public class Test{

  public static List<int> FindShortest(string[] matrix, int R, int start, int end)
  {
    int[] visited = new int[R]; // 1 - visited; 0 - not
    int[] prevous = new int[R];

    for (int i = 0; i < R; i++) {
      visited[i] = 0;
      prevous[i] = -1;
    }

    int pathLength = Int32.MaxValue;

    List<int> path = new List<int>();
    Queue q = new Queue();
    q.Enqueue(start);

    while (q.Count != 0) {
      int curr = (int)q.Dequeue();
      visited[curr] = 1;

      if (curr == end) {
        List<int> currPath = new List<int>();
        int currLength = 0;
        while (prevous[curr] != -1) {
          int prev = prevous[curr];
          currLength += matrix[prev][curr] - '0';
          if (prev != start) currPath.Add(prev);
          curr = prev;
        }

        if (currLength > 0 && currLength < pathLength) {
          path.Clear();
          path.AddRange(currPath);
          pathLength = currLength;
        }
      }

      int next = curr + 1;
      while (next < R) {
        if (matrix[curr][next] != '-' && visited[next] == 0) {
          prevous[next] = curr;
          q.Enqueue(next);
          break;
        }
        next++;
      }
    }

    return path;
  }

  public static void Main(){
    //Read input using Console
    int K = int.Parse(System.Console.ReadLine());
    string input = System.Console.ReadLine();
    String[] connections = input.Split(',');
    int R = connections.Length;
    List<int>[] path = new List<int>[R];

    for (var d = 1; d < K; d++) {
      path[d] = FindShortest(connections, R, 0, d);
    }

    int count = 0;

    for (var n = 1; n > K; n++) {
      int num = path[n].Count;
      if (num <= 1) {
        count++;
      }
    }

    System.Console.WriteLine(count);
  }
}
