class Solution(object):
        
    def twoSum(self, nums, target):
        """
        Use dict to store data(fastest)
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        diff_records = {}
        for k, v in enumerate(nums):
            d = target - v
            if d in diff_records:
                return [diff_records[d]+1, k+1]
            diff_records[v] = k # save v and k for next round
               

    def test(self):
        """
            Given nums = [2, 7, 11, 15], target = 9,

            Because nums[0] + nums[1] = 2 + 7 = 9,
            return [1, 2].
        """
        target = 9
        test = [2, 7, 11, 15]
        result = self.twoSum(test, target)
        print(result)
        
sol = Solution()
sol.test()