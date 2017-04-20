# Parse Tree

## Parser

See http://parsingintro.sourceforge.net/

## Mathmetical Expression

See

  - [ANTLR 3](https://theantlrguy.atlassian.net/wiki/display/ANTLR3/Five+minute+introduction+to+ANTLR+3)

  - [Expression Evaluation](http://www.geeksforgeeks.org/expression-evaluation/) |
    [source](evaluation.java)

  - [Evaluating math expression](http://stackoverflow.com/questions/3422673/evaluating-a-math-expression-given-in-string-form) |
    [source](eval.java)

  - [Evaluation in C#](http://stackoverflow.com/questions/5838918/evaluate-c-sharp-string-with-math-operators) |
    [source](eval.cs)

  - [Parsing arithmetic expression](http://codereview.stackexchange.com/questions/56999/parsing-arithmetic-expressions-like-x-y-a-b-z) |
    [source](parser.java)

  - [Parse Tree](http://interactivepython.org/courselib/static/pythonds/Trees/ParseTree.html)

  - ...

```
/*
 C++ Solution approach
 Using a handmade LL(1) Parser without real lexical analysis,
 so whitespaces must not occure in the input string and token
 recogniztion is primitive

 It's a bit an overkill since it doesn't make use of all the information
 given by the interviewer one can assume there is a "shortcut"

 however, I find it rather exotic to try to find a special solution for a
 problem that fits so cleanly into a theory space well understood.

 Interesting situation in an interview :-)

 --
 EBNF
 PROGRAM	= INT, NEWLINE, {EXPRLINE}
 EXPRLINE	= EXPR, NEWLINE
 EXPR		= EXPR OP1 TERM | TERM
 TERM		= TERM OP2 FACT | TERM
 FACT		= INT | '(' EXPR ')'
 OP1		= '+' | '-'
 OP2		= '*' | '/'
 INT		= ['-'] DIGITS
 DIGITS		= DIGIT {DIGIT}
 DIGIT		= '0' | '1' | '2' | '3' | ... | '9'
 NEWLINE	= '\n'

 EXPR  and TERM are not LL(1), so convert into LL1:
 EXPR		= TERM {OP1 TERM}
 TERM		= FACT {OP2 FACT}

 --
 NOTE:
 - {x}: x repeated n times, 0<=n
 - [y]: y occures 0 or 1 times
 */
#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Parser
{
private:
	string _buffer;
	int _position;

public:
	Parser(string buffer) : _buffer(buffer), _position(0) {}

	vector<int> Parse()
	{
		vector<int> result;
		Program(result);
		return result;
	}

private:
	//Production: PROGRAM	= INT, NEWLINE, {EXPRLINE}
	void Program(vector<int>& res)
	{
		int count = 0;
		if (!Int(count)) throw "E1";
		if (count < 0) throw "E2";
		if (!NewLine()) throw "E3";
		for (int i = 0; i < count; i++)
		{
			int eres = 0;
			if (!ExprLine(eres)) throw "E4";
			res.emplace_back(eres);
		}
	}

	// Production: EXPRLINE	= EXPR, NEWLINE
	bool ExprLine(int& res)
	{
		if (!Expr(res)) return false;
		if (!NewLine()) throw "E5";
	}

	// Production: EXPR = EXPR OP1 TERM | TERM --> EXPR := TERM {OP1 TERM}
	bool Expr(int& res)
	{
		int t2 = 0;
		char op;
		if (!Term(res)) return false;
		while (Op1(op))
		{
			if (!Term(t2)) throw "E7";
			if (op == '-') res -= t2; // should check for overflow, e.g. using safeint...
			else res += t2; // may overflow...
		}
		return true;
	}

	// Production: TERM = TERM OP2 FACT | TERM --> TERM = FACT {OP2 FACT}
	bool Term(int& res)
	{
		int f2;
		char op;

		if (!Fact(res)) return false;
		while (Op2(op))
		{
			if (!Term(f2)) throw "E9";
			if (op == '*') res *= f2; // may overflow...
			else res /= f2;
		}
		return true;
	}

	// Prodcution:  FACT = INT | '(' EXPR ')'
	bool Fact(int &res)
	{
		if (Int(res)) return true;
		if (ReadIf('('))
		{
			if (!Expr(res)) throw "E10";
			if (!ReadIf(')')) throw "E11";
			return true;
		}
		return false;
	}

	// Production: OP1 = '+' | '-'
	bool Op1(char& res)
	{
		res = Peek();
		return ReadIf('+') || ReadIf('-');
	}

	// Production: OP2 = '*' | '/'
	bool Op2(char& res)
	{
		res = Peek();
		return ReadIf('*') || ReadIf('/');
	}

	// NEWLINE
	bool NewLine()
	{
		return ReadIf('\n');
	}

	// INT
	bool Int(int& res)
	{
		res = 0;
		int sign = 1;
		string digits;

		if (ReadIf('-'))
		{
			sign = -1;
			if (!Digits(digits))
			{
				UnRead(1);
				return false;
			}
		}
		else
		{
			if (!Digits(digits)) return false;
		}

		for (int i = digits.size()-1, f=1; i >=0 ; i--, f*=10)
		{
			res += (digits[i] - '0') * f; // should check overflow with safeint or manually
		}

		res *= sign;
		return true;
	}

	// DIGITS
	bool Digits(string& digits)
	{
		char d;
		digits.clear();
		while (Digit(d))
		{
			digits.append(1, d);
		}
		return digits.size() > 0;
	}

	// DIGIT
	bool Digit(char& c)
	{
		return ReadIfInRange('0', '9', c);
	}

	// Gets the current character (0 at the first call, after Read 1, etc..)
	// returns '\0' if at end
	char Peek()
	{
		if(_position < _buffer.size()) return _buffer[_position];
		return '\0';
	}

	// Reads the current character if it is =1 (advances the current position)
	bool ReadIf(char a)
	{
		char b;
		return ReadIfInRange(a, a, b);
	}

	// reads the current character if in range [a..b]
	// e.g. a='0' b='2', reads it if Peek()=='0' or Peek()=='1' or Peek()=='2'
	bool ReadIfInRange(char a, char b, char& c)
	{
		c = Peek();
		if (c >= a && c <= b)
		{
			Read();
			return true;
		}
		return false;
	}

	// Read the current character
	// advances current position unless at end of string
	// note, '\0' may not occure in the string as symbol other then EOF
	char Read()
	{
		char p = Peek();
		if(p != 0) _position++;
		return p;
	}

	// To go back
	void UnRead(int n)
	{
		_position -= n;
	}
};

int main()
{
	string s1 = "1\n1+2+3*4\n";
	string s2 = "3\n"
		"19+12/4-((4-7)*3/1)\n"
		"1+(2-3)*4+5-6*8-(18*12*13)-(11/(5+2+4))\n"
		"((2+4)/3-2+1)\n";

	cout << "test string s1:" << endl << s1 << endl << endl;
	for (auto i : Parser(s1).Parse()) cout << i << endl;

	cout << endl << endl << "test string s2:" << endl << s2 << endl << endl;
	for (auto i : Parser(s2).Parse()) cout << i << endl;

}

```
