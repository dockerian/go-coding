/**
 * @see https://leetcode.com/problems/single-number/
 * Given an array of integers, every element appears twice except for one.
 * Find that single one.
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
    // note: arr.reduce(function callback(previous, current, currentIndex) { return current }, initialValue)
    // nums.reduce((p, c) => p^c)
    return nums.reduce(function (prev, curr, currIndx) {
        return prev ^ curr;
    });
};
