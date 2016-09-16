/**
 * Implement regular expression matching with support for '.' and '*'.
 * @see https://leetcode.com/problems/regular-expression-matching/
 * solution from 王敏 (https://github.com/mallow111)
 */

public class BasicRegex {
  /**
   * Returns true if string s matches to regular expression p; otherwise false.
   * <p/>
   * 注意'*'是不可能单独出现的，一定要和前面的一个或多个字符出现，
   * 比如'a*',表示1个或多个a。所以判断的时候要看第二个字符是不是'*'：
   * 如果不是'*'，一定要保证第一个字符是相同的或者第一个字符是'.'；
   * 如果是'*'，就要用'*'后面的字符和ｓ去匹配，并且如果ｓ的首字母和ｐ的首字母可以匹配
   * （相同或者ｐ的第一个字母是'.'）那么就不断去除ｓ的第一个字符。
   * 在这个匹配过程里是ｐ在主导ｓ，比如如果ｐ是`null`或者`length==0`，那么ｓ一定也必是一样。
   * 每次`substring`的时候都要提前判断是否可以`substring`，也就是长度是否`>0`
   * e.g s="aaaabb", p="a*bb", 如果拿ｐ中的bb直接和ｓ匹配，结果肯定是求不出来的，
   * 所以我们不断去掉ｓ的第一个字母，最后 s="bb"，`p.substring(2)`就可以和ｓ匹配。
   * 还有就是如果ｐ的长度为`1`，那么ｓ的长度也一定要为`1`，否则也无法匹配。
   * <p/>
   * @param  s    an input string for regular expression test
   * @param  p    the regular expression to match the input string
   * @return      true if the input matches to the regular expression
   * @see https://discuss.leetcode.com/topic/40371/easy-dp-java-solution-with-detailed-explanation
   * Here are some conditions to figure out, then the logic can be very straightforward.
   * 1, If p.charAt(j) == s.charAt(i) :  dp[i][j] = dp[i-1][j-1];
   * 2, If p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1];
   * 3, If p.charAt(j) == '*':
   *    here are two sub conditions:
   *    1) if p.charAt(j-1) != s.charAt(i) :
   *          dp[i][j] = dp[i][j-2]  // in this case, a* only counts as empty
   *    2) if p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.':
   *          dp[i][j] = dp[i-1][j]  // in this case, a* counts as multiple a
   *       or dp[i][j] = dp[i][j-1]  // in this case, a* counts as single a
   *       or dp[i][j] = dp[i][j-2]  // in this case, a* counts as empty
   */
  public boolean isMatch(String s, String p) {
    if (s == null) return false;
    if (p == null) return s == null;
    if (p.length() == 0) return s.length() == 0;

    int sLen = s.length(), pLen = p.length();
    boolean [][]dp = new boolean[sLen + 1][pLen + 1];

    dp[0][0] = true;
    for (int i = 1; i <= sLen; i++) {
      dp[i][0] = false; // no need? as default to false?
    }

    for (int i = 2; i <= pLen; i++) {
      dp[0][i] = dp[0][i-2] && p.charAt(i-1) == '*';
    }

    for (int i = 1 ; i <= sLen; i++) {
      for (int j = 1; j <= pLen; j++) {
        if (p.charAt(j-1) == '.' || p.charAt(j-1) == s.charAt(i-1))
          dp[i][j] = dp[i-1][j-1];
        if (p.charAt(j-1) == '*') {
          if (p.charAt(j-2) == '.' || p.charAt(j-2) == s.charAt(i-1))
            dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i][j-2];
          else
            dp[i][j] = dp[i][j-2];
        }
      }
    }

    return dp[sLen][pLen];
  }

  public boolean isMatchRecursive(String s, String p) {
    if (p == null) return s == null;
    if(p.length() == 0) return s.length() == 0;

    // since p.length() == 0 has been checked, it will only need to
    // check s.length() for corner case test
    if (p.length() == 1) {
      return s.length() == 1 && (s.charAt(0) == p.charAt(0) || p.charAt(0) == '.');
    }
    if (p.charAt(1) != '*') {
      return s.length() > 0 && (s.charAt(0) == p.charAt(0) || p.charAt(0) == '.') && isMatchRecursive(s.substring(1), p.substring(1));
    }

    while(s.length() > 0 && (p.charAt(0) == s.charAt(0) || p.charAt(0) == '.')) {
      if (isMatchRecursive(s, p.substring(2))) return true;
      s = s.substring(1);
    }

    return isMatchRecursive(s, p.substring(2));
  }
}
