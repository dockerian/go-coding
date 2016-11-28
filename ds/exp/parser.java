public class Parser {

  private static final String P_CLOSE = ")";
  private static final String P_OPEN = "(";


  private Map<Integer, Expression> frontCachingChart = new HashMap<>();
  private Map<Integer, Expression> backCachingChart = new HashMap<>();

  // maps opener to closer
  private Map<Integer, Integer> closingMapper = new HashMap<>();
  // maps closer to opener
  private Map<Integer, Integer> openingMapper = new HashMap<>();

  private LinkedList<Op> opsList = new LinkedList<>();

  public Expression parse(String input) {
    init();

    String org = input;
    String[] tokens = extractTokens(input);

    LinkedList<Integer> stack = new LinkedList<>();
    int depth = 0;
    for (int i = 0; i < tokens.length; i++) {
      String token = tokens[i];
      if (P_OPEN.equals(token)) {
        depth++;
        stack.push(i);
      } else if (P_CLOSE.equals(token)) {
        depth--;
        Integer opener = stack.pop();
        closingMapper.put(opener, i);
        openingMapper.put(i, opener);
      } else {
        OpType opType = Op.getOp(token);
        if (opType != null) {
          Op op;
          if (!stack.isEmpty()) {
            Integer opener = stack.peek();
            op = new Op(i, depth, opener, opType);
          } else {
            op = new Op(i, depth, 0, opType); // top level might have no opener
            closingMapper.put(0, tokens.length - 1);
          }
          opsList.add(op);
        } else {
          // this is a literal
          Expression value = new Expression();
          value.literal = token;

          // literals must have open and close as well
          value.open = i;
          value.close = i;

          frontCachingChart.put(i, value);
          backCachingChart.put(i, value);

          // additional literal parsing for things like
          // ((a)) or (((a)))
          additionalWrappingParsing(i, i + 1, tokens, value);
        }
      }
    }
    Comparator<Op> opPriorityComparator = new Comparator<Op>() {
      @Override
      public int compare(Op o1, Op o2) {
        // we want high priority op type first
        return o2.type.priority - o1.type.priority;
      }
    };
    // make sure higher priority op types are looked at first
    Collections.sort(opsList, opPriorityComparator);

    // make sure deeper ops are looked at first, the sort is stable so higher priority of same depth come first
    // we will do bottom up parsing in this case to ease parsing left to right of expression like a + b + c + d
    Collections.sort(opsList);

    Expression ex = parse(tokens);
    ex.org = org;
    return ex;
  }

  private void init() {
    frontCachingChart.clear();
    backCachingChart.clear();
    closingMapper.clear();
    openingMapper.clear();
    opsList.clear();
  }

  private String[] extractTokens(String input) {
    input = input.replaceAll("\\(", " ( ");
    input = input.replaceAll("\\)", " ) ");
    input = input.replaceAll("\\+", " + ");
    input = input.replaceAll("-", " - ");
    input = input.replaceAll("\\*", " * ");
    input = input.replaceAll("/", " / ");

    String[] tokens = input.trim().split("\\s+");
    return tokens;
  }

  private Expression parse(String[] tokens) {

    // this loop will be linear only
    Expression ret = null;
    while(!opsList.isEmpty()) {
      Op oneOp = opsList.poll();
      Expression subExpression = new Expression();

      subExpression.op = oneOp;
      subExpression.open = oneOp.opener;
      subExpression.close = closingMapper.get(subExpression.open);

      // working on left hand side
      int leftOfOpIndex = oneOp.index - 1;
      subExpression.left = backCachingChart.get(leftOfOpIndex);
      if (subExpression.left != null) {
        subExpression.open = subExpression.left.open;
      }

      // working on right hand side
      int rightOfOpIndex = oneOp.index + 1;
      subExpression.right = frontCachingChart.get(rightOfOpIndex);
      subExpression.close = subExpression.right.close;

      // simplify the special case of "(-2)"
      if (subExpression.left == null && subExpression.op.type == OpType.MINUS) {
        subExpression.literal = "-" + subExpression.right.literal;
        subExpression.op = null;
        subExpression.open = oneOp.index;
      }

      frontCachingChart.put(subExpression.open, subExpression);
      backCachingChart.put(subExpression.close, subExpression);

      additionalWrappingParsing(subExpression.open,
          subExpression.close + 1, tokens, subExpression);
      ret = subExpression;
    }

    return ret;
  }

  /**
   * Dealing with cases of '((a))' or '(((a)))' or '(((a + b)))'
   * @param subExpression
   */
  private void additionalWrappingParsing(int start, int end, String[] tokens, Expression expression) {
    int diff = 1;
    int leftIndex = start - diff;
    int rightIndex = end + diff - 1;
    while(leftIndex >= 0 && rightIndex < tokens.length) {
      String left = tokens[leftIndex];
      String right = tokens[rightIndex];
      if (P_OPEN.equals(left) && P_CLOSE.equals(right)) {
        // literals must have open and close too
        expression.open = leftIndex;
        expression.close = rightIndex;

        frontCachingChart.put(leftIndex, expression);
        backCachingChart.put(rightIndex, expression);
      } else {
        break;
      }
      diff++;
      leftIndex= start - diff;
      rightIndex = end + diff - 1;
    }
  }

  static class Expression {
    int open;
    int close;
    Op op;

    Expression left;
    Expression right;

    String literal;

    String org;

    @Override
    public String toString() {
      if (literal != null) {
        return literal;
      } else {
        return P_OPEN + left + ") " + op.type + " (" + right + P_CLOSE;
      }
    }
  }

  static class Op implements Comparable<Op> {
    OpType type;
    int index;
    int depth;
    int opener;

    Op(int index, int depth, int opener, OpType type) {
      this.index = index;
      this.depth = depth;
      this.opener = opener;
      this.type = type;
    }

    @Override
    public int compareTo(Op o) {
      return o.depth - depth;
    }

    public static OpType getOp(String token) {
      if ("+".equals(token)) {
        return OpType.PLUS;
      } else if ("-".equals(token)) {
        return OpType.MINUS;
      } else if ("*".equals(token)) {
        return OpType.MUL;
      } else if ("/".equals(token)) {
        return OpType.DIV;
      }
      return null;
    }

  }

  static enum OpType {
    PLUS(0, "+"), MINUS(0, "-"), MUL(1, "*"), DIV(1, "/");

    int priority;
    String str;

    private OpType(int priority, String str) {
      this.priority = priority;
      this.str = str;
    }

  }
}
