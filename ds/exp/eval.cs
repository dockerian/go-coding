using System;

class Program
{
  static void Main(string[] args)
  {
    string expression = "(((3.5 * 4.5) / (1 + 2)) + 5)";
    Console.WriteLine(string.Format("{0} = {1}", expression, new Expression.ExpressionTree(expression).Value));
    Console.WriteLine("\nShow's over folks, press a key to exit");
    Console.ReadKey(false);
  }
}

namespace Expression
{
  // -------------------------------------------------------

  abstract class NodeBase
  {
    public abstract double Value { get; }
  }

  // -------------------------------------------------------

  class ValueNode : NodeBase
  {
    public ValueNode(double value)
    {
      _double = value;
    }

    private double _double;
    public override double Value
    {
      get
      {
        return _double;
      }
    }
  }

  // -------------------------------------------------------

  abstract class ExpressionNodeBase : NodeBase
  {
    protected NodeBase GetNode(string expression)
    {
      // Remove parenthesis
      expression = RemoveParenthesis(expression);

      // Is expression just a number?
      double value = 0;
      if (double.TryParse(expression, out value))
      {
        return new ValueNode(value);
      }
      else
      {
        int pos = ParseExpression(expression);
        if (pos > 0)
        {
          string leftExpression = expression.Substring(0, pos - 1).Trim();
          string rightExpression = expression.Substring(pos).Trim();

          switch (expression.Substring(pos - 1, 1))
          {
            case "+":
              return new Add(leftExpression, rightExpression);
            case "-":
              return new Subtract(leftExpression, rightExpression);
            case "*":
              return new Multiply(leftExpression, rightExpression);
            case "/":
              return new Divide(leftExpression, rightExpression);
            default:
              throw new Exception("Unknown operator");
          }
        }
        else
        {
          throw new Exception("Unable to parse expression");
        }
      }
    }

    private string RemoveParenthesis(string expression)
    {
      if (expression.Contains("("))
      {
        expression = expression.Trim();

        int level = 0;
        int pos = 0;

        foreach (char token in expression.ToCharArray())
        {
          pos++;
          switch (token)
          {
            case '(':
              level++;
              break;
            case ')':
              level--;
              break;
          }

          if (level == 0)
          {
            break;
          }
        }

        if (level == 0 && pos == expression.Length)
        {
          expression = expression.Substring(1, expression.Length - 2);
          expression = RemoveParenthesis(expression);
        }
      }
      return expression;
    }

    private int ParseExpression(string expression)
    {
      int winningLevel = 0;
      byte winningTokenWeight = 0;
      int winningPos = 0;

      int level = 0;
      int pos = 0;

      foreach (char token in expression.ToCharArray())
      {
        pos++;

        switch (token)
        {
          case '(':
            level++;
            break;
          case ')':
            level--;
            break;
        }

        if (level <= winningLevel)
        {
          if (OperatorWeight(token) > winningTokenWeight)
          {
            winningLevel = level;
            winningTokenWeight = OperatorWeight(token);
            winningPos = pos;
          }
        }
      }
      return winningPos;
    }

    private byte OperatorWeight(char value)
    {
      switch (value)
      {
        case '+':
        case '-':
          return 3;
        case '*':
          return 2;
        case '/':
          return 1;
        default:
          return 0;
      }
    }
  }

  // -------------------------------------------------------

  class ExpressionTree : ExpressionNodeBase
  {
    protected NodeBase _rootNode;

    public ExpressionTree(string expression)
    {
      _rootNode = GetNode(expression);
    }

    public override double Value
    {
      get
      {
        return _rootNode.Value;
      }
    }
  }

  // -------------------------------------------------------

  abstract class OperatorNodeBase : ExpressionNodeBase
  {
    protected NodeBase _leftNode;
    protected NodeBase _rightNode;

    public OperatorNodeBase(string leftExpression, string rightExpression)
    {
      _leftNode = GetNode(leftExpression);
      _rightNode = GetNode(rightExpression);

    }
  }

  // -------------------------------------------------------

  class Add : OperatorNodeBase
  {
    public Add(string leftExpression, string rightExpression)
      : base(leftExpression, rightExpression)
    {
    }

    public override double Value
    {
      get
      {
        return _leftNode.Value + _rightNode.Value;
      }
    }
  }

  // -------------------------------------------------------

  class Subtract : OperatorNodeBase
  {
    public Subtract(string leftExpression, string rightExpression)
      : base(leftExpression, rightExpression)
    {
    }

    public override double Value
    {
      get
      {
        return _leftNode.Value - _rightNode.Value;
      }
    }
  }

  // -------------------------------------------------------

  class Divide : OperatorNodeBase
  {
    public Divide(string leftExpression, string rightExpression)
      : base(leftExpression, rightExpression)
    {
    }

    public override double Value
    {
      get
      {
        return _leftNode.Value / _rightNode.Value;
      }
    }
  }

  // -------------------------------------------------------

  class Multiply : OperatorNodeBase
  {
    public Multiply(string leftExpression, string rightExpression)
      : base(leftExpression, rightExpression)
    {
    }

    public override double Value
    {
      get
      {
        return _leftNode.Value * _rightNode.Value;
      }
    }
  }
}
