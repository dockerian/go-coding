function isString(str) {
    return typeof str === 'string' || str instanceof String
}

/**
 * Implement regular expression isMatch with support for '.' and '*'.
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 * @see https://leetcode.com/problems/regular-expression-matching/
 */
var isMatch = function(s, p) {
  if (!isString(s) || !isString(p)) return false
  if (s === null || p === null) return false;
  // console.log(`s= '${s}' [${s.length}], p= '${p}' [${p.length}]`)

  var dp = new Array(s.length + 1);
  for (var i = 0; i <= s.length; i++) {
    dp[i] = new Array(p.length + 1);
    for (var j = 0; j <= p.length; j++) {
      dp[i][j] = false;
    }
  }

  dp[0][0] = true;
  for (j = 2; j <= p.length; j++) {
    dp[0][j] = dp[0][j-2] && p[j-1] == '*';
  }

  for (i = 1 ; i <= s.length; i++) {
    for (j = 1; j <= p.length; j++) {
      if (p[j-1] == '.' || p[j-1] == s[i-1])
        dp[i][j] = dp[i-1][j-1];
      if (j >=2 && p[j-1] == '*') {
        if (p[j-2] == '.' || p[j-2] == s[i-1])
          dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i][j-2];
        else
          dp[i][j] = dp[i][j-2];
      }
    }
  }

  return dp[s.length][p.length];
};
